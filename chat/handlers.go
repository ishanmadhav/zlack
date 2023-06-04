package chat

import (
	"fmt"

	"github.com/gorilla/websocket"
)

func Join(conn *websocket.Conn, socketMessage SocketMessage) {
	fmt.Println("User join handler")
	_, ok := UserSet.Load(socketMessage.Username)
	if ok {
		err := conn.Close()
		if err != nil {
			fmt.Println("Error while closing websocket connection")
			fmt.Println(err)
			return
		}
		fmt.Println("Websocket connection closing")
	}

	UserSet.Store(socketMessage.Username, conn)
	fmt.Println("user added to set")
	channelID := socketMessage.ChannelID
	fmt.Println("The socket message is as follows")
	fmt.Print(socketMessage)
	serverResult := ServerMap[channelID]
	tempUserData := UserData{Username: socketMessage.Username, ChannelID: socketMessage.ChannelID, Conn: conn}
	if serverResult != nil {
		ServerMap[channelID] = append(ServerMap[channelID], tempUserData)
		fmt.Println("User added to server")
	} else {
		slice := []UserData{}
		slice = append(slice, tempUserData)
		ServerMap[channelID] = slice
		fmt.Println("User added to server")
		fmt.Print(ServerMap[channelID])
	}

}

func Send(conn *websocket.Conn, socketMessage SocketMessage) {
	fmt.Println("Message Send Handler")
	channelID := socketMessage.ChannelID
	serverResult := ServerMap[channelID]
	messageInBytes := socketMessage.Message
	fmt.Print(socketMessage)
	fmt.Println(messageInBytes)
	if serverResult != nil {
		for i := 0; i < len(serverResult); i++ {
			if serverResult[i].Username == socketMessage.Username {
				continue
			}
			serverResult[i].Conn.WriteMessage(1, []byte(messageInBytes))
		}
	} else {
		//Case to be handled later
		fmt.Println("User hasn't joined yet.")
	}
}

func Leave(conn *websocket.Conn, socketMessage SocketMessage) {
	fmt.Println("Closing connection")
	conn.Close()

}

func Default(conn *websocket.Conn, socketMessage SocketMessage) {
	fmt.Println("Do nothing")
	return
}
