package mysql

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tanmingqian/go-tam/app/apiserver/service/internal/biz"
	"github.com/tanmingqian/go-tam/app/apiserver/service/internal/data"
)

type userRepo struct {
	data *data.Data
	log  *log.Helper
}

func (r *userRepo) Save(ctx context.Context, user *biz.User) (*biz.User, error) {
	u := data.User{
		ObjectMeta:  user.ObjectMeta,
		Status:      user.Status,
		Nickname:    user.Nickname,
		Password:    user.Password,
		Email:       user.Email,
		Phone:       user.Phone,
		IsAdmin:     user.IsAdmin,
		TotalPolicy: user.TotalPolicy,
		LoginedAt:   user.LoginedAt,
	}
	result := r.data.GetDBIns().WithContext(ctx).Create(&u)
	return &biz.User{
		ObjectMeta:  user.ObjectMeta,
		Status:      user.Status,
		Nickname:    user.Nickname,
		Password:    user.Password,
		Email:       user.Email,
		Phone:       user.Phone,
		IsAdmin:     user.IsAdmin,
		TotalPolicy: user.TotalPolicy,
		LoginedAt:   user.LoginedAt,
	}, result.Error
}

func (r *userRepo) Update(ctx context.Context, user *biz.User) (*biz.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *userRepo) FindByID(ctx context.Context, user *biz.User) (*biz.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *userRepo) ListByUser(ctx context.Context, name string) (*biz.UserList, error) {
	//TODO implement me
	panic("implement me")
}

func (r *userRepo) ListAll(ctx context.Context) (*biz.UserList, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserRepo(data *data.Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/user")),
	}
}
