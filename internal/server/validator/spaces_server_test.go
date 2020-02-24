package validator

import (
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/pepeunlimited/files/pkg/rpc/files"
	"log"
	"strings"
	"testing"
)

func TestFileServerValidator_GetFileByFileIDOk(t *testing.T) {
	params := &files.GetFileParams{
		FileId:   &wrappers.Int64Value{
			Value: 2,
		},
		Filename: nil,
	}
	err := NewSpacesServerValidator().GetFile(params)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestFileServerValidator_GetFileByFileFilename(t *testing.T) {
	params := &files.GetFileParams{
		FileId:   nil,
		Filename: &files.Filename{
			Name:     "a",
			BucketId: &wrappers.Int64Value{Value: 1},
		},
	}
	err := NewSpacesServerValidator().GetFile(params)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestFileServerValidator_GetFile(t *testing.T) {
	endpoint := "fra1.digitaloceanspaces.com"
	split := strings.Split(endpoint, ".")
	log.Print(split)
}