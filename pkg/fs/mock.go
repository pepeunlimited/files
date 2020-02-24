package fs

import "context"

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
	panic("implement me")
}

func (f FileSystemMock) CreateBucket(ctx context.Context, bucket string) error {
	panic("implement me")
}

func (f FileSystemMock) BucketExist(bucket string) (bool, error) {
	panic("implement me")
}

func (f FileSystemMock) UploadFile(ctx context.Context, file File, meta FileMetaData, bucket string) error {
	panic("implement me")
}

func (f FileSystemMock) DeleteFile(filename string, bucket string) error {
	panic("implement me")
}

func (f FileSystemMock) UpdateFile(file File, meta FileMetaData) (*File, error) {
	panic("implement me")
}

func (f FileSystemMock) GetFile(filename string, bucket string) ([]byte, error) {
	panic("implement me")
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