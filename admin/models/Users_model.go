package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username, Password string
}

func (user User) Migrate() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	db.AutoMigrate(&user)
}

func (user User) Add() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	db.Create(&user)
}

// stringde int de gelebileceği için interface verdik
func (user User) Get(where ...interface{}) User {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return user
	}

	db.Where(where[0], where[1:]...).First(&user)
	return user
}

func (user User) GetAll(where ...interface{}) []User {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var users []User
	db.Find(&users, where...)
	return users
}

func (user User) Update(colum string, value interface{}) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	db.Model(&user).Update(colum, value)
}

func (user User) Updates(data User) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	db.Model(&user).Updates(data)
}

func (user User) Delete() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	db.Delete(&user, user.ID)
}
