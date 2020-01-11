package server

import (
	"context"
	rpc2 "github.com/pepeunlimited/authorization-twirp/rpc"
	"github.com/pepeunlimited/files/internal/app/app1/repository"
	"github.com/pepeunlimited/files/internal/app/app1/upload"
	"github.com/pepeunlimited/files/rpc"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestDOFileUploadServer_UploadDOV1FilesSuccess(t *testing.T) {
	ctx := context.TODO()
	authClient := rpc2.NewAuthorizationMock(nil)
	mock := upload.NewDosMock(nil)

	server := NewDOFileUploadServer(mock, repository.NewEntClient(), authClient)
	server.bucketRepository.Wipe(ctx)

	fileServer := NewDOFileServer(mock, repository.NewEntClient())
	fileServer.CreateBucket(ctx, &rpc.CreateBucketParams{
		BucketName: "bucket",
		Endpoint:   "fra1.mock.com",
	})

	body := strings.NewReader("Hello-World!\n\r Hei Maailma!")
	fileSize     := body.Size()
	contentType := "plain/text"
	contentLength := strconv.Itoa(int(fileSize))
	authorization := "Bearer A"

	// request
	req,_ := http.NewRequest(http.MethodPost, UploadDOV1Files, body)
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Content-Length", contentLength)
	req.Header.Add("Authorization", authorization)
	req.Header.Add("Meta-API-Args", "{\"filename\": \"filename.txt\"}")

	// recorder
	recorder := httptest.NewRecorder()
	server.UploadDOV1Files().ServeHTTP(recorder, req)

	if recorder.Code != 200 {
		t.Log(recorder.Code)
		t.Log(recorder.Body.String())
		t.FailNow()
	}
}


func TestDOFileUploadServer_UploadDOV1FilesFailed(t *testing.T) {


}