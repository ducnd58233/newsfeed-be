package service

import (
	"github.com/ducnd58233/newsfeed-be/configs"
	"github.com/ducnd58233/newsfeed-be/pkg/types/proto/pb/authen_and_post"
	"github.com/ducnd58233/newsfeed-be/pkg/types/proto/pb/newsfeed"
	"github.com/ducnd58233/newsfeed-be/pkg/client/authen_and_post"
	"github.com/ducnd58233/newsfeed-be/pkg/client/newsfeed"
)

type WebService struct {
	authenticateAndPostClient authen_and_post.AuthenticateAndPostClient
	newsfeedClient            newsfeed.NewsfeedClient
}

func NewWebService(conf *configs.WebConfig) (*WebService, error) {
	aapClient, err := authen_and_post_client.NewClient(conf.AuthenticateAndPost.Hosts)
	if err != nil {
		return nil, err
	}

	newsfeedClient, err := newsfeed_client.NewClient(conf.Newsfeed.Hosts)
	if err != nil {
		return nil, err
	}

	return &WebService{
		authenticateAndPostClient: aapClient,
		newsfeedClient: newsfeedClient,
	}, nil
}
