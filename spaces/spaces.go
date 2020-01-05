package spaces

import (
	"github.com/minio/minio-go"
	"log"
)


const (
	SpacesEndpoint 		= "SPACES_ENDPOINT"
	SpacesAccessKey 	= "SPACES_ACCESS_KEY"
	SpacesSecretKey 	= "SPACES_SECRET_KEY"
	SpacesBucketName    = "SPACES_BUCKET_NAME"
)

type spaces struct {
	//	region	 		string
	endpoint	 	string
	accessKey		string
	secretKey		string
	client 			*minio.Client
}

func NewSpaces(endpoint string, accessKey string, secretKey string) Bucket {
	// Initiate a client using DigitalOcean Spaces.
	client, err := minio.New(endpoint, accessKey, secretKey, true)
	if err != nil {
		log.Panic(err)
	}

	return &bucket{spaces:spaces{endpoint:endpoint, accessKey:accessKey, secretKey:secretKey, client:client}}
}