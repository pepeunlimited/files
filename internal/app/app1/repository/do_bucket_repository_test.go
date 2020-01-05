package repository

import (
	"context"
	"testing"
)

func TestDOBucketMySQL_CreateBucketAndDelete(t *testing.T) {
	ctx := context.TODO()
	bucketRepo := NewDOBucketRepository(NewEntClient())
	bucketRepo.DeleteAll(ctx)

	bucketName := "bucket-666"
	endpoint := "fra1.digitaloceanspaces.com"
	cdn := bucketName+".fra1.cdn.digitaloceanspaces.com"

	bucket, err := bucketRepo.CreateBucket(ctx, bucketName, endpoint, &cdn)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if bucket == nil {
		t.FailNow()
	}
	selected, err := bucketRepo.GetBucketByName(ctx, bucketName)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if selected.BucketName != bucketName {
		t.FailNow()
	}
	err = bucketRepo.DeleteBucketByName(ctx, bucketName)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, err = bucketRepo.GetBucketByName(ctx, bucketName)
	if err == nil {
		t.FailNow()
	}
	if err != ErrDOBucketNotExist {
		t.FailNow()
	}
	err = bucketRepo.DeleteBucketByName(ctx, bucketName)
	if err == nil {
		t.FailNow()
	}
	if err != ErrDOBucketNotExist {
		t.FailNow()
	}
}

func TestDobucketMySQL_GetBucketByID(t *testing.T) {
	ctx := context.TODO()
	bucketRepo := NewDOBucketRepository(NewEntClient())
	bucketRepo.DeleteAll(ctx)

	bucketName := "bucket-666"
	endpoint := "fra1.digitaloceanspaces.com"
	cdn := bucketName+".fra1.cdn.digitaloceanspaces.com"

	bucket, err := bucketRepo.CreateBucket(ctx, bucketName, endpoint, &cdn)
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