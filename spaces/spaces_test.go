package spaces

import (
	"context"
	"github.com/digitalocean/godo"
	"github.com/pepeunlimited/files/do"
	"strings"
	"testing"
)

// NOTE: requires to fill these constants:
const (
	//SpacesEndpoint		 string = "-"
	//SpacesAccessKey 	   	 string = "-"
	//SpacesSecretKey        string = "-"
	//SpacesBucketName       string = "-"
	//DoAccessToken          string = "-"
)

func TestSpacesCreateDeleteBucketAndObject(t *testing.T) {
	spaces := NewSpaces(Endpoint, AccessKey, SecretKey)
	if err := spaces.
		Files(BucketName).
		Delete("simo.txt").
		Delete("piia.txt").
		Execute(); err != nil {
		t.Error(err)
		t.FailNow()
	}
	if err := spaces.Delete(BucketName); err != nil {
		t.Error(err)
		t.FailNow()
	}
	body := strings.NewReader("hello-world!")
	file := File{mimeType:"plain/text", fileSize: int64(body.Len()), body:body}
	if err := spaces.
		Create(BucketName).
		Create(file, FileMetaData{filename:"simo.txt", isPublic:true}).
		Create(file, FileMetaData{filename:"piia.txt", isPublic:true}). // throw error if file exist?
		Execute(); err != nil {
		t.Error(err)
		t.FailNow()
	}
	bytes, err := spaces.Files(BucketName).Get("simo.txt")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if len(bytes) != int(body.Size()) {
		t.FailNow()
	}
	cdnOrigin := BucketName+"."+Endpoint
	_, _, err = do.NewDoClient(DoAccessToken).CDNs.Create(context.Background(), &godo.CDNCreateRequest{
		Origin: cdnOrigin,
		TTL:    3600,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(string(bytes))
	exist, err := spaces.Exist(BucketName)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if !exist {
		t.FailNow()
	}
	exist, err = spaces.Exist("asdasasaaaa")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if exist {
		t.FailNow()
	}
	spaces.Files(BucketName).Delete("simo.txt").Delete("piia.txt").Execute()
}