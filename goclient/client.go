package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/ishanmadhav/zlack/models"
)

func main() {
	dialer := *websocket.DefaultDialer
	conn, _, err := dialer.Dial("ws://127.0.0.1:5001/ws", nil)
	if err != nil {
		fmt.Printf("Could not establish websocket connection")
		return
	}

	go func() {
		fmt.Println("Reading messsages from the websocket server")
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("WebSocket read error:", err)
				return
			}
			fmt.Print(message)
			var msg models.Message
			err = json.Unmarshal(message, &msg)
			if err != nil {
				log.Println("Message read error")
			}
			fmt.Print(msg)
		}
	}()

	// Send a test message to the WebSocket server
	message := []byte("JOIN")
	err = conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Println("WebSocket write error:", err)
		return
	}
	fmt.Println("Message sent to server")

	// Sleep to keep the program running and receive messages
	time.Sleep(time.Second * 100)
}
