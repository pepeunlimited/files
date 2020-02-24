package twirp

import (
	"github.com/pepeunlimited/files/internal/pkg/ent"
	"github.com/pepeunlimited/files/pkg/rpc/files"
	"time"
)

func toFile(fromFile *ent.File, fromBuckets *ent.Bucket) *files.File {
	file := &files.File{
		Id:        int64(fromFile.ID),
		Filename:  fromFile.Filename,
		CreatedAt: fromFile.CreatedAt.Format(time.RFC3339),
		UpdatedAt: fromFile.CreatedAt.Format(time.RFC3339),
		MimeType:  fromFile.MimeType,
		FileSize:  fromFile.FileSize,
		UserId:    fromFile.UserID,
		IsDraft:   fromFile.IsDraft}
	if fromBuckets == nil {
		return file
	}

	if fromBuckets.CdnEndpoint == nil {
		file.FileUrl = "https://"+fromBuckets.Endpoint+"/"+file.Filename
	} else {
		file.FileUrl = "https://"+*fromBuckets.CdnEndpoint+"/"+file.Filename
	}
	file.SpacesId = int64(fromBuckets.ID)
	return file
}