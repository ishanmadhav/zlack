package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Text      string `json:"text"`
	UserID    uint   `json:"userId"`
	ChannelID uint   `json:"channelId"`
}
