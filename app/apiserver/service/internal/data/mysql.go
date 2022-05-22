package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tanmingqian/go-tam/app/apiserver/service/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func (r *userRepo) FindByName(ctx context.Context, name string) (*biz.User, error) {

	u := User{}
	result := r.data.GetDBIns().WithContext(ctx).First(&u, name)
	if result.Error != nil {
		return nil, result.Error
	}
	return &biz.User{
		ObjectMeta:  u.ObjectMeta,
		Status:      u.Status,
		Nickname:    u.Nickname,
		Password:    u.Password,
		Email:       u.Email,
		Phone:       u.Phone,
		IsAdmin:     u.IsAdmin,
		TotalPolicy: u.TotalPolicy,
		LoginedAt:   u.LoginedAt,
	}, result.Error
}

func (r *userRepo) Delete(ctx context.Context, name string) error {
	u := User{}
	result := r.data.db.WithContext(ctx).First(&u, name)
	if result.Error != nil {
		return result.Error
	}
	result = r.data.db.Delete(u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *userRepo) Save(ctx context.Context, user *biz.User) (*biz.User, error) {
	u := User{
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
		ObjectMeta:  u.ObjectMeta,
		Status:      u.Status,
		Nickname:    u.Nickname,
		Password:    u.Password,
		Email:       u.Email,
		Phone:       u.Phone,
		IsAdmin:     u.IsAdmin,
		TotalPolicy: u.TotalPolicy,
		LoginedAt:   u.LoginedAt,
	}, result.Error
}

func (r *userRepo) Update(ctx context.Context, user *biz.User) (*biz.User, error) {
	u := User{}
	result := r.data.GetDBIns().WithContext(ctx).First(&u, user.InstanceID)
	if result.Error != nil {
		return nil, result.Error
	}
	u.InstanceID = user.InstanceID
	result = r.data.GetDBIns().WithContext(ctx).Save(&u)
	if result.Error != nil {
		return nil, result.Error
	}
	return &biz.User{
		ObjectMeta:  u.ObjectMeta,
		Status:      u.Status,
		Nickname:    u.Nickname,
		Password:    u.Password,
		Email:       u.Email,
		Phone:       u.Phone,
		IsAdmin:     u.IsAdmin,
		TotalPolicy: u.TotalPolicy,
		LoginedAt:   u.LoginedAt,
	}, nil
}

func (r *userRepo) FindByID(ctx context.Context, user *biz.User) (*biz.User, error) {
	u := User{}
	result := r.data.GetDBIns().WithContext(ctx).First(&u, user.InstanceID)
	return &biz.User{
		ObjectMeta:  u.ObjectMeta,
		Status:      u.Status,
		Nickname:    u.Nickname,
		Password:    u.Password,
		Email:       u.Email,
		Phone:       u.Phone,
		IsAdmin:     u.IsAdmin,
		TotalPolicy: u.TotalPolicy,
		LoginedAt:   u.LoginedAt,
	}, result.Error
}

func (r *userRepo) ListByUser(ctx context.Context, name string) (*biz.UserList, error) {
	var us []User
	result := r.data.GetDBIns().WithContext(ctx).Find(&us, name)
	if result.Error != nil {
		return nil, result.Error
	}
	rv := make([]*biz.User, 0)
	for _, u := range us {
		rv = append(rv, &biz.User{
			ObjectMeta:  u.ObjectMeta,
			Status:      u.Status,
			Nickname:    u.Nickname,
			Password:    u.Password,
			Email:       u.Email,
			Phone:       u.Phone,
			IsAdmin:     u.IsAdmin,
			TotalPolicy: u.TotalPolicy,
			LoginedAt:   u.LoginedAt,
		})
	}
	return &biz.UserList{
		Items: rv,
		//TODO listmeta
	}, nil
}

func (r *userRepo) ListAll(ctx context.Context) (*biz.UserList, error) {
	var us []User
	result := r.data.GetDBIns().WithContext(ctx).Find(&us)
	if result.Error != nil {
		return nil, result.Error
	}
	rv := make([]*biz.User, 0)
	for _, u := range us {
		rv = append(rv, &biz.User{
			ObjectMeta:  u.ObjectMeta,
			Status:      u.Status,
			Nickname:    u.Nickname,
			Password:    u.Password,
			Email:       u.Email,
			Phone:       u.Phone,
			IsAdmin:     u.IsAdmin,
			TotalPolicy: u.TotalPolicy,
			LoginedAt:   u.LoginedAt,
		})
	}
	return &biz.UserList{
		Items: rv,
		//TODO listmeta
	}, nil

}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/user")),
	}
}
