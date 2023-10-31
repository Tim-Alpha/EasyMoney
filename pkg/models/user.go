package models

import (
	"github.com/jinzhu/gorm"
	"github.com/nk-code-lab/EasyMoney/pkg/config"
)

var db *gorm.DB

type User struct {
    gorm.Model
    Name       string `gorm:"column:name" json:"name"`
    Email      string `gorm:"column:email" json:"email"`
    Mobile     int64  `gorm:"column:mobile" json:"mobile"`
    Money      float64 `gorm:"column:money" json:"money"`
    Is_Verified bool `gorm:"column:is_verified" json:"is_verified"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) CreateUser() *User{
	db.NewRecord((u))
	db.Create(&u)
	return u
}

func GetAllUser() ([]User, error){
	var users []User
    if err := db.Find(&users).Error; err != nil {
        return nil, err
    }
	return users, nil
}

func GetUserByID(Id int64) (*User, *gorm.DB){
	var getUser User
	db := db.Where("ID=?", Id).Find(&getUser)
	return &getUser, db
}

func DeleteUser(Id int64) string{
	var user User
	db.Where("ID=?", Id).Delete(user)
	return "Deleted Successfully"
}