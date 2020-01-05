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


type CreateDOBucket struct {
	BucketName  string
	Endpoint    string
	CDNOrigin   *string
}

// only access point is from the server code. not provided any API's for it
func (server FileServer) CreateDOBucket(ctx context.Context, params CreateDOBucket) error {
	log.Print("files-server: creating DigitalOcean bucket..")
	exist, err := server.bucket.Exist(params.BucketName)
	if err != nil {
		return twirp.InternalErrorWith(err)
	}
	if *exist {
		log.Printf("files-server: the bucket=%v already exist in DigitalOcean's server..", params.BucketName)
		return nil
	}
	log.Printf("files-server: create new bucket=%v..",params.BucketName)
	// create the bucket
	if err := server.bucket.Create(params.BucketName).Execute(); err != nil {
		return twirp.InternalErrorWith(err)
	}

	if params.CDNOrigin == nil {
		return nil
	}
	// create the CDN's for bucket
	_, _, err = server.doClient.CDNs.Create(ctx, &godo.CDNCreateRequest{
		Origin: *params.CDNOrigin,
		TTL:    3600,
	})
	if err != nil {
		return twirp.InternalErrorWith(err)
	}
	return nil
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