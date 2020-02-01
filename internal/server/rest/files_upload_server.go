package rest

import (
	"github.com/pepeunlimited/authentication-twirp/pkg/authrpc"
	"github.com/pepeunlimited/files/internal/pkg/ent"
	"github.com/pepeunlimited/files/internal/pkg/mysql/bucketsrepo"
	"github.com/pepeunlimited/files/internal/pkg/mysql/filerepo"
	"github.com/pepeunlimited/files/internal/pkg/upload"
	"github.com/pepeunlimited/files/internal/server/validator"
	"github.com/pepeunlimited/files/pkg/filesrpc"
	"github.com/pepeunlimited/files/pkg/storage"
	"github.com/pepeunlimited/microservice-kit/httpz"
	"github.com/twitchtv/twirp"
	"log"
	"math/rand"
	"net/http"
)

const (
	// DigitalOcean's
	// version
	VersionV1            = "/v1"
	UploadV1Files string = UploadPath +VersionV1+ FilesPath
)

type UploadServer struct {
	validator      validator.UploadServerValidator
	actions        storage.Actions
	files          filerepo.FileRepository
	buckets        bucketsrepo.BucketsRepository
	authentication authrpc.AuthenticationService
}

// https://phil.tech/api/2016/01/04/http-rest-api-file-uploads/
func (server UploadServer) UploadV1Files() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header, args, err := server.validator.UploadSpacesV1Files(r.Header, r.Body)
		if err != nil {
			httpz.WriteError(w, err)
			return
		}
		// validate the access token..
		user, err := server.authentication.VerifyAccessToken(r.Context(), &authrpc.VerifyAccessTokenParams{
			AccessToken: header.Authorization,
		})
		if err != nil {
			switch err.(twirp.Error).Msg() {
			case authrpc.AccessTokenExpired:
				httpz.WriteError(w,httpz.NewMsgError(authrpc.AccessTokenExpired, http.StatusUnauthorized))
			case authrpc.AccessTokenMalformed:
				httpz.WriteError(w,httpz.NewMsgError(authrpc.AccessTokenMalformed, http.StatusBadRequest))
			case authrpc.AccessTokenUnknownError:
				httpz.WriteError(w,httpz.NewMsgError(authrpc.AccessTokenUnknownError, http.StatusInternalServerError))
			default:
				log.Print("buckets-upload: failed: "+err.Error())
				httpz.WriteError(w,httpz.NewMsgError(filesrpc.FileUploadFailed, http.StatusInternalServerError))
			}
			return
		}
		userId := user.UserId

		buckets,_, err := server.buckets.GetBuckets(r.Context(), 0, 20)
		if err != nil {
			log.Print("buckets-upload: failed: "+err.Error())
			httpz.WriteError(w, httpz.NewMsgError(filesrpc.FileUploadFailed, http.StatusInternalServerError))
			return
		}
		if len(buckets) == 0 {
			log.Print("buckets-upload: missing buckets!")
			httpz.WriteError(w, httpz.NewMsgError(filesrpc.FileUploadFailed, http.StatusInternalServerError))
			return
		}

		min := 0
		max := len(buckets)
		random := rand.Intn(max - min) + min
		bucket := buckets[random]

		exist, err := server.files.ExistInBucket(r.Context(), args.Filename, bucket.ID)
		if err != nil {
			log.Print("buckets-upload: failed: "+err.Error())
			httpz.WriteError(w, httpz.NewMsgError(filesrpc.FileUploadFailed, http.StatusInternalServerError))
			return
		}

		if *exist {
			httpz.WriteError(w, httpz.NewMsgError(filesrpc.FileExist, http.StatusBadRequest))
			return
		}

		// upload to the buckets..
		err = server.actions.Upload(
			r.Context(),
			storage.File{
			MimeType: header.ContentType,
			Body:     r.Body,
			FileSize: header.ContentLength,
		},
		storage.FileMetaData{
			Filename: args.Filename,
			IsPublic: true,
		},
		storage.Buckets{
			BucketName: bucket.Name,
			Endpoint:   bucket.Endpoint,
		})
		if err != nil {
			log.Print("buckets-upload: failed: " + err.Error())
			httpz.WriteError(w, httpz.NewMsgError(filesrpc.FileUploadFailed, http.StatusInternalServerError))
			return
		}

		// save to the DB
		file, err := server.files.CreateFile(r.Context(), args.Filename, header.ContentLength, header.ContentType, false, false, userId, bucket.ID)
		if err != nil {
			log.Print("buckets-upload: failed: " + err.Error())
			// rollback
			server.actions.Delete(
				r.Context(),
				storage.Buckets{
				BucketName: bucket.Name,
				Endpoint:   bucket.Endpoint,
			},
			args.Filename)
			httpz.WriteError(w, httpz.NewMsgError(filesrpc.FileUploadFailed, http.StatusInternalServerError))
			return
		}
		httpz.WriteOk(w, upload.UploadDOV1Files{
			Endpoint:    bucket.Endpoint,
			CDNEndpoint: *bucket.CdnEndpoint,
			FileId:      int64(file.ID),
			Filename:    args.Filename,
			URI:		 "https://"+*bucket.CdnEndpoint+"/"+args.Filename,
		})
	})
}

func NewUploadServer(actions storage.Actions, client *ent.Client, authentication authrpc.AuthenticationService) UploadServer {
	return UploadServer{
		actions:        actions,
		authentication: authentication,
		validator:      validator.NewSpacesUploadServerValidator(),
		files:          filerepo.NewFileRepository(client),
		buckets:        bucketsrepo.NewBucketsRepository(client),
	}
}

