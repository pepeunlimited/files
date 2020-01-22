package server

import (
	"context"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/pepeunlimited/files/internal/app/app1/mysql"
	"github.com/pepeunlimited/files/internal/app/app1/upload"
	"github.com/pepeunlimited/files/spacesrpc"
	"testing"
)

func TestDOFileServer_GetFileByID(t *testing.T) {
	ctx 	  := context.TODO()
	mock      := upload.NewDosMock(nil)
	server    := NewSpacesServer(mock, mysql.NewEntClient())
	server.spaces.Wipe(ctx)

	doBucket,_ := server.spaces.Create(ctx, "lol", "aaaa", nil)
	created,_ := server.files.CreateSpacesFile(ctx, "filename", 1, "mimetype", false, false, 1, doBucket.ID)

	resp0, err := server.GetFile(ctx, &spacesrpc.GetFileParams{
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
	ctx 	  := context.TODO()
	mock      := upload.NewDosMock(nil)
	server    := NewSpacesServer(mock, mysql.NewEntClient())
	server.spaces.Wipe(ctx)

	bucket0,_ := server.CreateSpaces(ctx, &spacesrpc.CreateSpacesParams{
		Name: "lol-1",
		Endpoint:   "endpoint1.aa.com",
	})
	bucket1,_ := server.CreateSpaces(ctx, &spacesrpc.CreateSpacesParams{
		Name: "lol-2",
		Endpoint:   "endpoint2.aa.com",
	})

	file0, err := server.files.CreateSpacesFile(ctx, "filename", 1, "mimetype2", false, false, 1, int(bucket0.SpacesId))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, err = server.files.CreateSpacesFile(ctx, "filename", 1, "mimetype2", false, false, 1, int(bucket1.SpacesId))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	resp0, err := server.GetFile(ctx, &spacesrpc.GetFileParams{
		Filename: &spacesrpc.Filename{
			Name:       "filename",
			BucketId:   &wrappers.Int64Value{
				Value: bucket0.SpacesId,
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
	if resp0.SpacesId != bucket0.SpacesId {
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
	resp0, err = server.GetFile(ctx, &spacesrpc.GetFileParams{
		Filename: &spacesrpc.Filename{
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
	if resp0.SpacesId != bucket0.SpacesId {
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
	mock      := upload.NewDosMock(nil)
	server    := NewSpacesServer(mock, mysql.NewEntClient())
	server.spaces.Wipe(ctx)
	bucket, err := server.CreateSpaces(ctx, &spacesrpc.CreateSpacesParams{
		Name:  "bucket-test",
		Endpoint:    "fra.endpoint.com",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	name, err := server.spaces.GetSpaceByName(ctx, "bucket-test")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if name.Name != bucket.Name {
		t.FailNow()
	}
	if !mock.(*upload.DosMock).IsCreateBucket {
		t.FailNow()
	}
	if mock.(*upload.DosMock).IsDeleteBucket {
		t.FailNow()
	}
}

func TestDOFileServer_Delete(t *testing.T) {
	ent       := mysql.NewEntClient()
	ctx 	  := context.TODO()
	mock      := upload.NewDosMock(nil)
	server    := NewSpacesServer(mock, ent)
	server.spaces.Wipe(ctx)

	doBucket,_ := server.spaces.Create(ctx, "bucket", "endpoint", nil)
	server.files.CreateSpacesFile(ctx, "filename.txt", 15, "plain/text", false, false, 1, doBucket.ID)
	file1,_ := server.files.CreateSpacesFile(ctx, "filename2.txt", 15, "plain/text", false, false, 1, doBucket.ID)
	_, err := server.Delete(ctx, &spacesrpc.DeleteParams{
		Filename: &spacesrpc.Filename{
			Name: "filename.txt",
			BucketName: &wrappers.StringValue{
				Value: "bucket",
			},
		},
		IsPermanent: false,
	})
	if mock.(*upload.DosMock).IsDeleteBucket {
		t.FailNow()
	}
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	query0,_, err := server.files.GetFileByFilenameSpacesID(ctx, "filename.txt", doBucket.ID, nil, nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if !query0.IsDeleted {
		t.FailNow()
	}
	_, err = server.Delete(ctx, &spacesrpc.DeleteParams{
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
	if mock.(*upload.DosMock).IsDeleteBucket {
		t.FailNow()
	}
}