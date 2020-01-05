package server

import (
	"context"
	"github.com/digitalocean/godo"
	"github.com/pepeunlimited/files/internal/app/app1/validator"
	"github.com/pepeunlimited/files/rpc"
	"github.com/pepeunlimited/files/spaces"
	"github.com/twitchtv/twirp"
	"log"
)

type FileServer struct {
	validator validator.FileServerValidator
	bucket    spaces.Bucket
	doClient    godo.Client
}

func (server FileServer) CreateBucket(ctx context.Context, params *rpc.CreateBucketParams) (*rpc.CreateBucketResponse, error) {
	log.Print("files-server: creating bucket..")
	exist, err := server.bucket.Exist(params.BucketName)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}
	if *exist {
		log.Printf("files-server: the bucket=%v already exist..", params.BucketName)
		return &rpc.CreateBucketResponse{}, nil
	}
	// create the bucket
	if err := server.bucket.Create(params.BucketName).Execute(); err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	if !params.IsCdn {
		return &rpc.CreateBucketResponse{}, nil
	}
	// create cdn for the bucket
	origin := params.BucketName+"."+params.Region+"."+params.Endpoint
	_, _, err = server.doClient.CDNs.Create(ctx, &godo.CDNCreateRequest{
		Origin: origin,
		TTL:    3600,
	})
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	return &rpc.CreateBucketResponse{}, nil
}

func (server FileServer) CreateFile(ctx context.Context, params *rpc.CreateFileParams) (*rpc.File, error) {
	//https://params.BucketName+region+endpoint+(fra1.digitaloceanspaces.com)+params.Filename
	//https://test-bucket-666.fra1.digitaloceanspaces.com/piia.txt
	return nil, nil
}



func NewFileServer(bucket spaces.Bucket, doClient godo.Client) FileServer {
	return FileServer{
		validator: validator.NewFileServerValidator(),
		bucket:    bucket,
		doClient:  doClient,
	}
}