package fs

import (
	"context"
	"log"
)

type FileSystemMock struct {
	IsDeleteBucket bool
	IsCreateBucket bool
	IsBucketExist  bool
	IsCreateFile   bool
	IsDeleteFile   bool
	IsUpdateFile   bool
	IsGetFile	   bool
	IsGetMetadata  bool

	IsBucketError bool
	IsFileError   bool
	Bucket 		 *Bucket
	File   		 *File
	FileMetaData *FileMetaData
	endpoint 	  string
}

func (f FileSystemMock) CdnEndpoint(bucket string) (string, error) {
	return f.Endpoint(), nil
}

func (f FileSystemMock) Endpoint() string {
	return f.endpoint
}

func (f FileSystemMock) DeleteBucket(bucket string) error {
	log.Print("deleting bucket...")
	return nil
}

func (f FileSystemMock) CreateBucket(ctx context.Context, bucket string) error {
	log.Print("creating bucket..")
	return nil
}

func (f FileSystemMock) BucketExist(bucket string) (bool, error) {
	return false, nil
}

func (f FileSystemMock) UploadFile(ctx context.Context, file File, meta FileMetaData, bucket string) error {
	log.Print("uploading file..")
	return nil
}

func (f FileSystemMock) DeleteFile(filename string, bucket string) error {
	log.Print("deleting file..")
	return nil
}

func (f FileSystemMock) UpdateFile(file File, meta FileMetaData) (*File, error) {
	panic("implement me")
}

func (f FileSystemMock) GetFile(filename string, bucket string) ([]byte, error) {
	log.Print("getting file..")
	return []byte("asd"), nil
}

func (f FileSystemMock) GetMetadata(filename string, bucket string) (*FileMetaData, error) {
	panic("implement me")
}

func (f FileSystemMock) GetBuckets() ([]*Bucket, error) {
	panic("implement me")
}

func NewMock(endpoint string) FileSystemMock {
	return FileSystemMock{endpoint:endpoint}
}