package upload

import (
	"github.com/pepeunlimited/authentication-twirp/pkg/rpc/auth"
	"github.com/pepeunlimited/files/internal/pkg/ent"
	"github.com/pepeunlimited/files/internal/pkg/mysql/bucket"
	"github.com/pepeunlimited/files/internal/pkg/mysql/file"
	"github.com/pepeunlimited/files/pkg/fs"
	"github.com/pepeunlimited/files/pkg/rpc/files"
	"github.com/pepeunlimited/microservice-kit/httpz"
	"github.com/twitchtv/twirp"
	"log"
	"net/http"
)

const (
	// version
	VersionV1            = "/v1"
	UploadV1Files string = UploadPath+VersionV1+FilesPath
)

type UploadServer struct {
	validator      UploadServerValidator
	files          file.FileRepository
	buckets        bucket.BucketRepository
	authentication auth.AuthenticationService
	filesystem     fs.FileSystem
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
		user, err := server.authentication.VerifyAccessToken(r.Context(), &auth.VerifyAccessTokenParams{
			AccessToken: header.Authorization,
		})
		if err != nil {
			switch err.(twirp.Error).Msg() {
			case auth.AccessTokenExpired:
				httpz.WriteError(w,httpz.NewMsgError(auth.AccessTokenExpired, http.StatusUnauthorized))
			case auth.AccessTokenMalformed:
				httpz.WriteError(w,httpz.NewMsgError(auth.AccessTokenMalformed, http.StatusBadRequest))
			case auth.AccessTokenUnknownError:
				httpz.WriteError(w,httpz.NewMsgError(auth.AccessTokenUnknownError, http.StatusInternalServerError))
			default:
				log.Print("buckets-upload: failed: "+err.Error())
				httpz.WriteError(w,httpz.NewMsgError(files.FileUploadFailed, http.StatusInternalServerError))
			}
			return
		}
		userId := user.UserId
		buckets,_, err := server.buckets.GetBuckets(r.Context(), 0, 20)
		if err != nil {
			log.Print("buckets-upload: failed: "+err.Error())
			httpz.WriteError(w, httpz.NewMsgError(files.FileUploadFailed, http.StatusInternalServerError))
			return
		}
		if len(buckets) == 0 {
			log.Print("buckets-upload: missing buckets!")
			httpz.WriteError(w, httpz.NewMsgError(files.FileUploadFailed, http.StatusInternalServerError))
			return
		}
		bucket, err := server.buckets.GetBucketsByEndpoint(r.Context(), server.filesystem.Endpoint())
		if err != nil {
			log.Print("buckets-upload: failed: "+err.Error())
			httpz.WriteError(w, httpz.NewMsgError(files.FileUploadFailed, http.StatusInternalServerError))
			return
		}
		exist, err := server.files.ExistInBucket(r.Context(), args.Filename, bucket.ID)
		if err != nil {
			log.Print("buckets-upload: failed: "+err.Error())
			httpz.WriteError(w, httpz.NewMsgError(files.FileUploadFailed, http.StatusInternalServerError))
			return
		}
		if *exist {
			httpz.WriteError(w, httpz.NewMsgError(files.FileExist, http.StatusBadRequest))
			return
		}
		// upload to the bucket..
		err = server.filesystem.UploadFile(
			r.Context(),
			fs.File{
			MimeType: header.ContentType,
			Body:     r.Body,
			FileSize: header.ContentLength,
		},
		fs.FileMetaData{
			Filename: args.Filename,
			IsPublic: true,
		},
		bucket.Name)
		if err != nil {
			log.Print("buckets-upload: failed: " + err.Error())
			httpz.WriteError(w, httpz.NewMsgError(files.FileUploadFailed, http.StatusInternalServerError))
			return
		}
		// save to the DB
		file, err := server.files.CreateFile(r.Context(), args.Filename, header.ContentLength, header.ContentType, false, false, userId, bucket.ID)
		if err != nil {
			log.Print("buckets-upload: failed: " + err.Error())
			// rollback
			if err := server.filesystem.DeleteFile(args.Filename, bucket.Name); err != nil {
				log.Print("buckets-upload: failed to delete: " + err.Error() + ", filename: "+args.Filename+", bucket: "+bucket.Name)
			}
			httpz.WriteError(w, httpz.NewMsgError(files.FileUploadFailed, http.StatusInternalServerError))
			return
		}
		httpz.WriteOk(w, UploadFile{
			Endpoint:    bucket.Endpoint,
			CDNEndpoint: *bucket.CdnEndpoint,
			FileId:      int64(file.ID),
			Filename:    args.Filename,
			URI:		 "https://"+*bucket.CdnEndpoint+"/"+args.Filename,
		})
	})
}

func NewUploadServer(client *ent.Client, authentication auth.AuthenticationService, fs fs.FileSystem) UploadServer {
	return UploadServer{
		authentication: authentication,
		validator:      NewSpacesUploadServerValidator(),
		files:          file.NewFileRepository(client),
		buckets:        bucket.New(client),
		filesystem:     fs,
	}
}

