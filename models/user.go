package models

import (
	_ "gorm.io/gorm"
)

type User struct {
	ID       string `json:"id" gorm:"primary_key"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
