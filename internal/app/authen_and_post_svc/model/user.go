package authen_and_post_svc_model

import (
	"errors"

	"github.com/ducnd58233/newsfeed-be/pkg/common"
)

var EntityName = "User"

type User struct {
	common.SQLModel `json:",inline"`
	Email           string `json:"email" gorm:"column:email;"`
	Password        string `json:"-" gorm:"column:password;"`
	LastName        string `json:"last_name" gorm:"column:last_name;"`
	FirstName       string `json:"first_name" gorm:"column:first_name;"`
	Salt            string `json:"-" gorm:"column:salt;"`
	Status          int    `json:"-" gorm:"column:status;"`
}

func (User) TableName() string {
	return "users"
}

type UserCreate struct {
	common.SQLModel `json:",inline"`
	Email           string `json:"email" gorm:"column:email;"`
	Password        string `json:"-" gorm:"column:password;"`
	LastName        string `json:"last_name" gorm:"column:last_name;"`
	FirstName       string `json:"first_name" gorm:"column:first_name;"`
	Salt            string `json:"-" gorm:"column:salt;"`
	Status          int    `json:"-" gorm:"column:status;"`
}

func (UserCreate) TableName() string {
	return "users"
}

type UserResponse struct {
	Id        uint    `json:"id" gorm:"column:id;"`
	Email     string `json:"email" gorm:"column:email;"`
	LastName  string `json:"last_name" gorm:"column:last_name;"`
	FirstName string `json:"first_name" gorm:"column:first_name;"`
}

func (UserResponse) TableName() string {
	return "users"
}

var (
	ErrUsernameOrPasswordInvalid = common.NewCustomError(
		errors.New("username or password invalid"),
		"username or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)

	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)
)
