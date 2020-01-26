package upload

import (
	"context"
	"github.com/pepeunlimited/files/spaces"
	"github.com/pepeunlimited/files/storage"
	"github.com/pepeunlimited/microservice-kit/misc"
)


// DigitalOceanSpaces
type dos struct {
	accessKey		string
	secretKey		string
	accessToken     string
}

func (dos dos) DeleteBucket(bucket storage.Buckets) error {
	return storage.NewSpacesBuilder(bucket.Endpoint, dos.accessKey, dos.secretKey, bucket.BucketName, nil).Delete()
}

func (dos dos) CreateBucket(ctx context.Context, bucket storage.Buckets) error {
	b := storage.NewSpacesBuilder(bucket.Endpoint, dos.accessKey, dos.secretKey, bucket.BucketName, &dos.accessToken)
	exist, err := b.Exist()
	if err != nil {
		return err
	}
	if exist {
		return ErrBucketExist
	}
	return b.Create().Execute(ctx)
}

func (dos dos) Delete(ctx context.Context, bucket storage.Buckets, filename string) error {
	return storage.NewSpacesBuilder(bucket.Endpoint, dos.accessKey, dos.secretKey, bucket.BucketName, nil).Files().Delete(filename).Execute(ctx)
}

func (dos dos) Upload(ctx context.Context, file storage.File, meta storage.FileMetaData, buckets storage.Buckets) error {
	return storage.NewSpacesBuilder(buckets.Endpoint, dos.accessKey, dos.secretKey, buckets.BucketName, nil).Files().Create(file, meta).Execute(ctx)
}

func NewDos() storage.Actions {
	spacesAccessKey    := misc.GetEnv(storage.SpacesAccessKey, spaces.AccessKey)
	spacesSecretKey    := misc.GetEnv(storage.SpacesSecretKey, spaces.SecretKey)
	doAccessToken      := misc.GetEnv(storage.DoAccessToken,   spaces.AccessToken)
	return dos{
		accessKey:   spacesAccessKey,
		secretKey:   spacesSecretKey,
		accessToken: doAccessToken,
	}
}