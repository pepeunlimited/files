package server

import (
	"github.com/pepeunlimited/files/internal/app/app1/ent"
	"github.com/pepeunlimited/files/rpcspaces"
	"time"
)

func toFile(files *ent.Files, spaces *ent.Spaces) *rpcspaces.File {
	file := &rpcspaces.File{
		Id:        int64(files.ID),
		Filename:  files.Filename,
		CreatedAt: files.CreatedAt.Format(time.RFC3339),
		UpdatedAt: files.CreatedAt.Format(time.RFC3339),
		MimeType:  files.MimeType,
		FileSize:  files.FileSize,
		UserId:    files.UserID,
		IsDraft:   files.IsDraft}
	if spaces == nil {
		return file
	}

	if spaces.CdnEndpoint == nil {
		file.FileUrl = "https://"+spaces.Endpoint+"/"+file.Filename
	} else {
		file.FileUrl = "https://"+*spaces.CdnEndpoint+"/"+file.Filename
	}

	file.SpacesId = int64(spaces.ID)
	return file
}