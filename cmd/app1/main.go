package main

import (
	rpc2 "github.com/pepeunlimited/authorization-twirp/rpc"
	"github.com/pepeunlimited/files/internal/app/app1/repository"
	"github.com/pepeunlimited/files/internal/app/app1/server"
	"github.com/pepeunlimited/files/internal/app/app1/upload"
	"github.com/pepeunlimited/files/rpc"
	"github.com/pepeunlimited/microservice-kit/headers"
	"github.com/pepeunlimited/microservice-kit/middleware"
	"github.com/pepeunlimited/microservice-kit/misc"
	"log"
	"net/http"
)

const (
	Version = "0.1"
)


func main() {
	log.Printf("Starting the FilesServer... version=[%v]", Version)


	authorizationAddress := misc.GetEnv(rpc2.RpcAuthorizationHost, "http://api.dev.pepeunlimited.com")
	// ent
	ent 	 	 	 := repository.NewEntClient()

	// DOsUpload
	dos				 := upload.NewDos()

	// DOs
	dfs := rpc.NewDOFileServiceServer(server.NewDOFileServer(dos, ent), nil)
	dus := server.NewDOFileUploadServer(dos, ent, rpc2.NewAuthorizationServiceProtobufClient(authorizationAddress, http.DefaultClient))

	mux := http.NewServeMux()
	mux.Handle(dfs.PathPrefix(), middleware.Adapt(dfs, headers.Username()))
	mux.Handle(server.UploadDOV1Files, dus.UploadDOV1Files())

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Panic(err)
	}
}