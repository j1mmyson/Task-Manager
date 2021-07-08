package models

import (
	"log"
	"time"

	_ "gorm.io/gorm"
)

type Session struct {
	SessionID string `gorm:"primary_key"`
	UserID    string `gorm:"size:191"`
	UpdatedAt time.Time
	User      User `gorm:"foreignKey:UserID"`
}

func CreateSession(sid string, uid string) {
	log.Println("CreateSession() !")
	var s Session
	s.SessionID = sid
	s.UserID = uid
	s.UpdatedAt = time.Now()

	DB.Create(&s)

}

func GetUserIdFromSession(sid string) (string, error) {
	var s Session

	if err := DB.First(&s, "session_id = ?", sid).Error; err != nil {
		return "", err
	}

	return s.UserID, nil
}

func GetSessionFromUserId(uid string) *Session {
	var s Session

	if err := DB.First(&s, "user_id = ?", uid).Error; err != nil {
		return nil
	}

	return &s
}

func UpdateCurrentTime(s Session) {
	DB.Save(&s)
}

func DeleteSession(sid string) {
	var s Session
	s.SessionID = sid
	DB.Delete(&s)
}
