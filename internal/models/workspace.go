package models

import "gorm.io/gorm"

//Workspace is a group of people, like an organization or big team
type Workspace struct {
	gorm.Model
	Name     string    `json:"name"`
	TeamID   uint      `json:"team_id"`
	Members  []User    `gorm:"many2many:workspace_members;"`
	Channels []Channel `gorm:"foreignKey:WorkspaceID"`
	Messages []Message `gorm:"foreignKey:WorkspaceID"`
}
