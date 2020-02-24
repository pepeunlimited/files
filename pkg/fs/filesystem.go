package fs

import (
	"context"
	"github.com/digitalocean/godo"
	"github.com/minio/minio-go"
	"github.com/twitchtv/twirp"
	"golang.org/x/oauth2"
	"io"
	"log"
	"strings"
)

const (
	FsEndpoint					= "FS_ENDPOINT"

	SpacesAccessKey 			= "SPACES_ACCESS_KEY"
	SpacesSecretKey 			= "SPACES_SECRET_KEY"
	DoAccessToken				= "DO_ACCESS_TOKEN"

	GoogleCloudStorageAccessKey = "GCS_ACCESS_KEY"
	GoogleCloudStorageSecretKey = "GCS_SECRET_KEY"
)

type FileSystem interface {
	CdnEndpoint(bucket string) (string, error)
	Endpoint() 												string
	DeleteBucket(bucket string) 		 				 	error
	CreateBucket(ctx context.Context, bucket string) 		error
	BucketExist(bucket string) 			 	 				(bool, error)
	UploadFile(ctx context.Context, file File, meta FileMetaData, bucket string)  error
	DeleteFile(filename string, bucket string) 				error
	UpdateFile(file File, meta FileMetaData)	 			    (*File, error)
	GetFile(filename string, bucket string)	 				([]byte, error)
	GetMetadata(filename string, bucket string)			 	(*FileMetaData, error)
	GetBuckets()											([]*Bucket, error)
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

type Bucket struct {
	BucketName string
	Endpoint   string
	IsCDN	   bool
}

type filesystem struct {
	digitalocean 	*godo.Client
	client 		   	*minio.Client
	isDigitalOcean 	bool
	endpoint 		string
}

func (f filesystem) CdnEndpoint(bucket string) (string, error) {
	if f.isDigitalOcean {
		endpoint := strings.Split(f.Endpoint(), ".")
		if len(endpoint) != 3 {
			return "", twirp.InvalidArgumentError("endpoint", "invalid_endpoint")
		}
		return bucket+"."+endpoint[0]+".cdn."+endpoint[1]+"."+endpoint[2], nil
	}
	return f.endpoint, nil
}


func (f filesystem) Endpoint() string {
	return f.endpoint
}

func (f filesystem) GetBuckets() ([]*Bucket, error) {
	buckets, err := f.client.ListBuckets()
	if err != nil {
		return nil, err
	}
	list := make([]*Bucket, 0)
	for _, bucket := range buckets {
		list = append(list, &Bucket{
			BucketName: bucket.Name,
			Endpoint:   f.endpoint,
			//TODO: isCDN
		})
	}
	return list, nil
}

func (f filesystem) DeleteBucket(bucket string) error {
	return f.client.RemoveBucket(bucket)
}

func (f filesystem) CreateBucket(ctx context.Context, bucket string) error {
	if err := f.client.MakeBucket(bucket, ""); err != nil {
		return err
	}
	if f.isDigitalOcean {
		_, _, err := f.digitalocean.CDNs.Create(ctx, &godo.CDNCreateRequest{
			Origin: bucket+"."+f.endpoint,
			TTL:    3600,
		})
		if err != nil {
			if err := f.DeleteBucket(bucket); err != nil {
				return err
			}
			return err
		}
	}
	return nil
}

func (f filesystem) BucketExist(bucket string) (bool, error) {
	return f.client.BucketExists(bucket)
}

func (f filesystem) UploadFile(ctx context.Context, file File, meta FileMetaData, bucket string) error {
	options := minio.PutObjectOptions{}
	options.ContentType = file.MimeType
	options.ContentDisposition = "inline; filename=\"" + meta.Filename + "\""
	options.UserMetadata = make(map[string]string)
	if meta.IsPublic {
		options.UserMetadata["x-amz-acl"] = "public-read"
	}
	if _, err := f.client.PutObjectWithContext(ctx, bucket, meta.Filename, file.Body, file.FileSize, options); err != nil {
		return err
	}
	return nil
}

func (f filesystem) DeleteFile(filename string, bucket string) error {
	if err := f.client.RemoveObject(bucket, filename); err != nil {
		return err
	}
	return nil
}

func (f filesystem) UpdateFile(file File, meta FileMetaData) (*File, error) {
	panic("implement me")
}

func (f filesystem) GetFile(filename string, bucket string) ([]byte, error) {
	file, err := f.client.GetObject(bucket, filename, minio.GetObjectOptions{})
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

func (f filesystem) GetMetadata(filename string, bucket string) (*FileMetaData, error) {
	_, err := f.client.GetObject(bucket, filename, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func NewDigitalOcean(endpoint string, accessKey string, secretKey string, accessToken string) FileSystem {
	f := filesystem{endpoint:endpoint}
	client, err := minio.New(endpoint, accessKey, secretKey, true)
	if err != nil {
		log.Panic(err)
	}
	f.isDigitalOcean = true
	f.digitalocean = newDoClient(accessToken)
	f.client = client
	return f
}

func NewGoogleCloud(endpoint string, accessKey string, secretKey string) FileSystem {
	f := filesystem{endpoint:endpoint}
	client, err := minio.New(endpoint, accessKey, secretKey, true)
	if err != nil {
		log.Panic(err)
	}
	f.isDigitalOcean = false
	f.client = client
	return filesystem{}
}




type tokenSource struct {
	accessToken string
}

func (t *tokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.accessToken,
	}
	return token, nil
}

func newDoClient(accessToken string) *godo.Client {
	oauth := oauth2.NewClient(context.Background(), &tokenSource{accessToken: accessToken})
	return godo.NewClient(oauth)
}