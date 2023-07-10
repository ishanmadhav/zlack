package models

import "gorm.io/gorm"

//Message is a message sent by a user in a channel. Cam be a DM or a channel message
type Message struct {
	gorm.Model
	Content     string `json:"content"`
	From        uint   `json:"from"`
	To          uint   `json:"to"`
	IsDM        bool   `json:"is_dm"`
	ChannelID   uint   `json:"channel_id"`
	WorkspaceID uint   `json:"workspace_id"`
}
