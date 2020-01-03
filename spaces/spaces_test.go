package spaces

import (
	"log"
	"strings"
	"testing"
)

// NOTE: requires to fill these constants:
const (
	//endpoint	 	string = "{ENDPOINT}"
	//accessKey 	string = "{ACCESS_KEY}"
	//secretKey    	string = "{SECRET_KEY}"
	//bucketName   	string = "{BUCKET_NAME}"
)

func TestSpacesCreateDeleteBucketAndObject(t *testing.T) {
	spaces := NewSpaces(endpoint, accessKey, secretKey)
	if err := spaces.
		Files(bucketName).
		Delete("simo.txt").
		Delete("piia.txt").
		Execute(); err != nil {
		t.Error(err)
		t.FailNow()
	}
	if err := spaces.Delete(bucketName); err != nil {
		t.Error(err)
		t.FailNow()
	}
	body := strings.NewReader("hello-world!")
	file := File{mimeType:"plain/text", fileSize: int64(body.Len()), body:body}
	if err := spaces.
		Create(bucketName).
		Create(file, FileMetaData{filename:"simo.txt", isPublic:true}).
		Create(file, FileMetaData{filename:"piia.txt", isPublic:true}). // throw error if file exist?
		Execute(); err != nil {
		t.Error(err)
		t.FailNow()
	}
	bytes, err := spaces.Files(bucketName).Get("simo.txt")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	log.Print(string(bytes))
}