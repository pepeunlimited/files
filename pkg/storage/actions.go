package storage

import "context"

type Actions interface {
	Delete(ctx context.Context, bucket Buckets, filename string) error
	Upload(ctx context.Context, file File, meta FileMetaData, buckets Buckets) error
	CreateBucket(ctx context.Context, bucket Buckets) error
	DeleteBucket(bucket Buckets) error
}