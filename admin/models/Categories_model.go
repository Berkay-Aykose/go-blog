package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Title, Slug string
}

func (category Category) Migrate() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	db.AutoMigrate(&category)
}

func (category Category) Add() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	db.Create(&category)
}

// stringde int de gelebileceği için interface verdik
func (category Category) Get(where ...interface{}) Category {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return category
	}

	db.First(&category, where)
	return category
}

func (category Category) GetAll(where ...interface{}) []Category {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var categories []Category
	db.Find(&categories, where...)
	return categories
}

func (category Category) Update(colum string, value interface{}) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	db.Model(&category).Update(colum, value)
}

func (category Category) Updates(data Category) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	db.Model(&category).Updates(data)
}

func (category Category) Delete() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	db.Delete(&category, category.ID)
}
