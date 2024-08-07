package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title, Slug, Description, Content, Picture_url string
	CategoryID                                     int
}

func (post Post) Migrate() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	db.AutoMigrate(&post)
}

func (post Post) Add() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	db.Create(&post)
}

// stringde int de gelebileceği için interface verdik
func (post Post) Get(where ...interface{}) Post {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return post
	}

	db.Where(where[0], where[1:]...).First(&post) //where çekmek için
	return post
}

func (post Post) GetAll(where ...interface{}) []Post {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var posts []Post
	db.Find(&posts, where...)
	return posts
}

func (post Post) Update(colum string, value interface{}) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	db.Model(&post).Update(colum, value)
}

func (post Post) Updates(data Post) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	db.Model(&post).Updates(data)
}

func (post Post) Delete() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	db.Delete(&post, post.ID)
}
