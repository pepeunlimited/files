package main

import (
	"github.com/pepeunlimited/files/internal/app/app1/server"
	"github.com/pepeunlimited/files/rpc"
	"github.com/pepeunlimited/microservice-kit/headers"
	"github.com/pepeunlimited/microservice-kit/middleware"
	"log"
	"net/http"
)

const (
	Version = "0.1"
)

func main() {
	log.Printf("Starting the TodoServer... version=[%v]", Version)

	is := rpc.NewFileServiceServer(server.NewFileServer(), nil)

	mux := http.NewServeMux()
	mux.Handle(is.PathPrefix(), middleware.Adapt(is, headers.Username()))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Panic(err)
	}
}