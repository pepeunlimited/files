package spaces

import (
	"context"
	"fmt"
	"github.com/digitalocean/godo"
	"github.com/minio/minio-go"
	"github.com/pepeunlimited/files/storage"
	"io"
	"log"
)

type spaces struct {
	endpoint	 	string
	accessKey		string
	secretKey		string
}

type bucket struct {
	isCreate 	bool
	isFiles    	bool
	bucketName  string
	spaces 		spaces
	client 		*minio.Client
	doClient    *godo.Client
}

func (b *bucket) Exist() (bool, error) {
	return b.client.BucketExists(b.bucketName)
}

func (b *bucket) Delete() error {
	return b.client.RemoveBucket(b.bucketName)
}

func (b *bucket) Files() storage.Object {
	executor := make(map[int]interface{})
	b.isFiles = true
	b.isCreate = false
	return &files{bucket:*b, executor:executor, order: 0}
}

func (b *bucket) Create() storage.Object {
	executor := make(map[int]interface{})
	b.isCreate = true
	b.isFiles = false
	return &files{bucket:*b, executor:executor, order: 0}
}

type files struct {
	bucket 		bucket
	executor  	map[int]interface{}
	order       int
}

type create struct {
	fileMetaData storage.FileMetaData
	file 		storage.File
}
type delete struct {
	filename 	string
}

type update struct {}

func (f *files) GetMetadata(filename string) (*minio.ObjectInfo, error) {
	file, err := f.get(filename)
	if err != nil {
		return nil, err
	}
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	return &stat, nil
}

func (f *files) Create(file storage.File, meta storage.FileMetaData) storage.Object {
	f.executor[f.order] = create{file:file,fileMetaData:meta}
	f.order++
	return f
}

func (f *files) Delete(filename string) storage.Object {
	f.executor[f.order] = delete{filename:filename}
	f.order++
	return f
}

func (f *files) Update() storage.Object {
	f.executor[f.order] = update{}
	f.order++
	return f
}

func (f *files) get(filename string) (*minio.Object, error) {
	return f.bucket.client.GetObject(f.bucket.bucketName, filename, minio.GetObjectOptions{})
}

func (f *files) Get(filename string) ([]byte, error) {
	file, err := f.get(filename)
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
		if err := f.bucket.client.MakeBucket(f.bucket.bucketName, ""); err != nil {
			return err
		}
		if f.bucket.doClient != nil {
			_, _, err := f.bucket.doClient.CDNs.Create(context.Background(), &godo.CDNCreateRequest{
				Origin: f.bucket.bucketName+"."+f.bucket.spaces.endpoint,
				TTL:    3600,
			})
			if err != nil {
				return err
			}
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
			options.ContentType = crt.file.MimeType
			options.ContentDisposition = "inline; filename=\"" + crt.fileMetaData.Filename + "\""
			options.UserMetadata = make(map[string]string)
			if crt.fileMetaData.IsPublic {
				options.UserMetadata["x-amz-acl"] = "public-read"
			}
			if _, err := f.bucket.client.PutObject(f.bucket.bucketName, crt.fileMetaData.Filename, crt.file.Body, crt.file.FileSize, options); err != nil {
				return err
			}
		case update:
			log.Panic("not implemented")
		case delete:
			delete := ec.(delete)
			if err := f.bucket.client.RemoveObject(f.bucket.bucketName, delete.filename); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unknown type for the executor")
		}
	}
	return nil
}