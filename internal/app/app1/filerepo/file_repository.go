package filerepo

import (
	"context"
	"errors"
	"github.com/pepeunlimited/files/internal/app/app1/ent"
	"github.com/pepeunlimited/files/internal/app/app1/ent/files"
	"github.com/pepeunlimited/files/internal/app/app1/ent/spaces"
	"github.com/pepeunlimited/files/internal/app/app1/spacesrepo"
	"time"
)

var (
	ErrFileExist 		= errors.New("files: already exist")
	ErrFileNotExist 	= errors.New("files: not exist")
)

// many-to-one do_buckets
type FileRepository interface {
	CreateSpacesFile(ctx context.Context, filename string, fileSize int64, mimeType string, isDeleted bool, isDraft bool, userId int64, spacesID int) (*ent.Files, error)

	GetFilesSpacesByID(ctx context.Context, fileID int, isDraft *bool, isDeleted *bool)		(*ent.Files, *ent.Spaces, error)
 	GetFileByID(ctx context.Context, fileID int, isDraft *bool, isDeleted *bool)				(*ent.Files, error)

	MarkAsDeletedByID(ctx context.Context, fileID int)										(*ent.Files, error)
	MarkAsDraftByID(ctx context.Context, fileID int)											(*ent.Files, error)

	// spacesID or spacesName is required because same filename may be in another spaces..
	GetFileByFilenameSpacesID(ctx context.Context, filename string, spacesID int, isDraft *bool, isDeleted *bool)	(*ent.Files, *ent.Spaces, error)
	GetFileByFilenameSpacesName(ctx context.Context, filename string, spacesName string, isDraft *bool, isDeleted *bool) (*ent.Files, *ent.Spaces, error)

	DeleteFileByID(ctx context.Context, fileID int) error
	ExistInSpaces(ctx context.Context, filename string, spacesID int) (*bool, error)
}

type filesMySQL struct {
	client  *ent.Client
}

func (f filesMySQL) ExistInSpaces(ctx context.Context, filename string, spacesID int) (*bool, error) {
	_, _, err := f.GetFileByFilenameSpacesID(ctx, filename, spacesID, nil, nil)
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

func (f filesMySQL) GetFilesSpacesByID(ctx context.Context, fileID int, isDraft *bool, isDeleted *bool) (*ent.Files, *ent.Spaces, error) {
	file, err := f.GetFileByID(ctx, fileID, isDraft, isDeleted)
	if err != nil {
		return nil, nil, err
	}
	bucket, err := file.QuerySpaces().Only(ctx)
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

func (f filesMySQL) GetFileByFilenameSpacesID(ctx context.Context, filename string, spacesID int, isDraft *bool, isDeleted *bool) (*ent.Files, *ent.Spaces, error) {
	query := f.client.Files.Query().Where(files.Filename(filename), files.HasSpacesWith(spaces.ID(spacesID)))

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
	spaces, err := file.QuerySpaces().Where(spaces.ID(spacesID)).Only(ctx)
	if err != nil {
		return nil, nil, f.isSpacesError(err)
	}
	return file, spaces, nil
}

func (f filesMySQL) GetFileByFilenameSpacesName(ctx context.Context, filename string, spacesName string, isDraft *bool, isDeleted *bool) (*ent.Files, *ent.Spaces, error) {
	query := f.client.Files.Query().Where(files.Filename(filename), files.HasSpacesWith(spaces.Name(spacesName)))
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
	spaces, err := file.QuerySpaces().Where(spaces.Name(spacesName)).Only(ctx)
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

func (f filesMySQL) CreateSpacesFile(ctx context.Context, filename string, fileSize int64, mimeType string, isDeleted bool, isDraft bool, userId int64, spacesID int) (*ent.Files, error) {
	if _,_, err := f.GetFileByFilenameSpacesID(ctx, filename, spacesID, nil, nil); err == nil {
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
		SetSpacesID(spacesID).
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
		return spacesrepo.ErrSpacesNotExist
	}
	return err
}

func NewFileRepository(client *ent.Client) FileRepository {
	return filesMySQL{client:client}
}