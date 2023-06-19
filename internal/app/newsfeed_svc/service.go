package newsfeed_svc

import (
	"fmt"
	"log"

	"github.com/ducnd58233/newsfeed-be/configs"
	"github.com/ducnd58233/newsfeed-be/pkg/types/proto/pb/newsfeed"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type NewsfeedService struct {
	newsfeed.UnimplementedNewsfeedServer
	db    *gorm.DB
	redis *redis.Client
}

func NewNewsfeedService(conf *configs.NewsfeedConfig) (*NewsfeedService, error) {
	db, err := gorm.Open(mysql.New(conf.MySQL), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
		return nil, err
	}

	rd := redis.NewClient(&conf.Redis)
	if rd == nil {
		return nil, fmt.Errorf("cannot init redis client")
	}

	return &NewsfeedService{
		db:    db,
		redis: rd,
	}, nil
}
