package models

import (
	"strconv"
	"time"

	_ "gorm.io/gorm"
)

type List struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	UserID  string `json:"user_id" gorm:"size:191"`
	Title   string `json:"title" gorm:"not null"`
	State   string `json:"state"`
	Content string `json:"content"`
	Date    int    `json:"date"`
	User    User   `gorm:"foreignKey:UserID"`
}

type CardData struct {
	UserID string
	// UserName   string
	Date       string
	Done       Box
	InProgress Box
	ToDo       Box
}

type Box struct {
	UserID string
	Date   string
	List   []List
}

const (
	Done       string = "Done"
	InProgress string = "InProgress"
	ToDo       string = "ToDo"
)

func GetCards(uid string, date int) CardData {
	var cards []List
	var cd CardData
	DB.Where("user_id = ? AND date = ?", uid, date).Find(&cards)
	// db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
	dateString := strconv.Itoa(date)
	cd.UserID = uid
	cd.Date = dateString

	cd.Done.UserID = uid
	cd.Done.Date = dateString

	cd.InProgress.UserID = uid
	cd.InProgress.Date = dateString

	cd.ToDo.UserID = uid
	cd.ToDo.Date = dateString

	for _, card := range cards {
		switch card.State {
		case Done:
			cd.Done.List = append(cd.Done.List, card)

		case InProgress:
			cd.InProgress.List = append(cd.InProgress.List, card)

		case ToDo:
			cd.ToDo.List = append(cd.ToDo.List, card)
		}
	}

	return cd
}

func GetDate(t time.Time) int {

	date, _ := strconv.Atoi(t.Format("20060102"))
	return date
}

func DateToString(date int) string {
	return strconv.Itoa(date)
}
