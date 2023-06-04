import { SetStateAction, useEffect, useState} from "react"

let ws=new WebSocket("ws://127.0.0.1:5001/ws")

export default function Chat() {

  const [socket, setSocket]=useState<WebSocket|null>(null)
  const [text, setText]=useState("")
  const [room, setRoom]=useState<string>("")
  const [username, setUsername]=useState("")

  useEffect(()=>{

    ws.onopen=(e)=>{
      console.log("The connection was established")
      //ws.send("Hello from the client")
      setSocket(ws)
    }
  }, [])

  useEffect(()=>{
    if (!socket)
    {
      return
    }

    socket.onmessage=(e)=>{
      console.log(e.data)
    }
  }, [socket])

  const handleChange=(e:any)=>{
    setText(e.target.value)
  }

  const handleUsernameChnage=(e:any)=>{
    setUsername(e.target.value)
  }

  const handleSendMessage=()=>{
    const num=parseInt(room)
    socket?.send(JSON.stringify({
      command: "SEND",
      message: text,
      username: username,
      channel: num
    }))
  }

  const handleRoomChange=(e:any)=>{
    setRoom(e.target.value)
  }

  const handleJoinRoom=()=>{
    const num=parseInt(room)
    socket?.send(JSON.stringify({
      command: "JOIN",
      message: text,
      username:  username,
      channel: num
    }))
  }
  return (
    <div>
      <button onClick={handleSendMessage}>Send Message</button>

      <p>Join</p>
      Username
      <input onChange={handleUsernameChnage} />
      Channel ID
      <input onChange={handleRoomChange}/>
      Message
      <input onChange={handleChange} />

      <button onClick={handleJoinRoom}>Join</button>
    </div>
  )
}