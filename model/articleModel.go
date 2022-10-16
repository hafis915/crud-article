package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Status   string `json:"status"`
}
