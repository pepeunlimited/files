package bucket

import (
	"context"
	"errors"
	"github.com/pepeunlimited/files/internal/pkg/ent"
	"github.com/pepeunlimited/files/internal/pkg/ent/bucket"
	"github.com/pepeunlimited/files/internal/pkg/ent/file"
	"time"
)

var (
	ErrBucketsNotExist = errors.New("buckets: not exist")
	ErrBucketsExist    = errors.New("buckets: already exist")
)

// one-to-many `files`
type BucketRepository interface {
	Create(ctx context.Context, name string, endpoint string, cdnEndpoint *string) (*ent.Bucket, error)
	GetBucketsByName(ctx context.Context, name string) 				(*ent.Bucket, error)
	GetBucketsByEndpoint(ctx context.Context, endpoint string) 			(*ent.Bucket, error)
	GetBucketByID(ctx context.Context, id int) 						(*ent.Bucket, error)
	DeleteBucketByName(ctx context.Context, name string) 			error
	GetBuckets(ctx context.Context, pageToken int64, pageSize int32) ([]*ent.Bucket, int64, error)

	// wipes spaces and files
	Wipe(ctx context.Context)
}

type bucketMySQL struct {
	client *ent.Client
}

func (b bucketMySQL) GetBucketsByEndpoint(ctx context.Context, endpoint string) (*ent.Bucket, error) {
	only, err := b.client.Bucket.Query().Where(bucket.Endpoint(endpoint)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrBucketsNotExist
		}
		return nil, err
	}
	return only, nil
}

//AND id > ? ORDER BY id ASC LIMIT ?
func (b bucketMySQL) GetBuckets(ctx context.Context, pageToken int64, pageSize int32) ([]*ent.Bucket, int64, error) {
	buckets, err := b.client.Bucket.Query().Where(bucket.IDGT(int(pageToken))).Order(ent.Asc(file.FieldID)).Limit(int(pageSize)).All(ctx)
	if err != nil {
		return nil, 0, err
	}
	if len(buckets) == 0 {
		return []*ent.Bucket{}, pageToken, nil
	}
	return buckets, int64(buckets[len(buckets) - 1].ID), nil
}

func (b bucketMySQL) GetBucketByID(ctx context.Context, id int) (*ent.Bucket, error) {
	bucket, err := b.client.Bucket.Get(ctx, id)
	if err != nil {
		return nil, b.isSpacesError(err)
	}
	return bucket, nil
}

func (b bucketMySQL) DeleteBucketByName(ctx context.Context, name string) error {
	if _, err := b.GetBucketsByName(ctx, name); err != nil {
		return err
	}
	_, err := b.client.Bucket.Delete().Where(bucket.Name(name)).Exec(ctx)
	return err
}

func (b bucketMySQL) Wipe(ctx context.Context) {
	b.client.File.Delete().ExecX(ctx)
	b.client.Bucket.Delete().ExecX(ctx)
}

func (b bucketMySQL) GetBucketsByName(ctx context.Context, bucketName string) (*ent.Bucket, error) {
	bucket, err := b.client.Bucket.Query().Where(bucket.Name(bucketName)).Only(ctx)
	if err != nil {
		return nil, b.isSpacesError(err)
	}
	return bucket, nil
}

func (b bucketMySQL) isSpacesError(err error) error {
	if ent.IsNotFound(err) {
		return ErrBucketsNotExist
	}
	// unknown
	return err
}

func (b bucketMySQL) Create(ctx context.Context, name string, endpoint string, cdnEndpoint *string) (*ent.Bucket, error) {
	if _, err := b.GetBucketsByName(ctx, name); err == nil {
		return nil, ErrBucketsExist
	}
	return b.client.
		Bucket.
		Create().
		SetName(name).
		SetEndpoint(endpoint).
		SetNillableCdnEndpoint(cdnEndpoint).
		SetCreatedAt(time.Now().UTC()).
		Save(ctx)
}

func New(client *ent.Client) BucketRepository {
	return &bucketMySQL{client: client}
}