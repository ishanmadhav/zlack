package models

import "gorm.io/gorm"

//Team is a group of people inside a workspace
type Team struct {
	gorm.Model
	Name    string `json:"name"`
	Members []User `gorm:"many2many:team_members;"`
}
