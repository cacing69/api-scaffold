package models

import (
	_ "github.com/jinzhu/gorm"
)

type UserModel struct {
	ID       uint `gorm:"primary_key"`
	Name     string
	Email    string
	Password string `json:"Password" gorm:"column:user_password"`
}

func (m UserModel) TableName() string {
	return "user"
}
