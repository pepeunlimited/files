package server

import (
	"github.com/pepeunlimited/authentication-twirp/rpcauth"
	"github.com/pepeunlimited/files/internal/app/app1/ent"
	"github.com/pepeunlimited/files/internal/app/app1/filerepo"
	"github.com/pepeunlimited/files/internal/app/app1/spacesrepo"
	"github.com/pepeunlimited/files/internal/app/app1/upload"
	"github.com/pepeunlimited/files/internal/app/app1/validator"
	"github.com/pepeunlimited/files/spacesrpc"
	"github.com/pepeunlimited/files/storage"
	"github.com/pepeunlimited/microservice-kit/httpz"
	"github.com/twitchtv/twirp"
	"log"
	"math/rand"
	"net/http"
)

const (
	// DigitalOcean's
	SpacesPath 				   = "/spaces"
	// version
	SpacesVersionV1            = "/v1"
	UploadSpacesV1Files string = UploadPath+ SpacesPath + SpacesVersionV1 +FilesPath
)

type SpacesUploadServer struct {
	validator   	validator.SpacesUploadServerValidator
	actions     	storage.Actions
	files        	filerepo.FileRepository
	spaces      	spacesrepo.SpacesRepository
	authentication  rpcauth.AuthenticationService
}

// https://phil.tech/api/2016/01/04/http-rest-api-file-uploads/
func (server SpacesUploadServer) UploadSpacesV1Files() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header, args, err := server.validator.UploadSpacesV1Files(r.Header, r.Body)
		if err != nil {
			httpz.WriteError(w, err)
			return
		}
		// validate the access token..
		user, err := server.authentication.VerifyAccessToken(r.Context(), &rpcauth.VerifyAccessTokenParams{
			AccessToken: header.Authorization,
		})
		if err != nil {
			if spacesrpc.IsReason(err.(twirp.Error), rpcauth.AccessTokenExpired) {
				httpz.WriteError(w,httpz.NewMsgError(rpcauth.AccessTokenExpired, http.StatusUnauthorized))
			} else if spacesrpc.IsReason(err.(twirp.Error), rpcauth.AccessTokenMalformed) {
				httpz.WriteError(w,httpz.NewMsgError(rpcauth.AccessTokenMalformed, http.StatusBadRequest))
			} else if spacesrpc.IsReason(err.(twirp.Error), rpcauth.AccessTokenUnknownError) {
				httpz.WriteError(w,httpz.NewMsgError(rpcauth.AccessTokenUnknownError, http.StatusInternalServerError))
			} else {
				log.Print("spaces-upload: failed: "+err.Error())
				httpz.WriteError(w,httpz.NewMsgError(spacesrpc.FileUploadFailed, http.StatusInternalServerError))
			}
			return
		}
		userId := user.UserId

		buckets,_, err := server.spaces.GetSpaces(r.Context(), 0, 20)
		if err != nil {
			log.Print("spaces-upload: failed: "+err.Error())
			httpz.WriteError(w, httpz.NewMsgError(spacesrpc.FileUploadFailed, http.StatusInternalServerError))
			return
		}
		if len(buckets) == 0 {
			log.Print("spaces-upload: missing buckets!")
			httpz.WriteError(w, httpz.NewMsgError(spacesrpc.FileUploadFailed, http.StatusInternalServerError))
			return
		}

		min := 0
		max := len(buckets)
		random := rand.Intn(max - min) + min
		bucket := buckets[random]

		exist, err := server.files.ExistInSpaces(r.Context(), args.Filename, bucket.ID)
		if err != nil {
			log.Print("spaces-upload: failed: "+err.Error())
			httpz.WriteError(w, httpz.NewMsgError(spacesrpc.FileUploadFailed, http.StatusInternalServerError))
			return
		}

		if *exist {
			httpz.WriteError(w, httpz.NewMsgError(spacesrpc.FileExist, http.StatusBadRequest))
			return
		}

		// upload to the spaces..
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
			BucketName: bucket.Name,
			Endpoint:   bucket.Endpoint,
		})
		if err != nil {
			log.Print("spaces-upload: failed: " + err.Error())
			httpz.WriteError(w, httpz.NewMsgError(spacesrpc.FileUploadFailed, http.StatusInternalServerError))
			return
		}

		// save to the DB
		file, err := server.files.CreateSpacesFile(r.Context(), args.Filename, header.ContentLength, header.ContentType, false, false, userId, bucket.ID)
		if err != nil {
			log.Print("spaces-upload: failed: " + err.Error())
			// rollback
			server.actions.Delete(storage.Buckets{
				BucketName: bucket.Name,
				Endpoint:   bucket.Endpoint,
			},
			args.Filename)
			httpz.WriteError(w, httpz.NewMsgError(spacesrpc.FileUploadFailed, http.StatusInternalServerError))
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

func NewSpacesUploadServer(actions storage.Actions, client *ent.Client, authentication rpcauth.AuthenticationService) SpacesUploadServer {
	return SpacesUploadServer{
		actions:    	 	actions,
		authentication: 	authentication,
		validator:   		validator.NewSpacesUploadServerValidator(),
		files:       		filerepo.NewFileRepository(client),
		spaces:      		spacesrepo.NewSpacesRepository(client),
	}
}