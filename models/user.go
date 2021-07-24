package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) CheckUser() error {
	err := db.Select("id").Where(&u).First(&u).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return  err
	}

	if u.ID > 0 {
		return nil
	}

	return nil
}

func (u *User) CreateUser() error {
	err := db.Create(&u).Error
	if err != nil {
		return err
	}

	return nil
}
