package models

import (
	"fmt"
	"goflow/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func initAdmin()  {
	user := User{}
	admin := User{Username: "admin", Password: "123456", IsAdmin: 1}
	db.Where("username = ?", "admin").First(&user)
	if user.ID == 0 {
		fmt.Println(user)
		db.Create(&admin)
	}
}

func SetUp()  {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
						setting.DatabaseSetting.User,
						setting.DatabaseSetting.Password,
						setting.DatabaseSetting.Host,
						setting.DatabaseSetting.Name)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}


	db.AutoMigrate(&User{})
	db.AutoMigrate(&DAG{})
	db.AutoMigrate(&HttpOperator{})
	db.AutoMigrate(&HttpOperatorInstance{})
	db.AutoMigrate(&Task{})
	db.AutoMigrate(&TaskToTask{})
	db.AutoMigrate(&TaskInstance{})

	initAdmin()
}

