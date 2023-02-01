package mysql

import (
	"context"
	"regexp"

	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"

	umock "GoStart/mock"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetUserByMobile(t *testing.T) {
	//注入
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("init sqlmock: %v", err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			t.Fatalf("close sqlmock: %v", err)
		}
	}(db)

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("open gorm: %v", err)
	}
	mobile := "18"

	//期望
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`mobile` = ? ORDER BY `users`.`mobile` LIMIT 1")).
		WithArgs(mobile).WillReturnRows(
		sqlmock.NewRows([]string{"mobile", "password", "nick_name"}).AddRow("18787878878", "sjojj", "bobby"))
	mock.ExpectClose()
	//调用
	userData := NewUser(gormDB)
	user, err := userData.GetUserByMobile(context.Background(), mobile)
	assert.Nil(t, err)

	expUser := umock.User{
		Mobile:   "18787878878",
		Password: "sjojj",
		NickName: "bobby1",
	}
	assert.Equal(t, expUser, user)

	/*
		fake 测试
		grpc 服务， rocketmq， kafka，
	*/
}
