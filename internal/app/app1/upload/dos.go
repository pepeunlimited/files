package upload

import (
	"errors"
	"github.com/pepeunlimited/files/spaces"
	"github.com/pepeunlimited/files/storage"
	"github.com/pepeunlimited/microservice-kit/misc"
)

var (
	ErrBucketExist 		= errors.New("spaces: bucket exist")
	ErrBucketNotExist 	= errors.New("spaces: bucket not found")
)

// DigitalOceanSpaces
type dos struct {
	accessKey		string
	secretKey		string
	accessToken     string
}

func (dos dos) DeleteBucket(bucket storage.Buckets) error {
	return spaces.NewBucket(bucket.Endpoint, dos.accessKey, dos.secretKey, bucket.BucketName).Delete()
}

func (dos dos) CreateBucket(bucket storage.Buckets) error {
	b := spaces.NewBucketCDN(bucket.Endpoint, dos.accessKey, dos.secretKey, bucket.BucketName, &dos.accessToken)
	exist, err := b.Exist()
	if err != nil {
		return err
	}
	if exist {
		return ErrBucketExist
	}
	return b.Create().Execute()
}

func (dos dos) Delete(bucket storage.Buckets, filename string) error {
	return spaces.NewBucket(bucket.Endpoint, dos.accessKey, dos.secretKey, bucket.BucketName).Files().Delete(filename).Execute()
}

func (dos dos) Upload(file storage.File, meta storage.FileMetaData, buckets storage.Buckets) error {
	return spaces.NewBucket(buckets.Endpoint, dos.accessKey, dos.secretKey, buckets.BucketName).Files().Create(file, meta).Execute()
}

func NewDos() storage.Actions {
	spacesAccessKey    := misc.GetEnv(spaces.SpacesAccessKey, spaces.AccessKey)
	spacesSecretKey    := misc.GetEnv(spaces.SpacesSecretKey, spaces.SecretKey)
	doAccessToken      := misc.GetEnv(spaces.DoAccessToken,   spaces.AccessToken)
	return dos{
		accessKey:   spacesAccessKey,
		secretKey:   spacesSecretKey,
		accessToken: doAccessToken,
	}
}