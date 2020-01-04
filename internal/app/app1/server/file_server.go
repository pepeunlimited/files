package server

import (
	"context"
	"github.com/pepeunlimited/files/internal/app/app1/validator"
	"github.com/pepeunlimited/files/rpc"
)

type FileServer struct {
	validator validator.FileServerValidator
}

func (server FileServer) CreateFile(ctx context.Context, params *rpc.CreateFileParams) (*rpc.File, error) {
	//https://params.BucketName+region+endpoint+(fra1.digitaloceanspaces.com)+params.Filename
	//https://test-bucket-666.fra1.digitaloceanspaces.com/piia.txt
	return nil, nil
}

func NewFileServer() FileServer {
	return FileServer{validator: validator.NewFileServerValidator()}
}