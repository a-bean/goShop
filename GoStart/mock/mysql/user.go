package mysql

import (
	"context"

	"GoStart/mock"
	"gorm.io/gorm"
)

type user struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *user {
	return &user{db: db}
}

func (u *user) GetUserByMobile(ctx context.Context, mobile string) (mock.User, error) {
	var user mock.User
	_ = u.db.Where(&mock.User{Mobile: mobile}).First(&user)
	return user, nil
}

var _ mock.UserData = &user{}
