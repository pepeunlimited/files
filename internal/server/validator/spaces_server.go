package validator

import (
	"github.com/pepeunlimited/files/pkg/rpc/files"
	"github.com/pepeunlimited/microservice-kit/validator"
	"github.com/twitchtv/twirp"
)

type SpacesServerValidator struct {}

func NewSpacesServerValidator() SpacesServerValidator {
	return SpacesServerValidator{}
}

func (valid SpacesServerValidator) GetFile(params *files.GetFileParams) error {
	if  params.Filename == nil && params.FileId == nil {
		return twirp.RequiredArgumentError("filename_or_fileId")
	}
	isFileId := params.FileId != nil && params.FileId.Value != 0
	if !isFileId && valid.filename(params.Filename)  {
		return twirp.RequiredArgumentError("filename_or_fileId")
	}
	return nil
}

func (valid SpacesServerValidator) Delete(params *files.DeleteParams) error {
	if  params.Filename == nil && params.FileId == nil {
		return twirp.RequiredArgumentError("filename_or_fileId")
	}
	isFileId := params.FileId != nil && params.FileId.Value != 0
	if !isFileId && valid.filename(params.Filename)  {
		return twirp.RequiredArgumentError("filename_or_fileId")
	}
	return nil
}

func (SpacesServerValidator) filename(params *files.Filename) bool {
	if params == nil {
		return false
	}
	isFilename := !validator.IsEmpty(params.Name)
	isBucketName := params.BucketName != nil && !validator.IsEmpty(params.BucketName.Value)
	isBucketId   := params.BucketId != nil && params.BucketId.Value != 0
	return !isFilename || !isBucketName && !isBucketId
}

func (valid SpacesServerValidator) CreateBucket(params *files.CreateBucketParams) error {
	if validator.IsEmpty(params.Name) {
		return twirp.RequiredArgumentError("name")
	}
	return nil
}