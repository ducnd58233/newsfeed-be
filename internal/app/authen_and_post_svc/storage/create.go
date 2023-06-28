package authen_and_post_svc_storage

import (
	"context"

	"github.com/ducnd58233/newsfeed-be/internal/app/authen_and_post_svc/model"
	"github.com/ducnd58233/newsfeed-be/pkg/common"
)

func (s *sqlStore) Create(ctx context.Context, data *authen_and_post_svc_model.UserCreate) error {
	db := s.db.Begin().Table(data.TableName())

	if err := db.Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
