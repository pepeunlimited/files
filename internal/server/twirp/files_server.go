package twirp

import (
	"context"
	"github.com/pepeunlimited/files/internal/pkg/ent"
	"github.com/pepeunlimited/files/internal/pkg/mysql/bucketrepo"
	"github.com/pepeunlimited/files/internal/pkg/mysql/filerepo"
	"github.com/pepeunlimited/files/internal/server/validator"
	"github.com/pepeunlimited/files/pkg/filesrpc"
	"github.com/pepeunlimited/files/pkg/storage"
	validator2 "github.com/pepeunlimited/microservice-kit/validator"
	"github.com/twitchtv/twirp"
	"log"
	"strings"
)

type FilesServer struct {
	validator validator.SpacesServerValidator
	spaces    bucketrepo.BucketRepository
	files     filerepo.FileRepository
	actions   storage.Actions // storage actions..
}

func (server FilesServer) CreateBucket(ctx context.Context, params *filesrpc.CreateBucketParams) (*filesrpc.CreateBucketResponse, error) {
	if err := server.validator.CreateBucket(params); err != nil {
		return nil, err
	}
	err := server.actions.CreateBucket(ctx, storage.Buckets{
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
	bucket, err := server.spaces.Create(ctx, params.Name, params.Endpoint, &cdn)
	if err != nil {
		server.actions.DeleteBucket(storage.Buckets{
			BucketName: params.Name,
			Endpoint:   params.Endpoint,
		})
		return nil, twirp.InternalErrorWith(err)
	}
	return &filesrpc.CreateBucketResponse{
		Endpoint:    storage.Endpoint,
		CdnEndpoint: *bucket.CdnEndpoint,
		Name:        bucket.Name,
		BucketId:    int64(bucket.ID),
	}, nil
}

func (server FilesServer) fileByFilename(ctx context.Context, params *filesrpc.Filename) (*ent.File, *ent.Bucket, error) {
	isDeleted := false
	if params.BucketName != nil && !validator2.IsEmpty(params.BucketName.Value) {
		return server.files.GetFileByFilenameBucketName(ctx, params.Name, params.BucketName.Value, nil, &isDeleted)
	}
	return server.files.GetFileByFilenameBucketID(ctx, params.Name, int(params.BucketId.Value), nil, &isDeleted)
}

func (server FilesServer) GetFile(ctx context.Context, params *filesrpc.GetFileParams) (*filesrpc.File, error) {
	if err := server.validator.GetFile(params); err != nil {
		return nil, err
	}
	var file *ent.File
	var err error
	var buckets *ent.Bucket
	isDraft := false
	isDeleted := false
	if params.Filename == nil {
		file, buckets, err = server.files.GetFilesBucketByID(ctx, int(params.FileId.Value), &isDraft, &isDeleted)
	} else {
		file, buckets, err = server.fileByFilename(ctx, params.Filename)
	}
	if err != nil {
		return nil, server.isFileError(err)
	}
	return toFile(file, buckets), nil
}

func (server FilesServer) isFileError(err error) error {
	switch err {
	case filerepo.ErrFileNotExist:
		return twirp.NotFoundError(filesrpc.FileNotFound)
	case bucketrepo.ErrBucketsNotExist:
		return twirp.NotFoundError(filesrpc.BucketNotFound)
	}
	log.Print("buckets-service: unknown: "+err.Error())
	return twirp.InternalErrorWith(err)
}

func (server FilesServer) Delete(ctx context.Context, params *filesrpc.DeleteParams) (*filesrpc.DeleteResponse, error) {
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
	file, buckets, err := server.files.GetFilesBucketByID(ctx, fileID, nil, nil)
	if err != nil {
		return nil, server.isFileError(err)
	}
	if params.IsPermanent {
		// call the actions for object delete..
		if err := server.actions.Delete(ctx, storage.Buckets{BucketName: buckets.Name, Endpoint:buckets.Endpoint}, file.Filename); err == nil {
			server.files.DeleteFileByID(ctx, fileID)
		}
	}
	return &filesrpc.DeleteResponse{}, nil
}

func (server FilesServer) GetBuckets(context.Context, *filesrpc.GetBucketsParams) (*filesrpc.GetBucketsResponse, error) {
	panic("implement me")
}

func (server FilesServer) Wipe(ctx context.Context, params *filesrpc.WipeParams) (*filesrpc.WipeParamsResponse, error) {
	panic("implement me")
}

func (server FilesServer) GetFiles(ctx context.Context, params *filesrpc.GetFilesParams) (*filesrpc.GetFilesResponse, error) {
	panic("implement me")
}

func (server FilesServer) Cut(ctx context.Context, params *filesrpc.CutParams) (*filesrpc.CutResponse, error) {
	panic("implement me")
}

func NewFilesServer(actions storage.Actions, client *ent.Client) FilesServer {
	return FilesServer{
		actions:   actions,
		validator: validator.NewSpacesServerValidator(),
		spaces:    bucketrepo.NewBucketRepository(client),
		files:     filerepo.NewFileRepository(client),
	}
}