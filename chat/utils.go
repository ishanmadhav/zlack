package chat

import "github.com/gorilla/websocket"

type SocketMessage struct {
	Command   string `json:"command"`
	Message   string `json:"message"`
	Username  string `json:"username"`
	ChannelID uint   `json:"channel"`
}

type UserData struct {
	Username  string `json:"username"`
	ChannelID uint   `json:"channel"`
	Conn      *websocket.Conn
}
