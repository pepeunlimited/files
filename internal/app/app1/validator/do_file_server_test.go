package validator

import (
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/pepeunlimited/files/rpc"
	"log"
	"strings"
	"testing"
)

func TestFileServerValidator_GetFileByFileIDOk(t *testing.T) {
	params := &rpc.GetFileParams{
		FileId:   &wrappers.Int64Value{
			Value: 2,
		},
		Filename: nil,
	}
	err := NewFileServerValidator().GetFile(params)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestFileServerValidator_GetFileByFileFilename(t *testing.T) {
	params := &rpc.GetFileParams{
		FileId:   nil,
		Filename: &rpc.Filename{
			Name:     "a",
			SpacesId: &wrappers.Int64Value{Value: 1},
		},
	}
	err := NewFileServerValidator().GetFile(params)
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