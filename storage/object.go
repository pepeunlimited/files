package storage

import (
	"github.com/minio/minio-go"
	"io"
)

type Object interface {
	Create(file File, meta FileMetaData)  Object
	Delete(filename string) 				 Object
	Update()							 Object
	Get(filename string)				 	 ([]byte, error)
	GetMetadata(filename string)			 (*minio.ObjectInfo, error)
	Execute()							 error
}

type FileMetaData struct {
	Filename 	string
	IsPublic 	bool
}

type File struct {
	MimeType 	string
	Body 		io.Reader
	FileSize 	int64
}