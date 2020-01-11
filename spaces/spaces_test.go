package spaces

import (
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
	spaces := NewBucketCDN(Endpoint, AccessKey, SecretKey, BucketName, &AccessToken)
	spaces.Files().Delete("simo.txt").Delete("piia.txt").Execute()
	spaces.Delete()
	body := strings.NewReader("hello-world!")
	file := storage.File{MimeType:"plain/text", FileSize: int64(body.Len()), Body:body}
	if err := spaces.
		Create().
		Create(file, storage.FileMetaData{Filename:"simo.txt", IsPublic:true}).
		Create(file, storage.FileMetaData{Filename:"piia.txt", IsPublic:true}). // throw error if file exist?
		Execute(); err != nil {
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
	exist, err = NewBucketCDN(Endpoint, AccessKey, SecretKey,"addsadsss", &AccessToken).Exist()
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

	spaces.Files().Delete("simo.txt").Delete("piia.txt").Execute()
}

func TestBucket_Delete(t *testing.T) {
	spaces := NewBucket(Endpoint, AccessKey, SecretKey, BucketName)
	if err := spaces.Delete(); err != nil {
		t.Error(err)
		t.FailNow()
	}
}