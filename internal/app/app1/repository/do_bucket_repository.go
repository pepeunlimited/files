package repository

import (
	"context"
	"errors"
	"github.com/pepeunlimited/files/internal/app/ent"
	"github.com/pepeunlimited/files/internal/app/ent/dobuckets"
	"time"
)

var (
	ErrDOBucketNotExist 	= errors.New("do-buckets: do-bucket not exist")
	ErrDOBucketExist 		= errors.New("do-buckets: do-bucket already exist")
)

type DOBucketRepository interface {
	CreateBucket(ctx context.Context, bucketName string, endpoint string, cdnEndpoint *string) (*ent.DoBuckets, error)
	GetBucketByName(ctx context.Context, bucketName string) (*ent.DoBuckets, error)
	GetBucketByID(ctx context.Context, id int) (*ent.DoBuckets, error)
	DeleteBucketByName(ctx context.Context, bucketName string) error
	DeleteAll(ctx context.Context)
}

type dobucketMySQL struct {
	client *ent.Client
}

func (b dobucketMySQL) GetBucketByID(ctx context.Context, id int) (*ent.DoBuckets, error) {
	bucket, err := b.client.DoBuckets.Get(ctx, id)
	if err != nil {
		return nil, b.isBucketError(err)
	}
	return bucket, nil
}

func (b dobucketMySQL) DeleteBucketByName(ctx context.Context, bucketName string) error {
	if _, err := b.GetBucketByName(ctx, bucketName); err != nil {
		return err
	}
	_, err := b.client.DoBuckets.Delete().Where(dobuckets.BucketName(bucketName)).Exec(ctx)
	return err
}

func (b dobucketMySQL) DeleteAll(ctx context.Context) {
	b.client.DoBuckets.Delete().ExecX(ctx)
}

func (b dobucketMySQL) GetBucketByName(ctx context.Context, bucketName string) (*ent.DoBuckets, error) {
	bucket, err := b.client.DoBuckets.Query().Where(dobuckets.BucketName(bucketName)).Only(ctx)
	if err != nil {
		return nil, b.isBucketError(err)
	}
	return bucket, nil
}

func (b dobucketMySQL) isBucketError(err error) error {
	if ent.IsNotFound(err) {
		return ErrDOBucketNotExist
	}
	// unknown
	return err
}


func (b dobucketMySQL) CreateBucket(ctx context.Context, bucketName string, endpoint string, cdnEndpoint *string) (*ent.DoBuckets, error) {
	if _, err := b.GetBucketByName(ctx, bucketName); err == nil {
		return nil, ErrDOBucketExist
	}
	return b.client.
		DoBuckets.
		Create().
		SetBucketName(bucketName).
		SetEndpoint(endpoint).
		SetNillableCdnEndpoint(cdnEndpoint).
		SetCreatedAt(time.Now().UTC()).
		Save(ctx)
}

func NewDOBucketRepository(client *ent.Client) DOBucketRepository {
	return &dobucketMySQL{client:client}
}