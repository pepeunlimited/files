package fs

import (
	"context"
	"log"
	"strings"
	"testing"
	"time"
)

// NOTE: requires to fill these constants:
const (
//SpacesEndpoint		 string = "-"
//SpacesAccessKey 	   	 string = "-"
//SpacesSecretKey        string = "-"
//SpacesBucketName       string = "-"
//PersonalAccessToken    string = "-"
)

func TestFilesystem_CreateBucketAndUploadFile(t *testing.T) {
	ctx := context.TODO()
	do := NewDigitalOcean(Endpoint, AccessKey, SecretKey, PersonalAccessToken)
	exist, err := do.BucketExist(BucketName)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if !exist {
		err := do.CreateBucket(ctx, BucketName)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
	}
	body := strings.NewReader("hello-world!")
	file := File{MimeType: "plain/text", FileSize: int64(body.Len()), Body:body}
	meta := FileMetaData{Filename: "simo.txt", IsPublic:true}
	err = do.UploadFile(ctx, file, meta, BucketName)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	bytes, err := do.GetFile(meta.Filename, BucketName)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if len(bytes) != int(body.Size()) {
		t.FailNow()
	}
	log.Print(string(bytes))
	time.Sleep(30 * time.Second)
	err = do.DeleteFile(meta.Filename, BucketName)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	err = do.DeleteBucket(BucketName)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
