package models

import (
	"time"

	_ "gorm.io/gorm"
)

type Session struct {
	SessionID   string `gorm:"primary_key"`
	UserID      string `gorm:"size:191"`
	CurrentTime time.Time
	User        User `gorm:"foreignKey:UserID"`
}
