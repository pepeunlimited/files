package bucketsrepo

import (
	"context"
	"errors"
	"github.com/pepeunlimited/files/internal/pkg/ent"
	"github.com/pepeunlimited/files/internal/pkg/ent/buckets"
	"github.com/pepeunlimited/files/internal/pkg/ent/files"
	"time"
)

var (
	ErrBucketsNotExist = errors.New("buckets: not exist")
	ErrBucketsExist    = errors.New("buckets: already exist")
)

// one-to-many `files`
type BucketsRepository interface {
	Create(ctx context.Context, name string, endpoint string, cdnEndpoint *string) (*ent.Buckets, error)
	GetBucketsByName(ctx context.Context, name string) 				(*ent.Buckets, error)
	GetBucketByID(ctx context.Context, id int) 						(*ent.Buckets, error)
	DeleteBucketByName(ctx context.Context, name string) 			error
	GetBuckets(ctx context.Context, pageToken int64, pageSize int32) ([]*ent.Buckets, int64, error)

	// wipes spaces and files
	Wipe(ctx context.Context)
}

type bucketsMySQL struct {
	client *ent.Client
}
//AND id > ? ORDER BY id ASC LIMIT ?
func (b bucketsMySQL) GetBuckets(ctx context.Context, pageToken int64, pageSize int32) ([]*ent.Buckets, int64, error) {
	buckets, err := b.client.Buckets.Query().Where(buckets.IDGT(int(pageToken))).Order(ent.Asc(files.FieldID)).Limit(int(pageSize)).All(ctx)
	if err != nil {
		return nil, 0, err
	}
	if len(buckets) == 0 {
		return []*ent.Buckets{}, pageToken, nil
	}
	return buckets, int64(buckets[len(buckets) - 1].ID), nil
}

func (b bucketsMySQL) GetBucketByID(ctx context.Context, id int) (*ent.Buckets, error) {
	bucket, err := b.client.Buckets.Get(ctx, id)
	if err != nil {
		return nil, b.isSpacesError(err)
	}
	return bucket, nil
}

func (b bucketsMySQL) DeleteBucketByName(ctx context.Context, name string) error {
	if _, err := b.GetBucketsByName(ctx, name); err != nil {
		return err
	}
	_, err := b.client.Buckets.Delete().Where(buckets.Name(name)).Exec(ctx)
	return err
}

func (b bucketsMySQL) Wipe(ctx context.Context) {
	b.client.Files.Delete().ExecX(ctx)
	b.client.Buckets.Delete().ExecX(ctx)
}

func (b bucketsMySQL) GetBucketsByName(ctx context.Context, bucketName string) (*ent.Buckets, error) {
	bucket, err := b.client.Buckets.Query().Where(buckets.Name(bucketName)).Only(ctx)
	if err != nil {
		return nil, b.isSpacesError(err)
	}
	return bucket, nil
}

func (b bucketsMySQL) isSpacesError(err error) error {
	if ent.IsNotFound(err) {
		return ErrBucketsNotExist
	}
	// unknown
	return err
}

func (b bucketsMySQL) Create(ctx context.Context, name string, endpoint string, cdnEndpoint *string) (*ent.Buckets, error) {
	if _, err := b.GetBucketsByName(ctx, name); err == nil {
		return nil, ErrBucketsExist
	}
	return b.client.
		Buckets.
		Create().
		SetName(name).
		SetEndpoint(endpoint).
		SetNillableCdnEndpoint(cdnEndpoint).
		SetCreatedAt(time.Now().UTC()).
		Save(ctx)
}

func NewBucketsRepository(client *ent.Client) BucketsRepository {
	return &bucketsMySQL{client: client}
}