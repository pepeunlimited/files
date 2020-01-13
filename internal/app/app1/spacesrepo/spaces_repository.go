package spacesrepo

import (
	"context"
	"errors"
	"github.com/pepeunlimited/files/internal/app/app1/ent"
	"github.com/pepeunlimited/files/internal/app/app1/ent/spaces"
	"time"
)

var (
	ErrSpacesNotExist 	= errors.New("spaces: not exist")
	ErrSpacesExist 		= errors.New("spaces: already exist")
)

// access `spaces`
// one-to-many `files`
type SpacesRepository interface {
	Create(ctx context.Context, name string, endpoint string, cdnEndpoint *string) (*ent.Spaces, error)
	GetSpaceByName(ctx context.Context, name string) 				(*ent.Spaces, error)
	GetSpaceByID(ctx context.Context, id int) 						(*ent.Spaces, error)
	DeleteSpaceByName(ctx context.Context, name string) 			error
	GetSpaces(ctx context.Context, pageToken int64, pageSize int32) ([]*ent.Spaces, int64, error)

	// wipes spaces and files
	Wipe(ctx context.Context)
}

type spacesMySQL struct {
	client *ent.Client
}
//AND id > ? ORDER BY id ASC LIMIT ?
func (b spacesMySQL) GetSpaces(ctx context.Context, pageToken int64, pageSize int32) ([]*ent.Spaces, int64, error) {
	buckets, err := b.client.Spaces.Query().Where(spaces.IDGT(int(pageToken))).Order(ent.Asc(spaces.FieldID)).Limit(int(pageSize)).All(ctx)
	if err != nil {
		return nil, 0, err
	}
	if len(buckets) == 0 {
		return []*ent.Spaces{}, pageToken, nil
	}
	return buckets, int64(buckets[len(buckets) - 1].ID), nil
}

func (b spacesMySQL) GetSpaceByID(ctx context.Context, id int) (*ent.Spaces, error) {
	bucket, err := b.client.Spaces.Get(ctx, id)
	if err != nil {
		return nil, b.isSpacesError(err)
	}
	return bucket, nil
}

func (b spacesMySQL) DeleteSpaceByName(ctx context.Context, name string) error {
	if _, err := b.GetSpaceByName(ctx, name); err != nil {
		return err
	}
	_, err := b.client.Spaces.Delete().Where(spaces.Name(name)).Exec(ctx)
	return err
}

func (b spacesMySQL) Wipe(ctx context.Context) {
	b.client.Files.Delete().ExecX(ctx)
	b.client.Spaces.Delete().ExecX(ctx)
}

func (b spacesMySQL) GetSpaceByName(ctx context.Context, bucketName string) (*ent.Spaces, error) {
	bucket, err := b.client.Spaces.Query().Where(spaces.Name(bucketName)).Only(ctx)
	if err != nil {
		return nil, b.isSpacesError(err)
	}
	return bucket, nil
}

func (b spacesMySQL) isSpacesError(err error) error {
	if ent.IsNotFound(err) {
		return ErrSpacesNotExist
	}
	// unknown
	return err
}

func (b spacesMySQL) Create(ctx context.Context, name string, endpoint string, cdnEndpoint *string) (*ent.Spaces, error) {
	if _, err := b.GetSpaceByName(ctx, name); err == nil {
		return nil, ErrSpacesExist
	}
	return b.client.
		Spaces.
		Create().
		SetName(name).
		SetEndpoint(endpoint).
		SetNillableCdnEndpoint(cdnEndpoint).
		SetCreatedAt(time.Now().UTC()).
		Save(ctx)
}

func NewSpacesRepository(client *ent.Client) SpacesRepository {
	return &spacesMySQL{client:client}
}