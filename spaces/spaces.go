package spaces

import (
	"context"
	"github.com/digitalocean/godo"
	"github.com/minio/minio-go"
	"github.com/pepeunlimited/files/storage"
	"golang.org/x/oauth2"
	"log"
)


const (
	SpacesAccessKey 		= "SPACES_ACCESS_KEY"
	SpacesSecretKey 		= "SPACES_SECRET_KEY"
	DoAccessToken			= "DO_ACCESS_TOKEN"
)

func NewBucket(endpoint string, accessKey string, secretKey string, bucketName string) storage.Bucket {
	return NewBucketCDN(endpoint, accessKey, secretKey, bucketName, nil)
}

func NewBucketCDN(endpoint string, accessKey string, secretKey string, bucketName string, accessToken *string) storage.Bucket {
	client, err := minio.New(endpoint, accessKey, secretKey, true)
	if err != nil {
		log.Panic(err)
	}
	spaces := spaces{endpoint: endpoint, accessKey: accessKey, secretKey: secretKey}
	b := &bucket{spaces: spaces, bucketName: bucketName, client: client}
	if accessToken == nil {
		return b
	}
	b.doClient = NewDoClient(*accessToken)
	return b
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

func NewDoClient(accessToken string) *godo.Client {
	oauth := oauth2.NewClient(context.Background(), &tokenSource{accessToken: accessToken})
	return godo.NewClient(oauth)
}