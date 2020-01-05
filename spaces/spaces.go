package spaces

import (
	"github.com/minio/minio-go"
	"log"
)


const (
	SpacesBucketEndpoint 	= "SPACES_BUCKET_ENDPOINT"
	SpacesBucketName        = "SPACES_BUCKET_NAME"
	SpacesAccessKey 		= "SPACES_ACCESS_KEY"
	SpacesSecretKey 		= "SPACES_SECRET_KEY"
)

type spaces struct {
	endpoint	 	string
	accessKey		string
	secretKey		string
	client 			*minio.Client
}

func NewSpaces(endpoint string, accessKey string, secretKey string) Bucket {
	client, err := minio.New(endpoint, accessKey, secretKey, true)
	if err != nil {
		log.Panic(err)
	}
	return &bucket{spaces:spaces{endpoint:endpoint, accessKey:accessKey, secretKey:secretKey, client:client}}
}