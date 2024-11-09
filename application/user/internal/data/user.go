package data

import (
	"context"

	"user/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func (r *userRepo) CreateUser(ctx context.Context, req *biz.UserRequest) (*biz.UserReply, error) {
	user, err := r.data.CreateUser(ctx, CreateUserParams{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &biz.UserReply{
		UserId: int32(user.ID),
	}, nil
}

func (r *userRepo) LoginUser(ctx context.Context, request *biz.LoginRequest) (*biz.UserReply, error) {
	user, err := r.data.LoginUser(ctx, request.Email)
	if err != nil {
		return nil, err
	}
	return &biz.UserReply{UserId: int32(user.ID)}, nil
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
