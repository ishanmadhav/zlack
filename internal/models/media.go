package models

import "gorm.io/gorm"

//Media is a file sent by a user in a channel or over DM. Can be an image, video, audio or any other file
type Media struct {
	gorm.Model
}
