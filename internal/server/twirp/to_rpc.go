package twirp

import (
	"github.com/pepeunlimited/files/internal/pkg/ent"
	"github.com/pepeunlimited/files/pkg/filesrpc"
	"time"
)

func toFile(files *ent.Files, buckets *ent.Buckets) *filesrpc.File {
	file := &filesrpc.File{
		Id:        int64(files.ID),
		Filename:  files.Filename,
		CreatedAt: files.CreatedAt.Format(time.RFC3339),
		UpdatedAt: files.CreatedAt.Format(time.RFC3339),
		MimeType:  files.MimeType,
		FileSize:  files.FileSize,
		UserId:    files.UserID,
		IsDraft:   files.IsDraft}
	if buckets == nil {
		return file
	}

	if buckets.CdnEndpoint == nil {
		file.FileUrl = "https://"+buckets.Endpoint+"/"+file.Filename
	} else {
		file.FileUrl = "https://"+*buckets.CdnEndpoint+"/"+file.Filename
	}

	file.SpacesId = int64(buckets.ID)
	return file
}