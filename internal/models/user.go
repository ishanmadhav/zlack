package models

import "gorm.io/gorm"

//User is a person who uses the application. Can be an admin, moderator or simple end user
type User struct {
	gorm.Model
	Username   string      `json:"username"`
	Email      string      `json:"email"`
	Firstname  string      `json:"firstname"`
	Lastname   string      `json:"lastname"`
	Password   string      `json:"password"`
	Roles      string      `json:"roles"`
	Workspaces []Workspace `gorm:"many2many:workspace_members;"`
	Teams      []Team      `gorm:"many2many:team_members;"`
}
