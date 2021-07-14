package models

import (
	"strconv"
	"time"

	_ "gorm.io/gorm"
)

type List struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	UserID  string `json:"user_id" gorm:"size:191"`
	Title   string `json:"title"`
	State   string `json:"state"`
	Content string `json:"content"`
	Date    int    `json:"date"`
	User    User   `gorm:"foreignKey:UserID"`
}

type CardData struct {
	UserID     string
	Date       string
	Done       []List
	InProgress []List
	ToDo       []List
}

const (
	Done       string = "Done"
	InProgress string = "InProgress"
	ToDo       string = "ToDo"
)

func GetCards(uid string, date int) []List {
	var cards []List
	DB.Where("user_id = ? AND date = ?", uid, date).Find(&cards)
	// db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
	return cards
}

func MakeStructFromCards(cards []List) CardData {
	var cd CardData

	// cd.UserID = cards[0].UserID
	// cd.Date = DateToString(cards[0].Date)

	cd.UserID = "hello;"
	cd.Date = "20210714"

	for _, card := range cards {
		switch card.State {
		case Done:
			cd.Done = append(cd.Done, card)

		case InProgress:
			cd.InProgress = append(cd.InProgress, card)

		case ToDo:
			cd.ToDo = append(cd.ToDo, card)
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
