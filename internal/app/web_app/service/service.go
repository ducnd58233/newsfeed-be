package service

import (
	"fmt"
	"net/http"

	"github.com/ducnd58233/newsfeed-be/configs"
	"github.com/ducnd58233/newsfeed-be/internal/pkg/types/web_app"
	"github.com/ducnd58233/newsfeed-be/pkg/client/authen_and_post"
	"github.com/ducnd58233/newsfeed-be/pkg/client/newsfeed"
	"github.com/ducnd58233/newsfeed-be/pkg/types/proto/pb/authen_and_post"
	"github.com/ducnd58233/newsfeed-be/pkg/types/proto/pb/newsfeed"
	"github.com/gin-gonic/gin"
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
		newsfeedClient:            newsfeedClient,
	}, nil
}

func (svc *WebService) CheckUserNameAndPassword(ctx *gin.Context) {
	var jsonRequest types.LoginRequest
	err := ctx.ShouldBindJSON(&jsonRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	authentication, err := svc.authenticateAndPostClient.CheckUserAuthentication(ctx, &authen_and_post.UserInfo{
		Email:        jsonRequest.Email,
		UserPassword: jsonRequest.Password,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if authentication.GetStatus() == authen_and_post.UserStatus_OK {
		ctx.Status(http.StatusOK)
		// change this later
		ctx.SetCookie("session_id", fmt.Sprintf("%d", authentication.Info.UserId), 0, "", "", false, false)
	}
}

func (svc *WebService) CreateUser(ctx *gin.Context) {
	var jsonRequest types.RegisterRequest
	err := ctx.ShouldBindJSON(&jsonRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	user, err := svc.authenticateAndPostClient.CreateUser(ctx, &authen_and_post.UserRegisterInfo{
		Email:     jsonRequest.Email,
		Password:  jsonRequest.Password,
		FirstName: jsonRequest.FirstName,
		LastName:  jsonRequest.LastName,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if user.GetStatus() == authen_and_post.UserStatus_OK {
		ctx.Status(http.StatusOK)
		// change this later
		ctx.SetCookie("session_id", fmt.Sprintf("%d", user.Info.UserId), 0, "", "", false, false)
	}
}
