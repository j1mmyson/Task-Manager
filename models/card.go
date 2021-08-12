package models

import (
	"fmt"
	"strconv"
	"time"

	_ "gorm.io/gorm"
)

type List struct {
	ID      uint      `json:"id" gorm:"primary_key"`
	UserID  string    `json:"user_id" gorm:"size:191"`
	Title   string    `json:"title" gorm:"not null"`
	State   string    `json:"state"`
	Content string    `json:"content"`
	Date    time.Time `json:"date"`
	User    User      `gorm:"foreignKey:UserID"`
}

type cdList struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	UserID  string `json:"user_id" gorm:"size:191"`
	Title   string `json:"title" gorm:"not null"`
	State   string `json:"state"`
	Content string `json:"content"`
	Date    string `json:"date"`
}

type CardData struct {
	UserID     string
	Done       Box
	InProgress Box
	ToDo       Box
}

type Box struct {
	UserID string
	List   []cdList
}

const (
	Done       string = "Done"
	InProgress string = "InProgress"
	ToDo       string = "ToDo"
)

// func GetCards(uid string, date int) CardData {
func GetCards(uid string) CardData {
	var cards []List
	var cd CardData
	DB.Where("user_id = ?", uid).Find(&cards)

	cd.UserID = uid

	cd.Done.UserID = uid

	cd.InProgress.UserID = uid

	cd.ToDo.UserID = uid

	for _, card := range cards {
		switch card.State {
		case Done:
			cd.Done.List = append(cd.Done.List, convertCard(card))

		case InProgress:
			cd.InProgress.List = append(cd.InProgress.List, convertCard(card))

		case ToDo:
			cd.ToDo.List = append(cd.ToDo.List, convertCard(card))
		}
	}

	return cd
}

func convertCard(card List) cdList {
	cd := cdList{}
	cd.ID = card.ID
	cd.Content = card.Content
	cd.State = card.State
	cd.Title = card.Title
	cd.UserID = card.UserID
	cd.Date = GetDate(card.Date)

	return cd
}

func GetDate(t time.Time) string {
	return FormatDate(t.Format("20060102"))
}

func DateToString(date int) string {
	return strconv.Itoa(date)
}

func FormatDate(date string) string {
	sDate := date
	year := sDate[:4]
	month := getMonth(sDate[4:6])
	day := sDate[6:]

	return fmt.Sprintf("%s %s, %s", month, day, year)
}

func getMonth(m string) string {
	switch m {
	case "01":
		return "Jan"
	case "02":
		return "Feb"
	case "03":
		return "March"
	case "04":
		return "Apr"
	case "05":
		return "May"
	case "06":
		return "Jun"
	case "07":
		return "Jul"
	case "08":
		return "Aug"
	case "09":
		return "Sep"
	case "10":
		return "Oct"
	case "11":
		return "Nov"
	}
	return "Dec"
}
