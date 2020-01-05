package server

import (
	"context"
	"github.com/pepeunlimited/files/do"
	"github.com/pepeunlimited/files/rpc"
	"github.com/pepeunlimited/files/spaces"
	"testing"
)

func TestFileServer_CreateFile(t *testing.T) {

}


func TestFileServer_CreateBucket(t *testing.T) {
	doClient := do.NewDoClient(spaces.DoAccessToken)
	bucket := spaces.NewSpaces(spaces.Endpoint, spaces.AccessKey, spaces.SecretKey)

	server := NewFileServer(bucket, *doClient)
	server.CreateBucket(context.TODO(), &rpc.CreateBucketParams{
		BucketName: spaces.BucketName,
		Region:     "fra1",
		Endpoint:   "digitaloceanspaces.com",
		IsCdn:      true,
	})
}