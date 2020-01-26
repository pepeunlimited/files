package server

import (
	"context"
	"github.com/pepeunlimited/authentication-twirp/authrpc"
	"github.com/pepeunlimited/files/internal/app/app1/mysql"
	"github.com/pepeunlimited/files/internal/app/app1/upload"
	"github.com/pepeunlimited/files/filesrpc"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestDOFileUploadServer_UploadDOV1FilesSuccess(t *testing.T) {
	ctx := context.TODO()
	authClient := authrpc.NewAuthenticationMock(nil)
	mock := upload.NewActionsMock(nil)

	server := NewUploadServer(mock, mysql.NewEntClient(), authClient)
	server.buckets.Wipe(ctx)

	fileServer := NewFilesServer(mock, mysql.NewEntClient())
	fileServer.CreateBucket(ctx, &filesrpc.CreateBucketParams{
		Name: "bucket",
		Endpoint:   "fra1.mock.com",
	})

	body := strings.NewReader("Hello-World!\n\r Hei Maailma!")
	fileSize     := body.Size()
	contentType := "plain/text"
	contentLength := strconv.Itoa(int(fileSize))
	authorization := "Bearer A"

	// request
	req,_ := http.NewRequest(http.MethodPost, UploadV1Files, body)
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Content-Length", contentLength)
	req.Header.Add("Authorization", authorization)
	req.Header.Add("Meta-API-Args", "{\"filename\": \"filename.txt\"}")

	// recorder
	recorder := httptest.NewRecorder()
	server.UploadV1Files().ServeHTTP(recorder, req)

	if recorder.Code != 200 {
		t.Log(recorder.Code)
		t.Log(recorder.Body.String())
		t.FailNow()
	}
}


func TestDOFileUploadServer_UploadDOV1FilesFailed(t *testing.T) {


}