package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Tag      string `json:"tag"`
}
