package main

import (
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

	// spaces..
	//spacesBucketName := misc.GetEnv(spaces.SpacesBucketName, "")
	spacesAccessKey  := misc.GetEnv(spaces.SpacesAccessKey, "")
	spacesSecretKey  := misc.GetEnv(spaces.SpacesSecretKey, "")
	spacesEndpoint   := misc.GetEnv(spaces.SpacesEndpoint, "")

	bucket := spaces.NewSpaces(spacesEndpoint, spacesAccessKey, spacesSecretKey)

	doAccessToken    := misc.GetEnv(do.DoAccessToken, "")
	doClient 		 := do.NewDoClient(doAccessToken)
	//exist, err := spaces.Exist(spacesBucketName)
	//if err != nil {
	//	log.Panic("files-service: failed to validate does the bucket exist: "+err.Error())
	//}
	//if !*exist {
	//	log.Print("files-service: create the new [" + spacesBucketName + "] bucket..")
	//	// create the bucket
	//	if err := spaces.Create(spacesBucketName).Execute(); err != nil {
	//		log.Panic("files-service: failed to create the bucket: "+err.Error())
	//	}
	//	// create cdn for the bucket
	//	// cdn requires to initialize do client..
	//
	//	_, _, err := do.NewDoClient(doAccessToken).CDNs.Create(context.Background(), &godo.CDNCreateRequest{
	//		Origin: spacesBucketName + "." + spacesEndpoint,
	//		TTL:    3600,
	//	})
	//	if err != nil {
	//		spaces.Delete(spacesBucketName)
	//		log.Panic("files-service: failed to enabled cdn for the bucket: "+err.Error())
	//	}
	//} else {
	//	log.Print("files-service: bucket already exist: "+spacesBucketName)
	//}

	is := rpc.NewFileServiceServer(server.NewFileServer(bucket, *doClient), nil)
	mux := http.NewServeMux()
	mux.Handle(is.PathPrefix(), middleware.Adapt(is, headers.Username()))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Panic(err)
	}
}