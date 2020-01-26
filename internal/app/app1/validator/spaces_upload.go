package validator

import (
	"encoding/json"
	"github.com/pepeunlimited/files/internal/app/app1/rest"
	"github.com/pepeunlimited/microservice-kit/httpz"
	"github.com/pepeunlimited/microservice-kit/jwt"
	"github.com/pepeunlimited/microservice-kit/validator"
	"io"
	"net/http"
	"strconv"
)

type UploadServerValidator struct {}

func NewSpacesUploadServerValidator() UploadServerValidator {
	return UploadServerValidator{}
}

func (UploadServerValidator) UploadSpacesV1Files(header http.Header, body io.ReadCloser) (*rest.UploadHeaders, *rest.MetaAPIArgs, error) {
	if body == nil {
		return nil, nil, httpz.NewMsgError("body", http.StatusBadRequest)
	}
	contentType := header.Get("Content-Type")
	if validator.IsEmpty(contentType) {
		return nil, nil, httpz.NewMsgError("content_type", http.StatusBadRequest)
	}
	contentLength := header.Get("Content-Length")
	if validator.IsEmpty(contentLength) {
		return nil, nil, httpz.NewMsgError("content_length", http.StatusBadRequest)
	}

	authorizaton := header.Get("Authorization")
	if validator.IsEmpty(authorizaton) {
		return nil, nil, httpz.NewMsgError("authorization", http.StatusBadRequest)
	}

	authorizaton, err := jwt.GetBearer(authorizaton)
	if validator.IsEmpty(authorizaton) {
		return nil, nil, httpz.NewMsgError("authorization", http.StatusBadRequest)
	}

	contentLength64, err := strconv.ParseInt(contentLength, 10, 64)
	if err != nil {
		return nil, nil, httpz.NewMsgError("invalid_content_length", http.StatusBadRequest)
	}

	metaApiArgs := header.Get("Meta-API-Args")
	if validator.IsEmpty(metaApiArgs) {
		return nil, nil, httpz.NewMsgError("meta_api_args", http.StatusBadRequest)
	}

	var args rest.MetaAPIArgs
	err = json.Unmarshal([]byte(metaApiArgs), &args)
	if err != nil {
		return nil, nil, httpz.NewMsgError("meta_api_args_not_json", http.StatusBadRequest)
	}

	if validator.IsEmpty(args.Filename) {
		return nil, nil, httpz.NewMsgError("filename", http.StatusBadRequest)
	}

	return &rest.UploadHeaders{
		ContentType: 	contentType,
		Authorization: authorizaton,
		ContentLength: 	contentLength64},
		&args, nil
}