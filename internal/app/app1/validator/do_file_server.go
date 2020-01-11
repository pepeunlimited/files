package validator

import (
	"github.com/pepeunlimited/files/rpc"
	"github.com/pepeunlimited/microservice-kit/validator"
	"github.com/twitchtv/twirp"
)

type FileServerValidator struct {}

func NewFileServerValidator() FileServerValidator {
	return FileServerValidator{}
}

func (valid FileServerValidator) GetFile(params *rpc.GetFileParams) error {
	if  params.Filename == nil && params.FileId == nil {
		return twirp.RequiredArgumentError("filename_or_fileId")
	}
	isFileId := params.FileId != nil && params.FileId.Value != 0
	if !isFileId && valid.filename(params.Filename)  {
		return twirp.RequiredArgumentError("filename_or_fileId")
	}
	return nil
}

func (valid FileServerValidator) Delete(params *rpc.DeleteParams) error {
	if  params.Filename == nil && params.FileId == nil {
		return twirp.RequiredArgumentError("filename_or_fileId")
	}
	isFileId := params.FileId != nil && params.FileId.Value != 0
	if !isFileId && valid.filename(params.Filename)  {
		return twirp.RequiredArgumentError("filename_or_fileId")
	}
	return nil
}

func (FileServerValidator) filename(params *rpc.Filename) bool {
	isFilename := params != nil && !validator.IsEmpty(params.Name)
	isBucketName := params.SpacesName != nil && !validator.IsEmpty(params.SpacesName.Value)
	isBucketId   := params.SpacesId != nil && params.SpacesId.Value != 0
	return !isFilename || !isBucketName && !isBucketId
}

func (valid FileServerValidator) CreateBucket(params *rpc.CreateSpacesParams) error {
	if validator.IsEmpty(params.Endpoint) {
		return twirp.RequiredArgumentError("endpoint")
	}
	if validator.IsEmpty(params.Name) {
		return twirp.RequiredArgumentError("name")
	}
	return nil
}