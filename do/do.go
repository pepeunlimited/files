package do

import (
	"context"
	"github.com/digitalocean/godo"
	"golang.org/x/oauth2"
)

const (
	DoAccessToken    	= "DO_ACCESS_TOKEN"
)

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