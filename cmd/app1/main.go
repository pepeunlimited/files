package main

import (
	"github.com/pepeunlimited/authentication-twirp/rpcauth"
	"github.com/pepeunlimited/files/internal/app/app1/mysql"
	"github.com/pepeunlimited/files/internal/app/app1/server"
	"github.com/pepeunlimited/files/internal/app/app1/upload"
	"github.com/pepeunlimited/files/rpcspaces"
	"github.com/pepeunlimited/microservice-kit/headers"
	"github.com/pepeunlimited/microservice-kit/middleware"
	"github.com/pepeunlimited/microservice-kit/misc"
	"log"
	"net/http"
)

const (
	Version = "0.1.4"
)


func main() {
	log.Printf("Starting the FilesServer... version=[%v]", Version)

	authenticationAddress := misc.GetEnv(rpcauth.RpcAuthenticationHost, "http://api.dev.pepeunlimited.com")
	// ent
	ent 	 	 	 := mysql.NewEntClient()

	// DOsUpload
	dos				 := upload.NewDos()

	// DOs
	sss := rpcspaces.NewSpacesServiceServer(server.NewSpacesServer(dos, ent), nil)
	sus := server.NewSpacesUploadServer(dos, ent, rpcauth.NewAuthenticationServiceProtobufClient(authenticationAddress, http.DefaultClient))

	mux := http.NewServeMux()
	mux.Handle(sss.PathPrefix(), middleware.Adapt(sss, headers.Username()))
	mux.Handle(server.UploadSpacesV1Files, sus.UploadSpacesV1Files())

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Panic(err)
	}
}