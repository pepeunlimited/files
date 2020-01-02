package server

import (
	"context"
	"github.com/pepeunlimited/images/internal/app/app1/validator"
	"github.com/pepeunlimited/images/rpc"
)

type ImageServer struct {
	validator validator.ImageServerValidator
}

func (server ImageServer) CreateImage(ctx context.Context, params *rpc.CreateImageParams) (*rpc.Image, error) {


	return nil, nil
}

func NewImageServer() ImageServer {
	return ImageServer{
		validator: validator.NewImageServerValidator(),
	}
}