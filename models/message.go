package models

type Message struct {
	Message  string `json:"message"`
	Username string `json:"user"`
	RoomID   int    `json:"roomID"`
	Commannd string `json:"command"`
}
