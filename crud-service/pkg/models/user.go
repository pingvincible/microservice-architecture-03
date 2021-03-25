package models

import (
	"github.com/pingvincible/crud-service/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	//Id          string `json:"id"`
	UserName	string `gorm:""json:"username"`
	FirstName	string `json:"firstName"`
	LastName	string `json:"lastName"`
	Email		string `json:"email"`
	Phone		string `json:"phone"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (b *User) CreateUser() *User {
	db.Create(&b)
	return b
}

func GetUserById(Id int64) (*User , *gorm.DB){
	var getUser User
	db:=db.Where("ID = ?", Id).Find(&getUser)
	return &getUser, db
}

func DeleteUser(ID int64) User {
	var user User
	db.Delete(&user, ID)
	return user
}