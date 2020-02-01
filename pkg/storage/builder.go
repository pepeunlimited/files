package storage

import (
	"context"
	"fmt"
	"github.com/digitalocean/godo"
	"github.com/minio/minio-go"
	"io"
	"log"
)

const (
	SpacesAccessKey 			= "SPACES_ACCESS_KEY"
	SpacesSecretKey 			= "SPACES_SECRET_KEY"
	DoAccessToken				= "DO_ACCESS_TOKEN"

	GoogleCloudStorageAccessKey = "GOOGLE_CLOUD_STORAGE_ACCESS_KEY"
	GoogleCloudStorageSecretKey = "GOOGLE_CLOUD_STORAGE_SECRET_KEY"
)

func newClient(endpoint string, accessKey string, secretKey string) *minio.Client {
	client, err := minio.New(endpoint, accessKey, secretKey, true)
	if err != nil {
		log.Panic(err)
	}
	return client
}

func NewSpacesBuilder(endpoint string, accessKey string, secretKey string, bucketName string, accessToken *string) Bucket {
	client := newClient(endpoint, accessKey, secretKey)
	service := service{endpoint: endpoint, accessKey: accessKey, secretKey: secretKey}
	b := &bucket{service: service, bucketName: bucketName, client: client, isSpaces:true, isGKE:false}
	if accessToken == nil {
		return b
	}
	b.spaces = &spaces{doClient: NewDoClient(*accessToken)}
	return b
}

func NewGoogleCloudStorageBuilder(endpoint string, accessKey string, secretKey string, bucketName string) Bucket {
	return nil
}

type service struct {
	endpoint	 	string
	accessKey		string
	secretKey		string
}

type bucket struct {
	isSpaces    bool
	isGKE       bool
	isCreate 	bool
	isFiles    	bool
	bucketName  string
	service 	service
	client 		*minio.Client
	spaces      *spaces
	gke         *gke
}

type spaces struct {
	doClient    *godo.Client
}

type gke struct {}

func (b *bucket) Exist() (bool, error) {
	return b.client.BucketExists(b.bucketName)
}

func (b *bucket) Delete() error {
	return b.client.RemoveBucket(b.bucketName)
}

func (b *bucket) Files() Object {
	executor := make(map[int]interface{})
	b.isFiles = true
	b.isCreate = false
	return &files{bucket:*b, executor:executor, order: 0}
}

func (b *bucket) Create() Object {
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
	fileMetaData FileMetaData
	file         File
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

func (f *files) Create(file File, meta FileMetaData) Object {
	f.executor[f.order] = create{file:file,fileMetaData:meta}
	f.order++
	return f
}

func (f *files) Delete(filename string) Object {
	f.executor[f.order] = delete{filename:filename}
	f.order++
	return f
}

func (f *files) Update() Object {
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

func (f *files) Execute(ctx context.Context) error {
	if f.bucket.isCreate {
		if err := f.bucket.client.MakeBucket(f.bucket.bucketName, ""); err != nil {
			return err
		}
		if f.bucket.isSpaces {
			if f.bucket.spaces != nil {
				_, _, err := f.bucket.spaces.doClient.CDNs.Create(ctx, &godo.CDNCreateRequest{
					Origin: f.bucket.bucketName+"."+f.bucket.service.endpoint,
					TTL:    3600,
				})
				if err != nil {
					return err
				}
			}
		} else {

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