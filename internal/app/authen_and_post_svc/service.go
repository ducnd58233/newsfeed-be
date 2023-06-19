package authen_and_post_svc

import (
	"fmt"
	"log"

	"github.com/ducnd58233/newsfeed-be/configs"
	"github.com/ducnd58233/newsfeed-be/pkg/types/proto/pb/authen_and_post"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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
