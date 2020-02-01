package filerepo

import (
	"context"
	"errors"
	"github.com/pepeunlimited/files/internal/pkg/ent"
	"github.com/pepeunlimited/files/internal/pkg/ent/buckets"
	"github.com/pepeunlimited/files/internal/pkg/ent/files"
	"github.com/pepeunlimited/files/internal/pkg/mysql/bucketsrepo"
	"time"
)

var (
	ErrFileExist 		= errors.New("files: already exist")
	ErrFileNotExist 	= errors.New("files: not exist")
)

// many-to-one do_buckets
type FileRepository interface {
	CreateFile(ctx context.Context, filename string, fileSize int64, mimeType string, isDeleted bool, isDraft bool, userId int64, bucketID int) (*ent.Files, error)

	GetFilesBucketByID(ctx context.Context, fileID int, isDraft *bool, isDeleted *bool)		(*ent.Files, *ent.Buckets, error)
 	GetFileByID(ctx context.Context, fileID int, isDraft *bool, isDeleted *bool)				(*ent.Files, error)

	MarkAsDeletedByID(ctx context.Context, fileID int)										(*ent.Files, error)
	MarkAsDraftByID(ctx context.Context, fileID int)											(*ent.Files, error)

	// spacesID or spacesName is required because same filename may be in another spaces..
	GetFileByFilenameBucketID(ctx context.Context, filename string, bucketsID int, isDraft *bool, isDeleted *bool)	(*ent.Files, *ent.Buckets, error)
	GetFileByFilenameBucketName(ctx context.Context, filename string, bucketsName string, isDraft *bool, isDeleted *bool) (*ent.Files, *ent.Buckets, error)

	DeleteFileByID(ctx context.Context, fileID int) error
	ExistInBucket(ctx context.Context, filename string, bucketsID int) (*bool, error)
}

type filesMySQL struct {
	client  *ent.Client
}

func (f filesMySQL) ExistInBucket(ctx context.Context, filename string, spacesID int) (*bool, error) {
	_, _, err := f.GetFileByFilenameBucketID(ctx, filename, spacesID, nil, nil)
	exist := false
	if err != nil {
		if err == ErrFileNotExist {
			return &exist, nil
		}
		return nil, err
	}
	exist = true
	return &exist, nil
}

func (f filesMySQL) GetFilesBucketByID(ctx context.Context, fileID int, isDraft *bool, isDeleted *bool) (*ent.Files, *ent.Buckets, error) {
	file, err := f.GetFileByID(ctx, fileID, isDraft, isDeleted)
	if err != nil {
		return nil, nil, err
	}
	bucket, err := file.QueryBuckets().Only(ctx)
	if err != nil {
		return nil, nil, f.isSpacesError(err)
	}
	return file, bucket, nil
}

func (f filesMySQL) GetFileByID(ctx context.Context, fileID int, isDraft *bool, isDeleted *bool) (*ent.Files, error) {
	query := f.client.Files.Query().Where(files.ID(fileID))
	if isDraft != nil {
		query.Where(files.IsDraft(*isDraft))
	}
	if isDeleted != nil {
		query.Where(files.IsDeleted(*isDeleted))
	}
	only, err := query.Only(ctx)
	if err != nil {
		return nil, f.isFilesError(err)
	}
	return only, nil
}

func (f filesMySQL) GetFileByFilenameBucketID(ctx context.Context, filename string, spacesID int, isDraft *bool, isDeleted *bool) (*ent.Files, *ent.Buckets, error) {
	query := f.client.Files.Query().Where(files.Filename(filename), files.HasBucketsWith(buckets.ID(spacesID)))

	if isDraft != nil {
		query.Where(files.IsDraft(*isDraft))
	}
	if isDeleted != nil {
		query.Where(files.IsDeleted(*isDeleted))
	}
	file, err := query.Only(ctx)
	if err != nil {
		return nil, nil, f.isFilesError(err)
	}
	spaces, err := file.QueryBuckets().Where(buckets.ID(spacesID)).Only(ctx)
	if err != nil {
		return nil, nil, f.isSpacesError(err)
	}
	return file, spaces, nil
}

func (f filesMySQL) GetFileByFilenameBucketName(ctx context.Context, filename string, spacesName string, isDraft *bool, isDeleted *bool) (*ent.Files, *ent.Buckets, error) {
	query := f.client.Files.Query().Where(files.Filename(filename), files.HasBucketsWith(buckets.Name(spacesName)))
	if isDraft != nil {
		query.Where(files.IsDraft(*isDraft))
	}
	if isDeleted != nil {
		query.Where(files.IsDeleted(*isDeleted))
	}
	file, err := query.Only(ctx)
	if err != nil {
		return nil, nil, f.isFilesError(err)
	}
	spaces, err := file.QueryBuckets().Where(buckets.Name(spacesName)).Only(ctx)
	if err != nil {
		return nil, nil, f.isSpacesError(err)
	}
	return file, spaces, nil
}

func (f filesMySQL) DeleteFileByID(ctx context.Context, fileID int) error {
	_, err := f.GetFileByID(ctx, fileID, nil, nil)
	if err != nil {
		return err
	}
	_, err = f.client.Files.Delete().Where(files.ID(fileID)).Exec(ctx)
	if err != nil {
		return err
	}
	return err
}

func (f filesMySQL) MarkAsDeletedByID(ctx context.Context, fileID int) (*ent.Files, error) {
	file, err := f.GetFileByID(ctx, fileID, nil, nil)
	if err != nil {
		return nil, err
	}
	return file.Update().SetIsDeleted(true).SetUpdatedAt(time.Now().UTC()).Save(ctx)
}

func (f filesMySQL) MarkAsDraftByID(ctx context.Context, fileID int) (*ent.Files, error) {
	panic("implement me")
}

func (f filesMySQL) CreateFile(ctx context.Context, filename string, fileSize int64, mimeType string, isDeleted bool, isDraft bool, userId int64, spacesID int) (*ent.Files, error) {
	if _,_, err := f.GetFileByFilenameBucketID(ctx, filename, spacesID, nil, nil); err == nil {
		return nil, ErrFileExist
	}
	return f.client.Files.Create().
		SetFileSize(fileSize).
		SetFilename(filename).
		SetUserID(userId).
		SetIsDraft(isDraft).
		SetCreatedAt(time.Now().UTC()).
		SetUpdatedAt(time.Now().UTC()).
		SetIsDeleted(isDeleted).
		SetBucketsID(spacesID).
		SetMimeType(mimeType).
		Save(ctx)
}

func (f filesMySQL) isFilesError(err error) error {
	if ent.IsNotFound(err) {
		return ErrFileNotExist
	}
	return err
}

func (f filesMySQL) isSpacesError(err error) error {
	if ent.IsNotFound(err) {
		return bucketsrepo.ErrBucketsNotExist
	}
	return err
}

func NewFileRepository(client *ent.Client) FileRepository {
	return filesMySQL{client:client}
}