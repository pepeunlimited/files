package main

import (
	rpc2 "github.com/pepeunlimited/authorization-twirp/rpc"
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
	Version = "0.1.2"
)


func main() {
	log.Printf("Starting the FilesServer... version=[%v]", Version)


	authorizationAddress := misc.GetEnv(rpc2.RpcAuthorizationHost, "http://api.dev.pepeunlimited.com")
	// ent
	ent 	 	 	 := mysql.NewEntClient()

	// DOsUpload
	dos				 := upload.NewDos()

	// DOs
	sss := rpcspaces.NewSpacesServiceServer(server.NewSpacesServer(dos, ent), nil)
	sus := server.NewSpacesUploadServer(dos, ent, rpc2.NewAuthorizationServiceProtobufClient(authorizationAddress, http.DefaultClient))

	mux := http.NewServeMux()
	mux.Handle(sss.PathPrefix(), middleware.Adapt(sss, headers.Username()))
	mux.Handle(server.UploadSpacesV1Files, sus.UploadSpacesV1Files())

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Panic(err)
	}
}