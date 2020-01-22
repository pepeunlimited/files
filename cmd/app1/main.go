package main

import (
	"github.com/pepeunlimited/authentication-twirp/authrpc"
	"github.com/pepeunlimited/files/internal/app/app1/mysql"
	"github.com/pepeunlimited/files/internal/app/app1/server"
	"github.com/pepeunlimited/files/internal/app/app1/upload"
	"github.com/pepeunlimited/files/spacesrpc"
	"github.com/pepeunlimited/microservice-kit/headers"
	"github.com/pepeunlimited/microservice-kit/middleware"
	"github.com/pepeunlimited/microservice-kit/misc"
	"log"
	"net/http"
)

const (
	Version = "0.1.5"
)


func main() {
	log.Printf("Starting the FilesServer... version=[%v]", Version)

	authenticationAddress := misc.GetEnv(authrpc.RpcAuthenticationHost, "http://api.dev.pepeunlimited.com")
	// ent
	ent 	 	 	 := mysql.NewEntClient()

	// DOsUpload
	dos				 := upload.NewDos()

	// DOs
	sss := spacesrpc.NewSpacesServiceServer(server.NewSpacesServer(dos, ent), nil)
	sus := server.NewSpacesUploadServer(dos, ent, authrpc.NewAuthenticationServiceProtobufClient(authenticationAddress, http.DefaultClient))

	mux := http.NewServeMux()
	mux.Handle(sss.PathPrefix(), middleware.Adapt(sss, headers.Username()))
	mux.Handle(server.UploadSpacesV1Files, sus.UploadSpacesV1Files())

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Panic(err)
	}
}