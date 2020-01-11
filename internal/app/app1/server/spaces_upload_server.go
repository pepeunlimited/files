package server

import (
	rpc2 "github.com/pepeunlimited/authorization-twirp/rpc"
	"github.com/pepeunlimited/files/internal/app/app1/ent"
	"github.com/pepeunlimited/files/internal/app/app1/repository"
	"github.com/pepeunlimited/files/internal/app/app1/upload"
	"github.com/pepeunlimited/files/internal/app/app1/validator"
	"github.com/pepeunlimited/files/rpc"
	"github.com/pepeunlimited/files/storage"
	"github.com/pepeunlimited/microservice-kit/httpz"
	"github.com/twitchtv/twirp"
	"log"
	"math/rand"
	"net/http"
)

const (
	// DigitalOcean's
	DoPath      = "/do"
	// do version
	DoVersionV1   = "/v1"
	UploadDOV1Files string = UploadPath+DoPath+DoVersionV1+FilesPath
)

type DOFileUploadServer struct {
	validator 			validator.FileUploadServerValidator
	actions 			storage.Actions
	filesRepository  	repository.FileRepository
	bucketRepository 	repository.DOBucketRepository
	authService 		rpc2.AuthorizationService
}

// https://phil.tech/api/2016/01/04/http-rest-api-file-uploads/
func (server DOFileUploadServer) UploadDOV1Files() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header, args, err := server.validator.UploadDOV1Files(r.Header, r.Body)
		if err != nil {
			httpz.WriteError(w, err)
			return
		}
		// validate the access token..
		user, err := server.authService.VerifyAccessToken(r.Context(), &rpc2.VerifyAccessTokenParams{
			AccessToken: header.Authorization,
		})
		if err != nil {
			if rpc.IsReason(err.(twirp.Error), rpc2.AccessTokenExpired) {
				httpz.WriteError(w,httpz.NewMsgError(rpc2.AccessTokenExpired, http.StatusUnauthorized))
			} else if rpc.IsReason(err.(twirp.Error), rpc2.AccessTokenMalformed) {
				httpz.WriteError(w,httpz.NewMsgError(rpc2.AccessTokenMalformed, http.StatusBadRequest))
			} else if rpc.IsReason(err.(twirp.Error), rpc2.AccessTokenUnknownError) {
				httpz.WriteError(w,httpz.NewMsgError(rpc2.AccessTokenUnknownError, http.StatusInternalServerError))
			} else {
				log.Print("do-file-upload: failed: "+err.Error())
				httpz.WriteError(w,httpz.NewMsgError(rpc.FileUploadFailed, http.StatusInternalServerError))
			}
			return
		}
		userId := user.UserId

		buckets,_, err := server.bucketRepository.GetBuckets(r.Context(), 0, 20)
		if err != nil {
			log.Print("do-file-upload: failed: "+err.Error())
			httpz.WriteError(w, httpz.NewMsgError(rpc.FileUploadFailed, http.StatusInternalServerError))
			return
		}
		if len(buckets) == 0 {
			log.Print("do-file-upload: missing buckets!")
			httpz.WriteError(w, httpz.NewMsgError(rpc.FileUploadFailed, http.StatusInternalServerError))
			return
		}

		min := 0
		max := len(buckets)
		random := rand.Intn(max - min) + min
		bucket := buckets[random]

		exist, err := server.filesRepository.Exist(r.Context(), args.Filename, bucket.ID)
		if err != nil {
			log.Print("do-file-upload: failed: "+err.Error())
			httpz.WriteError(w, httpz.NewMsgError(rpc.FileUploadFailed, http.StatusInternalServerError))
			return
		}

		if *exist {
			httpz.WriteError(w, httpz.NewMsgError(rpc.FileExist, http.StatusBadRequest))
			return
		}

		// upload to the bucket..
		err = server.actions.Upload(storage.File{
			MimeType: header.ContentType,
			Body:     r.Body,
			FileSize: header.ContentLength,
		},
		storage.FileMetaData{
			Filename: args.Filename,
			IsPublic: true,
		},
		storage.Buckets{
			BucketName: bucket.BucketName,
			Endpoint:   bucket.Endpoint,
		})
		if err != nil {
			log.Print("do-file-upload: failed: " + err.Error())
			httpz.WriteError(w, httpz.NewMsgError(rpc.FileUploadFailed, http.StatusInternalServerError))
			return
		}

		// save to the DB
		file, err := server.filesRepository.CreateSpacesFile(r.Context(), args.Filename, header.ContentLength, header.ContentType, false, false, userId, bucket.ID)
		if err != nil {
			log.Print("do-file-upload: failed: " + err.Error())
			// rollback
			server.actions.Delete(storage.Buckets{
				BucketName: bucket.BucketName,
				Endpoint:   bucket.Endpoint,
			},
			args.Filename)
			httpz.WriteError(w, httpz.NewMsgError(rpc.FileUploadFailed, http.StatusInternalServerError))
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

func NewDOFileUploadServer(actions storage.Actions, client *ent.Client, authService rpc2.AuthorizationService) DOFileUploadServer {
	return DOFileUploadServer{
		actions:		  actions,
		authService:      authService,
		validator:        validator.NewDOFileUploadServerValidator(),
		filesRepository:   repository.NewFileRepository(client),
		bucketRepository: repository.NewDOBucketRepository(client),
	}
}