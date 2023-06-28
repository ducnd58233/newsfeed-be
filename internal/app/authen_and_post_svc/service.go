package authen_and_post_svc

import (
	"context"
	"fmt"
	"log"

	"github.com/ducnd58233/newsfeed-be/configs"
	"github.com/ducnd58233/newsfeed-be/internal/app/authen_and_post_svc/biz"
	"github.com/ducnd58233/newsfeed-be/internal/app/authen_and_post_svc/model"
	"github.com/ducnd58233/newsfeed-be/internal/app/authen_and_post_svc/storage"
	"github.com/ducnd58233/newsfeed-be/pkg/types/proto/pb/authen_and_post"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type AuthenAndPostStore interface {
	// CheckUserAuthentication(ctx context.Context, info *authen_and_post.UserInfo) (*authen_and_post.UserResult, error)
	CreateUser(ctx context.Context, info *authen_and_post.UserDetailInfo) (*authen_and_post.UserResult, error)
	// EditUser(ctx context.Context, info *authen_and_post.UserDetailInfo) (*authen_and_post.UserResult, error)
	// GetUserFollower(ctx context.Context, info *authen_and_post.UserInfo) (*authen_and_post.UserFollower, error)
	// GetPostDetail(ctx context.Context, request *authen_and_post.GetPostRequest) (*authen_and_post.Post, error)
}

func (aap *AuthenticateAndPostService) CreateUser(ctx context.Context, info *authen_and_post.UserRegisterInfo) (*authen_and_post.UserResult, error) {
	userInfo := &authen_and_post_svc_model.UserCreate{
		LastName:       info.LastName,
		FirstName:      info.FirstName,
		Password: info.Password,
		Email:          info.Email,
	}

	store := authen_and_post_svc_storage.NewSQLStore(aap.db)
	biz := authen_and_post_svc_biz.NewRegisterBiz(store)
	data, err := biz.Register(ctx, userInfo)

	if err != nil {
		return nil, err
	}

	userResp := &authen_and_post.UserDetailInfo{
		UserId: int64(data.Id),
		Email: data.Email,
		FirstName: data.FirstName,
		LastName: data.LastName,
	}

	return &authen_and_post.UserResult{
		Status: authen_and_post.UserStatus_OK,
		Info: userResp,
	}, nil
}

type AuthenticateAndPostService struct {
	authen_and_post.UnimplementedAuthenticateAndPostServer
	db    *gorm.DB
	redis *redis.Client
}

func NewAuthenticateAndPostService(conf *configs.AuthenticateAndPostConfig) (*AuthenticateAndPostService, error) {
	db, err := gorm.Open(mysql.New(conf.MySQL), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}

	rd := redis.NewClient(&conf.Redis)
	if rd == nil {
		return nil, fmt.Errorf("cannot init redis client")
	}

	return &AuthenticateAndPostService{
		db:    db,
		redis: rd,
	}, nil
}
