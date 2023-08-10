package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"strconv"
	"tiktok/internal/biz"
)

type userRepo struct {
	data *Database
	log  *log.Helper
}

func NewUserRepo(data *Database, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u userRepo) CreateUser(ctx context.Context, req *biz.RegisterLoginRequest) (*biz.RegisterLoginResponse, error) {

	return &biz.RegisterLoginResponse{
		StatusCode: strconv.Itoa(200),
		StatusMsg:  "OK",
		UserId:     1,
		Token:      "",
	}, nil
}

func (u userRepo) VerifyPassword(ctx context.Context, req *biz.RegisterLoginRequest) (*biz.RegisterLoginResponse, error) {
	panic("implement me")
}
