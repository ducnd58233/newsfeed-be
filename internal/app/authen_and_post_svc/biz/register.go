package authen_and_post_svc_biz

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/ducnd58233/newsfeed-be/internal/app/authen_and_post_svc/model"
	"github.com/ducnd58233/newsfeed-be/pkg/common"
	"golang.org/x/crypto/bcrypt"
)

type RegisterStore interface {
	Find(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*authen_and_post_svc_model.User, error)
	Create(ctx context.Context, data *authen_and_post_svc_model.UserCreate) error
}

type registerBiz struct {
	registerStore RegisterStore
}

func NewRegisterBiz(registerStore RegisterStore) *registerBiz {
	return &registerBiz{registerStore: registerStore}
}

func (biz *registerBiz) Register(ctx context.Context, data *authen_and_post_svc_model.UserCreate) (*authen_and_post_svc_model.UserResponse, error) {
	user, err := biz.registerStore.Find(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		return nil, authen_and_post_svc_model.ErrEmailExisted
	}

	if err != nil && err == common.ErrRecordNotFound {
		salt := common.GenSalt(common.DefaultSaltLength)
		passwordWithSalt := fmt.Sprintf("%s%s", data.Password, salt)
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordWithSalt), bcrypt.DefaultCost)
		if err != nil {
			return nil, fmt.Errorf("cannot hash password")
		}

		data.Salt = salt
		data.Password = hex.EncodeToString(hashedPassword)
		data.Status = 1

		if err := biz.registerStore.Create(ctx, data); err != nil {
			return nil, common.ErrCannotCreateEntity(authen_and_post_svc_model.EntityName, err)
		}

		return &authen_and_post_svc_model.UserResponse{
			Id: data.Id,
			Email: data.Email,
			LastName: data.LastName,
			FirstName: data.FirstName,
		}, nil
	}

	return nil, common.ErrDB(err)
}
