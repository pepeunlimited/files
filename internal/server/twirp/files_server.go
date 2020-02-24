package twirp

import (
	"context"
	"github.com/pepeunlimited/files/internal/pkg/ent"
	"github.com/pepeunlimited/files/internal/pkg/mysql/bucket"
	"github.com/pepeunlimited/files/internal/pkg/mysql/file"
	"github.com/pepeunlimited/files/internal/server/errorz"
	"github.com/pepeunlimited/files/internal/server/validator"
	"github.com/pepeunlimited/files/pkg/fs"
	"github.com/pepeunlimited/files/pkg/rpc/files"
	validator2 "github.com/pepeunlimited/microservice-kit/validator"
	"log"
)

type FilesServer struct {
	validator validator.SpacesServerValidator
	buckets   bucket.BucketRepository
	files     file.FileRepository
	fs        fs.FileSystem
}

func (server FilesServer) CreateBucket(ctx context.Context, params *files.CreateBucketParams) (*files.CreateBucketResponse, error) {
	if err := server.validator.CreateBucket(params); err != nil {
		return nil, err
	}
	if err := server.fs.CreateBucket(ctx, params.Name); err != nil {
		return nil, errorz.Fs(err)
	}
	cdn, err := server.fs.CdnEndpoint(params.Name)
	if err != nil {
		return nil, err
	}
	bucket, err := server.buckets.Create(ctx, params.Name, server.fs.Endpoint(), &cdn)
	if err != nil {
		if err := server.fs.DeleteBucket(params.Name); err != nil {
			log.Print("failed to delete bucket: "+ err.Error())
		}
		return nil, errorz.File(err)
	}
	return &files.CreateBucketResponse{
		Endpoint:    server.fs.Endpoint(),
		CdnEndpoint: *bucket.CdnEndpoint,
		Name:        bucket.Name,
		BucketId:    int64(bucket.ID),
	}, nil
}

func (server FilesServer) fileByFilename(ctx context.Context, params *files.Filename) (*ent.File, *ent.Bucket, error) {
	isDeleted := false
	if params.BucketName != nil && !validator2.IsEmpty(params.BucketName.Value) {
		return server.files.GetFileByFilenameBucketName(ctx, params.Name, params.BucketName.Value, nil, &isDeleted)
	}
	return server.files.GetFileByFilenameBucketID(ctx, params.Name, int(params.BucketId.Value), nil, &isDeleted)
}

func (server FilesServer) GetFile(ctx context.Context, params *files.GetFileParams) (*files.File, error) {
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
		return nil, errorz.File(err)
	}
	return toFile(file, buckets), nil
}

func (server FilesServer) Delete(ctx context.Context, params *files.DeleteParams) (*files.DeleteResponse, error) {
	if err := server.validator.Delete(params); err != nil {
		return nil, err
	}
	var fileID int
	if params.Filename != nil {
		file,_, err := server.fileByFilename(ctx, params.Filename)
		if err != nil {
			return nil, errorz.File(err)
		}
		fileID = file.ID
	} else {
		fileID = int(params.FileId.Value)
	}
	isDeleted := false
	if !params.IsPermanent {
		_, err := server.files.GetFileByID(ctx, fileID, nil, &isDeleted)
		if err != nil {
			return nil, errorz.File(err)
		}
	}
	_, err := server.files.MarkAsDeletedByID(ctx, fileID)
	if err != nil {
		return nil, errorz.File(err)
	}
	file, buckets, err := server.files.GetFilesBucketByID(ctx, fileID, nil, nil)
	if err != nil {
		return nil, errorz.File(err)
	}
	if params.IsPermanent {
		err := server.fs.DeleteFile(file.Filename, buckets.Name)
		if err != nil {
			return nil, errorz.Fs(err)
		}
		err = server.files.DeleteFileByID(ctx, fileID)
		if err != nil {
			return nil, errorz.File(err)
		}
	}
	return &files.DeleteResponse{}, nil
}

func (server FilesServer) GetBuckets(context.Context, *files.GetBucketsParams) (*files.GetBucketsResponse, error) {
	panic("implement me")
}

func (server FilesServer) Wipe(ctx context.Context, params *files.WipeParams) (*files.WipeParamsResponse, error) {
	panic("implement me")
}

func (server FilesServer) GetFiles(ctx context.Context, params *files.GetFilesParams) (*files.GetFilesResponse, error) {
	panic("implement me")
}

func (server FilesServer) Cut(ctx context.Context, params *files.CutParams) (*files.CutResponse, error) {
	panic("implement me")
}

func NewFilesServer(client *ent.Client, fs fs.FileSystem) FilesServer {
	return FilesServer{
		validator: validator.NewSpacesServerValidator(),
		buckets:   bucket.New(client),
		files:     file.NewFileRepository(client),
		fs:        fs,
	}
}