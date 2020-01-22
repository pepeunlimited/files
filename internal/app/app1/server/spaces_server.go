package server

import (
	"context"
	"github.com/pepeunlimited/files/internal/app/app1/ent"
	"github.com/pepeunlimited/files/internal/app/app1/filerepo"
	"github.com/pepeunlimited/files/internal/app/app1/spacesrepo"
	"github.com/pepeunlimited/files/internal/app/app1/validator"
	"github.com/pepeunlimited/files/spacesrpc"
	"github.com/pepeunlimited/files/storage"
	"github.com/pepeunlimited/microservice-kit/rpcz"
	validator2 "github.com/pepeunlimited/microservice-kit/validator"
	"github.com/twitchtv/twirp"
	"log"
	"strings"
)

type SpacesServer struct {
	validator validator.SpacesServerValidator
	spaces    spacesrepo.SpacesRepository
	files      filerepo.FileRepository
	actions   storage.Actions // storage actions..
}

func (server SpacesServer) CreateSpaces(ctx context.Context, params *spacesrpc.CreateSpacesParams) (*spacesrpc.CreateSpacesResponse, error) {
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
	spaces, err := server.spaces.Create(ctx, params.Name, params.Endpoint, &cdn)
	if err != nil {
		server.actions.DeleteBucket(storage.Buckets{
			BucketName: params.Name,
			Endpoint:   params.Endpoint,
		})
		return nil, twirp.InternalErrorWith(err)
	}
	return &spacesrpc.CreateSpacesResponse{
		Endpoint:    spaces.Endpoint,
		CdnEndpoint: *spaces.CdnEndpoint,
		Name:        spaces.Name,
		SpacesId:    int64(spaces.ID),
	}, nil
}

func (server SpacesServer) fileByFilename(ctx context.Context, params *spacesrpc.Filename) (*ent.Files, *ent.Spaces, error) {
	isDeleted := false
	if params.BucketName != nil && !validator2.IsEmpty(params.BucketName.Value) {
		return server.files.GetFileByFilenameSpacesName(ctx, params.Name, params.BucketName.Value, nil, &isDeleted)
	}
	return server.files.GetFileByFilenameSpacesID(ctx, params.Name, int(params.BucketId.Value), nil, &isDeleted)
}

func (server SpacesServer) GetFile(ctx context.Context, params *spacesrpc.GetFileParams) (*spacesrpc.File, error) {
	if err := server.validator.GetFile(params); err != nil {
		return nil, err
	}
	var file *ent.Files
	var err error
	var spaces *ent.Spaces
	isDraft := false
	isDeleted := false
	if params.Filename == nil {
		file, spaces, err = server.files.GetFilesSpacesByID(ctx, int(params.FileId.Value), &isDraft, &isDeleted)
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
	case filerepo.ErrFileNotExist:
		return twirp.NotFoundError("file not exist").WithMeta(rpcz.Reason, spacesrpc.FileNotFound)
	case spacesrepo.ErrSpacesNotExist:
		return twirp.NotFoundError("spaces not exist").WithMeta(rpcz.Reason, spacesrpc.SpacesNotFound)
	}
	log.Print("spaces-service: unknown: "+err.Error())
	return twirp.NewError(twirp.Internal ,"unknown: "+err.Error())
}

func (server SpacesServer) Delete(ctx context.Context, params *spacesrpc.DeleteParams) (*spacesrpc.DeleteResponse, error) {
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
		_, err := server.files.GetFileByID(ctx, fileID, nil, &isDeleted)
		if err != nil {
			return nil, server.isFileError(err)
		}
	}
	_, err := server.files.MarkAsDeletedByID(ctx, fileID)
	if err != nil {
		return nil, server.isFileError(err)
	}
	file, buckets, err := server.files.GetFilesSpacesByID(ctx, fileID, nil, nil)
	if err != nil {
		return nil, server.isFileError(err)
	}
	if params.IsPermanent {
		// call the actions for object delete..
		if err := server.actions.Delete(storage.Buckets{BucketName: buckets.Name, Endpoint:buckets.Endpoint}, file.Filename); err == nil {
			server.files.DeleteFileByID(ctx, fileID)
		}
	}
	return &spacesrpc.DeleteResponse{}, nil
}

func (server SpacesServer) GetSpaces(context.Context, *spacesrpc.GetSpacesParams) (*spacesrpc.GetSpacesResponse, error) {
	panic("implement me")
}

func (server SpacesServer) Wipe(ctx context.Context, params *spacesrpc.WipeParams) (*spacesrpc.WipeParamsResponse, error) {
	panic("implement me")
}

func (server SpacesServer) GetFiles(ctx context.Context, params *spacesrpc.GetFilesParams) (*spacesrpc.GetFilesResponse, error) {
	panic("implement me")
}

func (server SpacesServer) Cut(ctx context.Context, params *spacesrpc.CutParams) (*spacesrpc.CutResponse, error) {
	panic("implement me")
}

func NewSpacesServer(actions storage.Actions, client *ent.Client) SpacesServer {
	return SpacesServer{
		actions:   actions,
		validator: validator.NewSpacesServerValidator(),
		spaces:    spacesrepo.NewSpacesRepository(client),
		files:     filerepo.NewFileRepository(client),
	}
}