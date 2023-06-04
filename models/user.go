package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `json:"username"`
	Messages  []Message
	ChannelID uint `json:"channel"`
}
