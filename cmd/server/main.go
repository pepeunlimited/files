package main

import (
	"github.com/pepeunlimited/authentication-twirp/authrpc"
	"github.com/pepeunlimited/files/internal/pkg/ent"
	"github.com/pepeunlimited/files/internal/pkg/upload"
	"github.com/pepeunlimited/files/internal/server/rest"
	"github.com/pepeunlimited/files/internal/server/twirp"
	"github.com/pepeunlimited/files/pkg/filesrpc"
	"github.com/pepeunlimited/microservice-kit/headers"
	"github.com/pepeunlimited/microservice-kit/middleware"
	"github.com/pepeunlimited/microservice-kit/misc"
	"log"
	"net/http"
)

const (
	Version = "0.1.7"
)


func main() {
	log.Printf("Starting the FilesServer... version=[%v]", Version)

	authenticationAddress := misc.GetEnv(authrpc.RpcAuthenticationHost, "http://api.dev.pepeunlimited.com")
	// ent
	ent 	 	 	 := ent.NewEntClient()

	// StorageService
	dos				 := upload.NewDos()
	//gcs 			 := upload.NewGoogleCloudStorage()

	sss := filesrpc.NewFilesServiceServer(twirp.NewFilesServer(dos, ent), nil)
	sus := rest.NewUploadServer(dos, ent, authrpc.NewAuthenticationServiceProtobufClient(authenticationAddress, http.DefaultClient))

	mux := http.NewServeMux()
	mux.Handle(sss.PathPrefix(), middleware.Adapt(sss, headers.Username()))
	mux.Handle(rest.UploadV1Files, sus.UploadV1Files())

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Panic(err)
	}
}