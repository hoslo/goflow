package models

import (
	"errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	IsAdmin  int    `json:"is_admin"`
}

func (u *User) CheckUser() error {
	err := db.Select("id").Where(&u).First(&u).Error

	if err != nil {
		return  err
	}

	if u.ID > 0 {
		return nil
	}

	return nil
}

func (u *User) CreateUser() error {
	var existUser User
	db.Where("username = ?", u.Username).First(&existUser)
	if existUser.ID > 0 {
		return errors.New("exist user")
	}

	err := db.Create(&u).Error
	if err != nil {
		return err
	}

	return nil
}

func QueryUsers() ([]User, error) {
	var existUsers []User
	err := db.Find(&existUsers).Error

	if err != nil {
		return nil, err
	}

	return existUsers, nil
}
