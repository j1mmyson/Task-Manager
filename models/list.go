package models

import (
	_ "gorm.io/gorm"
)

type List struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	UserID  string `json:"user_id" gorm:"size:191"`
	Title   string `json:"title"`
	State   string `json:"state"`
	Content string `json:"content"`
	User    User   `gorm:"foreignKey:UserID"`
}
