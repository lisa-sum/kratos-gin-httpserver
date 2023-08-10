package service

import (
	"context"
	"golang.org/x/exp/slog"
	"tiktok/api/user"
	"tiktok/internal/biz"
)

type UserService struct {
	user.UnimplementedUserServer

	uc *biz.UserUseCase
}

func NewUserService(uc *biz.UserUseCase) *UserService {
	return &UserService{uc: uc}
}

func (u *UserService) Register(ctx context.Context, req *user.UserRegisterLoginRequest) (*user.UserRegisterLoginResponse, error) {
	slog.Debug("UserService.Login")

	result, err := u.uc.CreateUser(ctx, &biz.RegisterLoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &user.UserRegisterLoginResponse{
		StatusCode: 200,
		StatusMsg:  "OK",
		UserId:     result.UserId,
		Token:      result.Token,
	}, nil
}

func (u *UserService) Login(ctx context.Context, req *user.UserRegisterLoginRequest) (*user.UserRegisterLoginResponse, error) {
	slog.Debug("UserService.Login")

	result, err := u.uc.VerifyPassword(ctx, &biz.RegisterLoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &user.UserRegisterLoginResponse{
		StatusCode: 200,
		StatusMsg:  "OK",
		UserId:     result.UserId,
		Token:      result.Token,
	}, nil
}
