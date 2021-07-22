package models

import (
	"time"

	_ "gorm.io/gorm"
)

type Session struct {
	SessionID string `gorm:"primary_key"`
	UserID    string `gorm:"size:191;unique"`
	UpdatedAt time.Time
	User      User `gorm:"foreignKey:UserID"`
}

func CreateSession(sid string, uid string) {
	var s Session

	s.SessionID, s.UserID, s.UpdatedAt = sid, uid, time.Now()
	DB.Create(&s)
}

func GetUserIdFromSession(sid string) (string, error) {
	var s Session

	if err := DB.First(&s, "session_id = ?", sid).Error; err != nil {
		return "", err
	}

	return s.UserID, nil
}

func GetSessionFromUserId(uid string) Session {
	var s Session

	if err := DB.First(&s, "user_id = ?", uid).Error; err != nil {
		return Session{}
	}

	return s
}

func UpdateCurrentTime(s Session) {
	DB.Save(&s)
}

func DeleteSession(sid string) {
	var s Session

	s.SessionID = sid
	DB.Delete(&s)
}

func CleanSessions() {
	var sessions []Session

	DB.Find(&sessions)
	for _, s := range sessions {
		if time.Now().Sub(s.UpdatedAt) > (time.Minute * 30) {
			DeleteSession(s.SessionID)
		}
	}

	DbSessionCleaned = time.Now()
}
