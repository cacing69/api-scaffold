package mod

import (
	_ "github.com/jinzhu/gorm"
)

type User struct {
	ID       uint `gorm:"primary_key"`
	Name     string
	Email    string
	Password string `json:"Password" gorm:"column:user_password"`
}

func (m User) TableName() string {
	return "user"
}
