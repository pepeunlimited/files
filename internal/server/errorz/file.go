package errorz

import (
	"github.com/pepeunlimited/files/internal/pkg/mysql/bucket"
	"github.com/pepeunlimited/files/internal/pkg/mysql/file"
	"github.com/pepeunlimited/files/pkg/rpc/files"
	"github.com/twitchtv/twirp"
	"log"
)

func File(err error) error {
	switch err {
	case file.ErrFileNotExist:
		return twirp.NotFoundError(files.FileNotFound)
	case bucket.ErrBucketsNotExist:
		return twirp.NotFoundError(files.BucketNotFound)
	}
	log.Print("file-service: unknown: "+err.Error())
	return twirp.InternalErrorWith(err)
}