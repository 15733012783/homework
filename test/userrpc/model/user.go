package model

import (
	"errors"
	"gorm.io/gorm"
)

type User struct {
	ID       int
	Username string
	Password string
	Tel      string
}

func UserLogin(u User) (error, User) {
	var user User
	err := db.Model(&User{}).Where("username = ? ", u.Username).First(&user).Error
	return err, user
}

func UserRegister(us User) int {
	db.Model(&User{}).Create(&us)
	return us.ID
}

func UserCreate(Username string) error {
	var user User
	err := db.Model(&User{}).Where("username = ? ", Username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}
