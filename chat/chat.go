package chat

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

//we need to track the users that are currently connected
//so we add to some data strcuture whenever a new user gets connected
//and whenever that user gets disconnected, we remove that user from the data structure
//we can use map
//sync.Map?
//we should maintain another map that keeps track of the channel for that user
//we maintain servers for each channel that should be up and running because the user is connected that belongs in that channel
//we check if a channel server is up and running, if not, we spin up a new server and add the user to it
//should it be a map of sets?
//or map of maps?
//we only want to maintain one conenction to one user for now.

type ChatServerInterface interface {
	Startup()
}

type ChatServer struct {
}

// Data structures
var UserSet sync.Map
var ServerMap map[uint][]UserData

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}

func handler(conn *websocket.Conn) {
	fmt.Println("Some client connected to the chat server")
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		var socketMessage SocketMessage
		err = json.Unmarshal(msg, &socketMessage)
		if err != nil {
			fmt.Println(err)
			return
		}

		switch socketMessage.Command {
		case "JOIN":
			Join(conn, socketMessage)
		case "SEND":
			Send(conn, socketMessage)
		case "LEAVE":
			Leave(conn, socketMessage)
			return
		default:
			Default(conn, socketMessage)
		}
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}

	ws.WriteMessage(1, []byte("Hi new client"))
	go handler(ws)
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func setupSets() {
	ServerMap = make(map[uint][]UserData)
}

func Startup() {
	setupRoutes()
	setupSets()
	fmt.Println("Websocket server starting up on port 5001")
	log.Fatal(http.ListenAndServe(":5001", nil))
}

func (s *ChatServer) Startup() {
	setupRoutes()
	setupSets()
	fmt.Println("Websocket server starting up on port 5001")
	log.Fatal(http.ListenAndServe(":5001", nil))
}

func NewChatServer() *ChatServer {
	chatServerObj := ChatServer{}
	return &chatServerObj
}
