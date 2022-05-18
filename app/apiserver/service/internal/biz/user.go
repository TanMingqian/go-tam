package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tanmingqian/go-tam/pkg/metadata"
	"time"
)

// User is a user model. It is also used as gorm model.
type User struct {
	// Standard object's metadata
	metadata.ObjectMeta `json:"metadata,omitempty"`

	Status int `json:"ID" `

	// Required: true
	Nickname string `json:"nickname" gorm:"column:nickname" validate:"required,min=1,max=30"`

	// Required: true
	Password string `json:"password,omitempty" gorm:"column:password" validate:"required"`

	// Required: true
	Email string `json:"email" gorm:"column:email" validate:"required,email,min=1,max=100"`

	Phone string `json:"phone" gorm:"column:phone" validate:"omitempty"`

	IsAdmin int `json:"isAdmin,omitempty" gorm:"column:isAdmin" validate:"omitempty"`

	TotalPolicy int64 `json:"totalPolicy" gorm:"-" validate:"omitempty"`

	LoginedAt time.Time `json:"loginedAt,omitempty" gorm:"column:loginedAt"`
}

// UserList is the whole list of all users which have been stored in storage.
type UserList struct {
	// May add TypeMeta in the future.

	// Standard list metadata.
	// +optional
	metadata.ListMeta `json:",inline"`

	Items []*User `json:"items"`
}

type UserRepo interface {
	Save(ctx context.Context, user *User) (*User, error)
	Update(ctx context.Context, user *User) (*User, error)
	FindByID(ctx context.Context, user *User) (*User, error)
	ListByUser(ctx context.Context, name string) (*UserList, error)
	ListAll(ctx context.Context) (*UserList, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/order")),
	}
}

func (uc *UserUseCase) Create(ctx context.Context, u *User) (*User, error) {
	return uc.repo.Save(ctx, u)
}

func (uc *UserUseCase) Get(ctx context.Context, u *User) (*User, error) {
	return uc.repo.FindByID(ctx, u)
}

func (uc *UserUseCase) Update(ctx context.Context, u *User) (*User, error) {
	return uc.repo.Update(ctx, u)
}

func (uc *UserUseCase) List(ctx context.Context) (*UserList, error) {
	return uc.repo.ListAll(ctx)
}
