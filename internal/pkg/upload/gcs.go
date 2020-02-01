package upload

import (
	"context"
	"github.com/pepeunlimited/files/pkg/storage"
	"github.com/pepeunlimited/microservice-kit/misc"
)

type gcs struct {
	accessKey		string
	secretKey		string
}

func (gcs gcs) DeleteBucket(bucket storage.Buckets) error {
	return storage.NewGoogleCloudStorageBuilder(bucket.Endpoint, gcs.accessKey, gcs.secretKey, bucket.BucketName).Delete()
}

func (gcs gcs) CreateBucket(ctx context.Context, bucket storage.Buckets) error {
	b := storage.NewGoogleCloudStorageBuilder(bucket.Endpoint, gcs.accessKey, gcs.secretKey, bucket.BucketName)
	exist, err := b.Exist()
	if err != nil {
		return err
	}
	if exist {
		return ErrBucketExist
	}
	return b.Create().Execute(ctx)
}

func (gcs gcs) Delete(ctx context.Context, bucket storage.Buckets, filename string) error {
	return storage.NewGoogleCloudStorageBuilder(bucket.Endpoint, gcs.accessKey, gcs.secretKey, bucket.BucketName).Files().Delete(filename).Execute(ctx)
}

func (gcs gcs) Upload(ctx context.Context, file storage.File, meta storage.FileMetaData, buckets storage.Buckets) error {
	return storage.NewGoogleCloudStorageBuilder(buckets.Endpoint, gcs.accessKey, gcs.secretKey, buckets.BucketName).Files().Create(file, meta).Execute(ctx)
}

func NewGoogleCloudStorage() storage.Actions {
	accessKey    := misc.GetEnv(storage.GoogleCloudStorageAccessKey, storage.AccessKey)
	secretKey    := misc.GetEnv(storage.GoogleCloudStorageSecretKey, storage.SecretKey)
	return gcs{
		accessKey: accessKey,
		secretKey:secretKey,
	}
}