package server

import (
	"context"
	"github.com/pepeunlimited/files/do"
	"github.com/pepeunlimited/files/spaces"
	"testing"
)

func TestFileServer_CreateFile(t *testing.T) {

}


func TestFileServer_CreateBucket(t *testing.T) {
	doClient := do.NewDoClient(spaces.DoAccessToken)
	bucket := spaces.NewSpaces(spaces.Endpoint, spaces.AccessKey, spaces.SecretKey)

	if err := bucket.Delete(spaces.BucketName); err != nil {
		t.Error(err)
		t.FailNow()
	}


	server := NewFileServer(bucket, *doClient)

	err := server.CreateDOBucket(context.TODO(), CreateDOBucket{
		BucketName: spaces.BucketName,
		Endpoint:   spaces.Endpoint,
		IsCDN:		true,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}