package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/ishanmadhav/zlack/models"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var roomList []models.Room
var userList []models.User
var roomMap map[int][]models.User
var curr int

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}

func reader(conn *websocket.Conn) {
	curr++
	user := models.User{Username: "random_name", Id: curr, Conn: conn}
	userList = append(userList, user)
	fmt.Print(userList)
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Print(p)
		var message models.Message
		err = json.Unmarshal(p, &message)

		if err != nil {
			fmt.Println("There was an error with the message")
			fmt.Println(err)
			return
		}

		if message.Commannd == "JOIN" {
			roomMap[message.RoomID] = append(roomMap[message.RoomID], user)
			fmt.Println("Joining room")
		} else if message.Commannd == "SEND" {
			arr := roomMap[message.RoomID]
			for i := 0; i < len(arr); i++ {
				currUser := arr[i]
				err = currUser.Conn.WriteMessage(1, []byte(message.Message))
				if err != nil {
					fmt.Println("There was an error sending the message")
					log.Println(err)
				}
			}
		}

	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Printf("New client connected")
	ws.WriteMessage(1, []byte("Hi, client!"))
	go reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	fmt.Printf("Hello World")
	curr = 0
	roomMap = make(map[int][]models.User)
	setupRoutes()
	log.Fatal(http.ListenAndServe(":5000", nil))
}
