package upload

import (
	"context"
	"github.com/pepeunlimited/files/pkg/storage"
	"github.com/pepeunlimited/microservice-kit/errorz"
	"log"
)

type ActionsMock struct {
	IsUpload       bool
	IsDelete       bool
	IsCreateBucket bool
	IsDeleteBucket bool
	Errors         errorz.Stack
	File           storage.File
	Metadata       storage.FileMetaData
	To             storage.Buckets
}

func (d *ActionsMock) DeleteBucket(bucket storage.Buckets) error {
	log.Print("StorageActions: DeleteBucket..")
	d.IsDeleteBucket = true
	return d.Errors.Pop()
}

func (d *ActionsMock) CreateBucket(ctx context.Context, bucket storage.Buckets) error {
	log.Print("StorageActions: CreateBucket..")
	d.IsCreateBucket = true
	return d.Errors.Pop()
}

func (d *ActionsMock) Delete(ctx context.Context, bucket storage.Buckets, filename string) error {
	log.Print("StorageActions: Delete..")
	d.IsDelete = true
	return d.Errors.Pop()
}

func (d *ActionsMock) Upload(ctx context.Context, file storage.File, meta storage.FileMetaData, buckets storage.Buckets) error {
	log.Print("StorageActions: Upload..")
	d.IsUpload = true
	d.File = file
	d.Metadata = meta
	d.To = buckets
	log.Printf("File=%v, Meta=%v, To=%v",file, meta, buckets)
	return d.Errors.Pop()
}

func NewActionsMock(errors []error) storage.Actions {
	return &ActionsMock{Errors: errorz.NewErrorStack(errors)}
}