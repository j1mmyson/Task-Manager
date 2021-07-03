package models

import (
	_ "gorm.io/gorm"
)

type List struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	User    string `json:"user"`
	Title   string `json:"title"`
	State   string `json:"state"`
	Content string `json:"content"`
}
