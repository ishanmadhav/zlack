package models

import "gorm.io/gorm"

type Channel struct {
	gorm.Model
	Name  string `json:"name"`
	Users []User
}
