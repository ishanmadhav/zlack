package models

import "github.com/gorilla/websocket"

type User struct {
	Username string
	Id       int
	Conn     *websocket.Conn
}
