package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type RegisterLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterLoginResponse struct {
	StatusCode string `json:"username"`
	StatusMsg  string `json:"password"`
	UserId     int64  `json:"user_id"`
	Token      string `json:"token"`
}

type User struct {
	Id              int64  // 用户id
	Name            string // 用户名称
	FollowCount     int64  // 关注总数
	FollowerCount   int64  // 粉丝总数
	IsFollow        bool   // true-已关注，false-未关注
	Avatar          string // 用户头像
	BackgroundImage string // 用户个人页顶部大图
	Signature       string // 个人简介
	TotalFavorited  int64  // 获赞数量
	WorkCount       int64  // 作品数量
	FavoriteCount   int64  // 点赞数量
}

type UserRepo interface {
	CreateUser(ctx context.Context, req *RegisterLoginRequest) (*RegisterLoginResponse, error)
	VerifyPassword(ctx context.Context, req *RegisterLoginRequest) (*RegisterLoginResponse, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUseCase) CreateUser(ctx context.Context, req *RegisterLoginRequest) (*RegisterLoginResponse, error) {
	uc.log.WithContext(ctx).Infof("req%v", req)

	return uc.repo.CreateUser(ctx, req)
}

func (uc *UserUseCase) VerifyPassword(ctx context.Context, req *RegisterLoginRequest) (*RegisterLoginResponse, error) {
	uc.log.WithContext(ctx).Infof("req", req)

	return uc.repo.VerifyPassword(ctx, req)
}
