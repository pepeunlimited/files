package filerepo

import (
	"context"
	"github.com/pepeunlimited/files/internal/app/app1/mysql"
	"github.com/pepeunlimited/files/internal/app/app1/bucketsrepo"
	"testing"
)

func TestFilesMySQL_CreateDOFile(t *testing.T) {
	ctx    := context.TODO()
	client := mysql.NewEntClient()
	bucket := bucketsrepo.NewBucketsRepository(client)
	files := NewFileRepository(client)
	bucket.Wipe(ctx)

	filename := "filename.txt"
	fileSize := int64(12)
	mimeType := "plain/text"
	isDeleted := false
	isDraft := false
	userId := int64(1)

	created, err := bucket.Create(ctx, "bucket-name", "e", nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, err = files.CreateFile(ctx, filename, fileSize, mimeType, isDeleted, isDraft, userId, created.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestFilesMySQL_GetFilesDOBucketByID2(t *testing.T) {
	ctx    := context.TODO()
	client := mysql.NewEntClient()
	bucket := bucketsrepo.NewBucketsRepository(client)
	files := NewFileRepository(client)
	bucket.Wipe(ctx)

	filename := "filename.txt"
	fileSize := int64(12)
	mimeType := "plain/text"
	isDeleted := false
	isDraft := false
	userId := int64(1)

	doBucket,_ := bucket.Create(ctx, "bucket-name", "e", nil)
	file, err := files.CreateFile(ctx, filename, fileSize, mimeType, isDeleted, isDraft, userId, doBucket.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	file2, doBucket2, err := files.GetFilesBucketByID(ctx, file.ID, nil, nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if doBucket2.ID != doBucket.ID {
		t.FailNow()
	}
	if file2.ID != file.ID {
		t.FailNow()
	}
}

func TestFilesMySQL_GetFilesDOBucketByIDNotFound(t *testing.T) {
	ctx    := context.TODO()
	client := mysql.NewEntClient()
	bucket := bucketsrepo.NewBucketsRepository(client)
	files := NewFileRepository(client)
	bucket.Wipe(ctx)
	_, _, err := files.GetFilesBucketByID(ctx, 11111111, nil, nil)
	if err == nil {
		t.FailNow()
	}
	if err != ErrFileNotExist {
		t.FailNow()
	}
}

func TestFilesMySQL_GetFileByFilename(t *testing.T) {
	ctx    := context.TODO()
	client := mysql.NewEntClient()
	bucket := bucketsrepo.NewBucketsRepository(client)
	files := NewFileRepository(client)
	bucket.Wipe(ctx)

	filename := "filename.txt"
	fileSize := int64(12)
	mimeType := "plain/text"
	isDeleted := false
	isDraft := false
	userId := int64(1)

	bucket0,_ := bucket.Create(ctx, "bucket-name-0", "e0", nil)
	bucket1,_ := bucket.Create(ctx, "bucket-name-1", "e1", nil)

	_, err := files.CreateFile(ctx, filename, fileSize, mimeType, isDeleted, isDraft, userId, bucket0.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, err = files.CreateFile(ctx, filename, fileSize, mimeType, isDeleted, isDraft, userId, bucket1.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	_,_, err = files.GetFileByFilenameBucketName(ctx, filename, bucket0.Name, nil, nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_,_, err = files.GetFileByFilenameBucketID(ctx, filename, bucket1.ID, nil, nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestFilesMySQL_GetFileByID(t *testing.T) {
	ctx    := context.TODO()
	client := mysql.NewEntClient()
	bucket := bucketsrepo.NewBucketsRepository(client)
	files := NewFileRepository(client)
	bucket.Wipe(ctx)

	filename := "filename.txt"
	fileSize := int64(12)
	mimeType := "plain/text"
	isDeleted := false
	isDraft := false
	userId := int64(1)

	bucket0,_ := bucket.Create(ctx, "bucket-name-0", "e0", nil)
	bucket1,_ := bucket.Create(ctx, "bucket-name-1", "e1", nil)

	_, err := files.CreateFile(ctx, filename, fileSize, mimeType, isDeleted, isDraft, userId, bucket0.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	file1, err := files.CreateFile(ctx, filename, fileSize, mimeType, isDeleted, isDraft, userId, bucket1.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	_,_, err = files.GetFilesBucketByID(ctx, file1.ID, nil, nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_,_, err = files.GetFileByFilenameBucketName(ctx, filename, bucket1.Name, nil, nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestFilesMySQL_MarkAsDeletedByID(t *testing.T) {
	ctx    := context.TODO()
	client := mysql.NewEntClient()
	bucket := bucketsrepo.NewBucketsRepository(client)
	files := NewFileRepository(client)
	bucket.Wipe(ctx)

	filename := "filename.txt"
	fileSize := int64(12)
	mimeType := "plain/text"
	isDeleted := false
	isDraft := false
	userId := int64(1)

	bucket0,_ := bucket.Create(ctx, "bucket-name-0", "e0", nil)
	file, err := files.CreateFile(ctx, filename, fileSize, mimeType, isDeleted, isDraft, userId, bucket0.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, err = files.MarkAsDeletedByID(ctx, file.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	selected, err := files.GetFileByID(ctx, file.ID, nil, nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if !selected.IsDeleted {
		t.FailNow()
	}
}