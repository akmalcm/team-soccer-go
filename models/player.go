package models

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	PlayerID int    `json:"ID" gorm:"unique"`
	Forename string `json:"Forename"`
	Surname  string `json:"Surname"`
	ImageURL string `json:"ImageUrl"`
}
