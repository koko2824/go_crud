package models

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	Title string
	Text string
	Status string
}