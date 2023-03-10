package mock

import "context"

/*
gomock 测试
安装：
	go get -u github.com/golang/mock/gomock

	go install github.com/golang/mock/mockgen
*/

type User struct {
	Mobile   string
	Password string
	NickName string
}

type UserServer struct {
	Db UserData
}

func (us *UserServer) GetUserByMobile(ctx context.Context, mobile string) (User, error) {
	user, err := us.Db.GetUserByMobile(ctx, mobile)
	if err != nil {
		return User{}, err
	}
	if user.NickName == "bobby18" {
		user.NickName = "bobby17"
	}
	return user, nil
}

type UserData interface {
	GetUserByMobile(ctx context.Context, mobile string) (User, error)
}
