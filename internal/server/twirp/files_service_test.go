package twirp

import (
	"context"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/pepeunlimited/files/internal/pkg/ent"
	"github.com/pepeunlimited/files/pkg/fs"
	"github.com/pepeunlimited/files/pkg/rpc/files"
	"testing"
)

func TestDOFileServer_GetFileByID(t *testing.T) {
	ctx 	  := context.TODO()
	fs        := fs.NewMock("mock.endpoint")
	server    := NewFilesServer(ent.NewEntClient(), fs)
	server.buckets.Wipe(ctx)

	doBucket,_ := server.buckets.Create(ctx, "lol", "aaaa", nil)
	created,_ := server.files.CreateFile(ctx, "filename", 1, "mimetype", false, false, 1, doBucket.ID)

	resp0, err := server.GetFile(ctx, &files.GetFileParams{
		FileId: &wrappers.Int64Value{
			Value: int64(created.ID),
		},
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if resp0.Id != int64(created.ID) {
		t.FailNow()
	}
	if resp0.SpacesId != int64(doBucket.ID) {
		t.FailNow()
	}
	if resp0.UserId != created.UserID {
		t.FailNow()
	}
	if resp0.MimeType != created.MimeType {
		t.FailNow()
	}
	if resp0.Filename != created.Filename {
		t.FailNow()
	}
}

func TestDOFileServer_GetFileByFilename(t *testing.T) {
	ctx 	   := context.TODO()
	fs        := fs.NewMock("mock.endpoint")
	server    := NewFilesServer(ent.NewEntClient(), fs)
	server.buckets.Wipe(ctx)

	bucket0,_ := server.CreateBucket(ctx, &files.CreateBucketParams{
		Name: "lol-1",
	})
	bucket1,_ := server.CreateBucket(ctx, &files.CreateBucketParams{
		Name: "lol-2",
	})

	file0, err := server.files.CreateFile(ctx, "filename", 1, "mimetype2", false, false, 1, int(bucket0.BucketId))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, err = server.files.CreateFile(ctx, "filename", 1, "mimetype2", false, false, 1, int(bucket1.BucketId))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	resp0, err := server.GetFile(ctx, &files.GetFileParams{
		Filename: &files.Filename{
			Name:       "filename",
			BucketId:   &wrappers.Int64Value{
				Value: bucket0.BucketId,
			},
			BucketName: nil,
		},
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if resp0.Id != int64(file0.ID) {
		t.FailNow()
	}
	if resp0.SpacesId != bucket0.BucketId {
		t.FailNow()
	}
	if resp0.UserId != file0.UserID {
		t.FailNow()
	}
	if resp0.MimeType != file0.MimeType {
		t.FailNow()
	}
	if resp0.Filename != file0.Filename {
		t.FailNow()
	}
	resp0, err = server.GetFile(ctx, &files.GetFileParams{
		Filename: &files.Filename{
			Name:       "filename",
			BucketName: &wrappers.StringValue{
				Value: "lol-1",
			},
		},
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if resp0.Id != int64(file0.ID) {
		t.FailNow()
	}
	if resp0.SpacesId != bucket0.BucketId {
		t.FailNow()
	}
	if resp0.UserId != file0.UserID {
		t.FailNow()
	}
	if resp0.MimeType != file0.MimeType {
		t.FailNow()
	}
	if resp0.Filename != file0.Filename {
		t.FailNow()
	}
}

func TestDOFileServer_CreateBucket(t *testing.T) {
	ctx 	  := context.TODO()
	fs        := fs.NewMock("mock.endpoint")
	server    := NewFilesServer(ent.NewEntClient(), fs)
	server.buckets.Wipe(ctx)
	bucket, err := server.CreateBucket(ctx, &files.CreateBucketParams{
		Name:  "bucket-test",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	name, err := server.buckets.GetBucketsByName(ctx, "bucket-test")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if name.Name != bucket.Name {
		t.FailNow()
	}
	//if !mock.(*upload.ActionsMock).IsCreateBucket {
	//	t.FailNow()
	//}
	//if mock.(*upload.ActionsMock).IsDeleteBucket {
	//	t.FailNow()
	//}
}

func TestDOFileServer_Delete(t *testing.T) {
	ent       := ent.NewEntClient()
	ctx 	  := context.TODO()
	fs        := fs.NewMock("mock.endpoint")
	server    := NewFilesServer(ent, fs)
	server.buckets.Wipe(ctx)

	doBucket,_ := server.buckets.Create(ctx, "bucket", "endpoint", nil)
	server.files.CreateFile(ctx, "filename.txt", 15, "plain/text", false, false, 1, doBucket.ID)
	file1,_ := server.files.CreateFile(ctx, "filename2.txt", 15, "plain/text", false, false, 1, doBucket.ID)
	_, err := server.Delete(ctx, &files.DeleteParams{
		Filename: &files.Filename{
			Name: "filename.txt",
			BucketName: &wrappers.StringValue{
				Value: "bucket",
			},
		},
		IsPermanent: false,
	})
	//if mock.(*upload.ActionsMock).IsDeleteBucket {
	//	t.FailNow()
	//}
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	query0,_, err := server.files.GetFileByFilenameBucketID(ctx, "filename.txt", doBucket.ID, nil, nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if !query0.IsDeleted {
		t.FailNow()
	}
	_, err = server.Delete(ctx, &files.DeleteParams{
		FileId:      &wrappers.Int64Value{
			Value: int64(file1.ID),
		},
		IsPermanent: false,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	query1, err := server.files.GetFileByID(ctx, file1.ID, nil, nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if !query1.IsDeleted {
		t.FailNow()
	}
	//if mock.(*upload.ActionsMock).IsDeleteBucket {
	//	t.FailNow()
	//}
}