package models

import (
	"time"

	_ "gorm.io/gorm"
)

type User struct {
	ID        string    `json:"user_id" gorm:"primary_key"`
	Password  string    `json:"password"`
	Name      string    `json:"user_name"`
	CreatedAt time.Time `json:"created_at"`
}

func GetUserFromUserId(uid string) User {
	var u User

	DB.First(&u, "id = ?", uid)
	return u
}
