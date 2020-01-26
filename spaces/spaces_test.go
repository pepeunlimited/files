package spaces

import (
	"context"
	"github.com/pepeunlimited/files/storage"
	"strings"
	"testing"
)

// NOTE: requires to fill these constants:
const (
	//SpacesEndpoint		 string = "-"
	//SpacesAccessKey 	   	 string = "-"
	//SpacesSecretKey        string = "-"
	//SpacesBucketName       string = "-"
	//AccessToken            string = "-"
)

func TestSpacesCreateDeleteBucketAndObject(t *testing.T) {
	ctx := context.TODO()
	spaces := storage.NewSpacesBuilder(Endpoint, AccessKey, SecretKey, BucketName, &AccessToken)
	spaces.Files().Delete("simo.txt").Delete("piia.txt").Execute(ctx)
	spaces.Delete()
	body := strings.NewReader("hello-world!")
	file := storage.File{MimeType:"plain/text", FileSize: int64(body.Len()), Body:body}
	if err := spaces.
		Create().
		Create(file, storage.FileMetaData{Filename:"simo.txt", IsPublic:true}).
		Create(file, storage.FileMetaData{Filename:"piia.txt", IsPublic:true}). // throw error if file exist?
		Execute(ctx); err != nil {
		t.Error(err)
		t.FailNow()
	}
	bytes, err := spaces.Files().Get("simo.txt")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if len(bytes) != int(body.Size()) {
		t.FailNow()
	}
	t.Log(string(bytes))
	exist, err := spaces.Exist()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if !exist {
		t.FailNow()
	}
	exist, err = storage.NewSpacesBuilder(Endpoint, AccessKey, SecretKey,"addsadsss", &AccessToken).Exist()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if exist {
		t.FailNow()
	}
	_, err = spaces.Files().GetMetadata("simo.txt")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	spaces.Files().Delete("simo.txt").Delete("piia.txt").Execute(ctx)
}

func TestBucket_Delete(t *testing.T) {
	spaces := storage.NewSpacesBuilder(Endpoint, AccessKey, SecretKey, BucketName, nil)
	if err := spaces.Delete(); err != nil {
		t.Error(err)
		t.FailNow()
	}
}