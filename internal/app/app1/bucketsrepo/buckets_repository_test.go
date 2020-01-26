package bucketsrepo

import (
	"context"
	"github.com/pepeunlimited/files/internal/app/app1/mysql"
	"testing"
)

func TestDOBucketMySQL_CreateBucketAndDelete(t *testing.T) {
	ctx := context.TODO()
	spaces := NewBucketsRepository(mysql.NewEntClient())
	spaces.Wipe(ctx)

	bucketName := "bucket-666"
	endpoint := "fra1.digitaloceanspaces.com"
	cdn := bucketName+".fra1.cdn.digitaloceanspaces.com"

	bucket, err := spaces.Create(ctx, bucketName, endpoint, &cdn)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if bucket == nil {
		t.FailNow()
	}
	selected, err := spaces.GetBucketsByName(ctx, bucketName)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if selected.Name != bucketName {
		t.FailNow()
	}
	err = spaces.DeleteBucketByName(ctx, bucketName)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, err = spaces.GetBucketsByName(ctx, bucketName)
	if err == nil {
		t.FailNow()
	}
	if err != ErrBucketsNotExist {
		t.FailNow()
	}
	err = spaces.DeleteBucketByName(ctx, bucketName)
	if err == nil {
		t.FailNow()
	}
	if err != ErrBucketsNotExist {
		t.FailNow()
	}
}

func TestDobucketMySQL_GetBucketByID(t *testing.T) {
	ctx := context.TODO()
	bucketRepo := NewBucketsRepository(mysql.NewEntClient())
	bucketRepo.Wipe(ctx)

	bucketName := "bucket-666"
	endpoint := "fra1.digitaloceanspaces.com"
	cdn := bucketName+".fra1.cdn.digitaloceanspaces.com"

	bucket, err := bucketRepo.Create(ctx, bucketName, endpoint, &cdn)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	selected, err := bucketRepo.GetBucketByID(ctx, bucket.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if selected.ID != bucket.ID {
		t.FailNow()
	}
}

func TestDobucketsMySQL_GetBuckets(t *testing.T) {
	ctx := context.TODO()
	bucketRepo := NewBucketsRepository(mysql.NewEntClient())
	bucketRepo.Wipe(ctx)

	bucketName := "bucket-666"
	endpoint := "fra1.digitaloceanspaces.com"
	cdn := bucketName+".fra1.cdn.digitaloceanspaces.com"

	if _, err := bucketRepo.Create(ctx, bucketName+"-1", endpoint, &cdn); err != nil {
		t.Error(err)
		t.FailNow()
	}
	bucketRepo.Create(ctx, bucketName+"-2", endpoint, &cdn)
	bucketRepo.Create(ctx, bucketName+"-3", endpoint, &cdn)
	bucketRepo.Create(ctx, bucketName+"-4", endpoint, &cdn)

	buckets0, nextPageToken0, err := bucketRepo.GetBuckets(ctx, 0, 1)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if len(buckets0) != 1 {
		t.FailNow()
	}
	buckets1, nextPageToken1, err := bucketRepo.GetBuckets(ctx, nextPageToken0, 20)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if len(buckets1) != 3 {
		t.FailNow()
	}
	buckets2, nextPageToken2, err := bucketRepo.GetBuckets(ctx, nextPageToken1, 20)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if len(buckets2) != 0 {
		t.FailNow()
	}
	if nextPageToken2 == 0 {
		t.FailNow()
	}
}