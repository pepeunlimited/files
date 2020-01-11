package rpc

import (
	"github.com/pepeunlimited/microservice-kit/rpcz"
	"github.com/pepeunlimited/microservice-kit/validator"
	"github.com/twitchtv/twirp"
)

const (
	FileNotFound     = "file_not_found"
	SpacesNotFound   = "spaces_not_found"
	FileUploadFailed = "file_upload_failed"
	FileExist        = "file_exist"
)

func IsReason(error twirp.Error, key string) bool {
	reason := error.Meta(rpcz.Reason)
	if validator.IsEmpty(reason) {
		return false
	}
	return reason == key
}