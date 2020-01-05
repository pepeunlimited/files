package spaces

import (
	"github.com/minio/minio-go"
)

type Bucket interface {
	Delete(bucketName string) error
	Files(bucketName string)  Files
	Create(bucketName string) Files
	Exist(bucketName string)  (*bool, error)
}

type bucket struct {
	name 		string
	isCreate 	bool
	isFiles    	bool
	spaces 		spaces
}

func (b *bucket) Exist(bucketName string) (*bool, error) {
	_, err := b.spaces.client.GetBucketLocation(bucketName)
	exist := false
	if err != nil {
		minioerr, ok := err.(minio.ErrorResponse)
		if !ok {
			return nil, err
		}
		if minioerr.Code == "NoSuchBucket" {
			return &exist, nil
		}
		return nil, err
	}
	exist = true
	return &exist, nil
}

func (b *bucket) Delete(bucketName string) error {
	return b.spaces.client.RemoveBucket(bucketName)
}

func (b *bucket) Files(bucketName string) Files {
	executor := make(map[int]interface{})
	b.isFiles = true
	b.isCreate = false
	b.name = bucketName
	return &files{bucket:*b, executor:executor, order: 0}
}

func (b *bucket) Create(bucketName string) Files {
	executor := make(map[int]interface{})
	b.isCreate = true
	b.isFiles = false
	b.name = bucketName
	return &files{bucket:*b, executor:executor, order: 0}
}

