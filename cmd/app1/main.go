package main

import (
	"context"
	"github.com/pepeunlimited/files/do"
	"github.com/pepeunlimited/files/internal/app/app1/server"
	"github.com/pepeunlimited/files/rpc"
	"github.com/pepeunlimited/files/spaces"
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

	// DOs Spaces
	spacesAccessKey    := misc.GetEnv(spaces.SpacesAccessKey, "")
	spacesSecretKey    := misc.GetEnv(spaces.SpacesSecretKey, "")
	spacesEndpoint     := misc.GetEnv(spaces.SpacesBucketEndpoint, "")
	spacesBucketName   := misc.GetEnv(spaces.SpacesBucketName, "")
	bucket := spaces.NewSpaces(spacesEndpoint, spacesAccessKey, spacesSecretKey)

	// DOs APIClient
	doAccessToken    := misc.GetEnv(do.DoAccessToken, "")
	doClient 		 := do.NewDoClient(doAccessToken)

	fileServer := server.NewFileServer(bucket, *doClient)

	// Create the Bucket if not exist...
	err := fileServer.CreateDOBucket(context.Background(), server.CreateDOBucket{
		BucketName: spacesBucketName,
		Endpoint:   spacesEndpoint,
		IsCDN:  	true,
	})

	if err != nil {
		log.Panic("files-server: interrupt server startup: "+err.Error())
	}

	is := rpc.NewFileServiceServer(fileServer, nil)
	mux := http.NewServeMux()
	mux.Handle(is.PathPrefix(), middleware.Adapt(is, headers.Username()))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Panic(err)
	}
}