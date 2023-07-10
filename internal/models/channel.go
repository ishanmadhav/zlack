package models

import "gorm.io/gorm"

//Channel is a group of people inside a workspace focussed on some particular topic
type Channel struct {
	gorm.Model
	Name        string `json:"name"`
	WorkspaceID uint   `json:"workspace_id"`
	Messages    []Message
}
