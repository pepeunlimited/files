package spaces

import (
	"fmt"
	"github.com/minio/minio-go"
	"io"
	"log"
)

type Files interface {
	Create(file File, meta FileMetaData)  Files
	Delete(filename string) 				 Files
	Update()							 Files
	Get(filename string)				 	 ([]byte, error)
	Execute()							 error
}

type files struct {
	bucket 		bucket
	executor  	map[int]interface{}
	order       int
}

type FileMetaData struct {
	filename 	string
	isPublic 	bool
}

type File struct {
	mimeType 	string
	body 		io.Reader
	fileSize 	int64
}

type create struct {
	fileMetaData FileMetaData
	file 		File
}
type delete struct {
	filename 	string
}

type update struct {}

func (f *files) Create(file File, meta FileMetaData) Files {
	f.executor[f.order] = create{file:file,fileMetaData:meta}
	f.order++
	return f
}

func (f *files) Delete(filename string) Files {
	f.executor[f.order] = delete{filename:filename}
	f.order++
	return f
}

func (f *files) Update() Files {
	f.executor[f.order] = update{}
	f.order++
	return f
}

func (f *files) Get(filename string) ([]byte, error) {
	file, err := f.bucket.spaces.client.GetObject(f.bucket.name, filename, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	bytes := make([]byte, stat.Size)
	_, err = file.Read(bytes)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return bytes, nil
}

func (f *files) Execute() error {
	if f.bucket.isCreate {
		if err := f.bucket.spaces.client.MakeBucket(f.bucket.name, ""); err != nil {
			return err
		}
	}
	if f.bucket.isFiles {
		//TODO: validate does the bucket exist
	}
	for k := range f.executor {
		ec := f.executor[k]
		switch ec.(type) {
		case create:
			crt := ec.(create)
			options := minio.PutObjectOptions{}
			options.ContentType = crt.file.mimeType
			options.UserMetadata = make(map[string]string)
			if crt.fileMetaData.isPublic {
				options.UserMetadata["x-amz-acl"] = "public-read"
			}
			if _, err := f.bucket.spaces.client.PutObject(f.bucket.name, crt.fileMetaData.fileName, crt.file.body, crt.file.fileSize, options); err != nil {
				return err
			}
		case update:
			log.Panic("not implemented")
		case delete:
			delete := ec.(delete)
			if err := f.bucket.spaces.client.RemoveObject(f.bucket.name, delete.filename); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unknown type for the executor")
		}
	}
	return nil
}