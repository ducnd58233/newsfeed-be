package authen_and_post_svc_storage

import (
	"context"

	"github.com/ducnd58233/newsfeed-be/internal/app/authen_and_post_svc/model"
	"github.com/ducnd58233/newsfeed-be/pkg/common"
	"gorm.io/gorm"
)

func (s *sqlStore) Find(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*authen_and_post_svc_model.User, error) {
	db := s.db.Table(authen_and_post_svc_model.User{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var user authen_and_post_svc_model.User

	if err := db.Where(conditions).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &user, nil
}
