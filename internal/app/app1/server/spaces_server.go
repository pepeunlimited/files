package server

import (
	"context"
	"github.com/pepeunlimited/files/internal/app/app1/ent"
	"github.com/pepeunlimited/files/internal/app/app1/repository"
	"github.com/pepeunlimited/files/internal/app/app1/validator"
	"github.com/pepeunlimited/files/rpc"
	"github.com/pepeunlimited/files/storage"
	"github.com/pepeunlimited/microservice-kit/rpcz"
	validator2 "github.com/pepeunlimited/microservice-kit/validator"
	"github.com/twitchtv/twirp"
	"log"
	"strings"
)

type SpacesServer struct {
	validator 		validator.SpacesServerValidator
	spacesRepo 		repository.SpacesRepository
	filesRepo        repository.FileRepository
	actions         storage.Actions // storage actions..
}

func (server SpacesServer) CreateSpaces(ctx context.Context, params *rpc.CreateSpacesParams) (*rpc.CreateSpacesResponse, error) {
	if err := server.validator.CreateBucket(params); err != nil {
		return nil, err
	}
	err := server.actions.CreateBucket(storage.Buckets{
		BucketName: params.Name,
		Endpoint:   params.Endpoint,
	})

	endpoint := strings.Split(params.Endpoint, ".")
	if len(endpoint) != 3 {
		return nil, twirp.InvalidArgumentError("endpoint", "invalid_endpoint")
	}
	cdn := params.Name+"."+endpoint[0]+".cdn."+endpoint[1]+"."+endpoint[2]

	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}
	spaces, err := server.spacesRepo.Create(ctx, params.Name, params.Endpoint, &cdn)
	if err != nil {
		server.actions.DeleteBucket(storage.Buckets{
			BucketName: params.Name,
			Endpoint:   params.Endpoint,
		})
		return nil, twirp.InternalErrorWith(err)
	}
	return &rpc.CreateSpacesResponse{
		Endpoint:    spaces.Endpoint,
		CdnEndpoint: *spaces.CdnEndpoint,
		Name:        spaces.Name,
		SpacesId:    int64(spaces.ID),
	}, nil
}

func (server SpacesServer) fileByFilename(ctx context.Context, params *rpc.Filename) (*ent.Files, *ent.Spaces, error) {
	isDeleted := false
	if params.BucketName != nil && !validator2.IsEmpty(params.BucketName.Value) {
		return server.filesRepo.GetFileByFilenameSpacesName(ctx, params.Name, params.BucketName.Value, nil, &isDeleted)
	}
	return server.filesRepo.GetFileByFilenameSpacesID(ctx, params.Name, int(params.BucketId.Value), nil, &isDeleted)
}

func (server SpacesServer) GetFile(ctx context.Context, params *rpc.GetFileParams) (*rpc.File, error) {
	if err := server.validator.GetFile(params); err != nil {
		return nil, err
	}
	var file *ent.Files
	var err error
	var spaces *ent.Spaces
	isDraft := false
	isDeleted := false
	if params.Filename == nil {
		file, spaces, err = server.filesRepo.GetFilesSpacesByID(ctx, int(params.FileId.Value), &isDraft, &isDeleted)
	} else {
		file, spaces, err = server.fileByFilename(ctx, params.Filename)
	}

	if err != nil {
		return nil, server.isFileError(err)
	}

	return toFile(file, spaces), nil
}

func (server SpacesServer) isFileError(err error) error {
	switch err {
	case repository.ErrFileNotExist:
		return twirp.NotFoundError("file not exist").WithMeta(rpcz.Reason, rpc.FileNotFound)
	case repository.ErrSpacesNotExist:
		return twirp.NotFoundError("spaces not exist").WithMeta(rpcz.Reason, rpc.SpacesNotFound)
	}
	log.Print("spaces-service: unknown: "+err.Error())
	return twirp.NewError(twirp.Internal ,"unknown: "+err.Error())
}

func (server SpacesServer) Delete(ctx context.Context, params *rpc.DeleteParams) (*rpc.DeleteResponse, error) {
	if err := server.validator.Delete(params); err != nil {
		return nil, err
	}
	var fileID int
	if params.Filename != nil {
		file,_, err := server.fileByFilename(ctx, params.Filename)
		if err != nil {
			return nil, server.isFileError(err)
		}
		fileID = file.ID
	} else {
		fileID = int(params.FileId.Value)
	}
	isDeleted := false
	if !params.IsPermanent {
		_, err := server.filesRepo.GetFileByID(ctx, fileID, nil, &isDeleted)
		if err != nil {
			return nil, server.isFileError(err)
		}
	}
	_, err := server.filesRepo.MarkAsDeletedByID(ctx, fileID)
	if err != nil {
		return nil, server.isFileError(err)
	}
	file, buckets, err := server.filesRepo.GetFilesSpacesByID(ctx, fileID, nil, nil)
	if err != nil {
		return nil, server.isFileError(err)
	}
	if params.IsPermanent {
		// call the actions for object delete..
		if err := server.actions.Delete(storage.Buckets{BucketName: buckets.Name, Endpoint:buckets.Endpoint}, file.Filename); err == nil {
			server.filesRepo.DeleteFileByID(ctx, fileID)
		}
	}
	return &rpc.DeleteResponse{}, nil
}

func (server SpacesServer) GetSpaces(context.Context, *rpc.GetSpacesParams) (*rpc.GetSpacesResponse, error) {
	panic("implement me")
}

func (server SpacesServer) Wipe(ctx context.Context, params *rpc.WipeParams) (*rpc.WipeParamsResponse, error) {
	panic("implement me")
}

func (server SpacesServer) GetFiles(ctx context.Context, params *rpc.GetFilesParams) (*rpc.GetFilesResponse, error) {
	panic("implement me")
}

func (server SpacesServer) Cut(ctx context.Context, params *rpc.CutParams) (*rpc.CutResponse, error) {
	panic("implement me")
}

func NewSpacesServer(actions storage.Actions, client *ent.Client) SpacesServer {
	return SpacesServer{
		actions:		 actions,
		validator: 		 validator.NewSpacesServerValidator(),
		spacesRepo: 	 repository.NewSpacesRepository(client),
		filesRepo:  		 repository.NewFileRepository(client),
	}
}