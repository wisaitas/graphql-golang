package entity

import (
	"strconv"
	"time"

	"github.com/wisaitas/graphql-golang/internal/appv2/model"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	Username  string         `json:"username" gorm:"unique;not null"`
	Email     string         `json:"email" gorm:"unique;not null"`
	Password  string         `json:"-" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (User) TableName() string {
	return "tbl_users"
}

func (e *User) EntityToModel() *model.User {
	return &model.User{
		ID:       strconv.FormatUint(uint64(e.ID), 10),
		Username: e.Username,
		Email:    e.Email,
	}
}
