package main

import (
	"github.com/pepeunlimited/authentication-twirp/pkg/rpc/auth"
	"github.com/pepeunlimited/files/internal/pkg/ent"
	"github.com/pepeunlimited/files/internal/server/twirp"
	"github.com/pepeunlimited/files/internal/server/upload"
	"github.com/pepeunlimited/files/pkg/fs"
	"github.com/pepeunlimited/files/pkg/rpc/files"
	"github.com/pepeunlimited/microservice-kit/middleware"
	"github.com/pepeunlimited/microservice-kit/misc"
	validator2 "github.com/pepeunlimited/microservice-kit/validator"
	"log"
	"net/http"
)

const (
	Version = "0.1.7.1"
)


func main() {
	log.Printf("Starting the FilesServer... version=[%v]", Version)

	authenticationAddress := misc.GetEnv(auth.RpcAuthenticationHost, "http://api.dev.pepeunlimited.com")
	authClient := auth.NewAuthenticationServiceProtobufClient(authenticationAddress, http.DefaultClient)

	var filesystem fs.FileSystem

	endpoint 			:= misc.GetEnv(fs.FsEndpoint, "")
	// DigitalOcean's
	spacesAccessKey 	:= misc.GetEnv(fs.SpacesAccessKey, "")
	spacesSecretKey 	:= misc.GetEnv(fs.SpacesSecretKey, "")
	doAccessToken 		:= misc.GetEnv(fs.DoAccessToken, "")

	// GoogleCloudService's
	gcsAccessKey 	:= misc.GetEnv(fs.GoogleCloudStorageAccessKey, "")
	gcsSecretKey 	:= misc.GetEnv(fs.GoogleCloudStorageSecretKey, "")

	if !validator2.IsEmpty(spacesAccessKey) && !validator2.IsEmpty(spacesSecretKey) && !validator2.IsEmpty(doAccessToken)  {
		log.Print("using filesystem as DigitalOcean's")
		filesystem = fs.NewDigitalOcean(endpoint, spacesAccessKey, spacesSecretKey, doAccessToken)
	} else if !validator2.IsEmpty(gcsAccessKey) && !validator2.IsEmpty(gcsSecretKey) {
		log.Print("using filesystem as Google Cloud Service's")
		filesystem = fs.NewGoogleCloud(endpoint, gcsAccessKey, gcsSecretKey)
	} else {
		log.Print("using filesystem as Mock")
		filesystem = fs.NewMock(endpoint)
	}

	client := ent.NewEntClient()
	sss := files.NewFilesServiceServer(twirp.NewFilesServer(client, filesystem), nil)
	sus := upload.NewUploadServer(client, authClient, filesystem)

	mux := http.NewServeMux()
	mux.Handle(sss.PathPrefix(), middleware.Adapt(sss))
	mux.Handle(upload.UploadV1Files, sus.UploadV1Files())

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Panic(err)
	}
}