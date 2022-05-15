package data

import (
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
