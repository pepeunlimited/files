package upload

import (
	"github.com/pepeunlimited/files/storage"
	"github.com/pepeunlimited/microservice-kit/errorz"
	"log"
)

type DosMock struct {
	IsUpload 		bool
	IsDelete 		bool
	IsCreateBucket 	bool
	IsDeleteBucket 	bool
	Errors  		errorz.Stack
	File     		storage.File
	Metadata 		storage.FileMetaData
	To 				storage.Buckets
}

func (d *DosMock) DeleteBucket(bucket storage.Buckets) error {
	log.Print("StorageActions: DeleteBucket..")
	d.IsDeleteBucket = true
	return d.Errors.Pop()
}

func (d *DosMock) CreateBucket(bucket storage.Buckets) error {
	log.Print("StorageActions: CreateBucket..")
	d.IsCreateBucket = true
	return d.Errors.Pop()
}

func (d *DosMock) Delete(bucket storage.Buckets, filename string) error {
	log.Print("StorageActions: Delete..")
	d.IsDelete = true
	return d.Errors.Pop()
}

func (d *DosMock) Upload(file storage.File, meta storage.FileMetaData, buckets storage.Buckets) error {
	log.Print("StorageActions: Upload..")
	d.IsUpload = true
	d.File = file
	d.Metadata = meta
	d.To = buckets
	log.Printf("File=%v, Meta=%v, To=%v",file, meta, buckets)
	return d.Errors.Pop()
}

func NewDosMock(errors []error) storage.Actions {
	return &DosMock{Errors:errorz.NewErrorStack(errors)}
}