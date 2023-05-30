import { SetStateAction, useEffect, useState} from "react"



export default function Chat() {

  const [socket, setSocket]=useState<WebSocket|null>(null)
  const [text, setText]=useState("")
  const [room, setRoom]=useState<string>("")

  useEffect(()=>{
    let ws=new WebSocket("ws://127.0.0.1:5000/ws")
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

  const handleClick=()=>{
    const num=parseInt(room)
    socket?.send(JSON.stringify({
      message: text,
      user: "Ishan",
      roomID: num,
      command: "SEND"
    }))
  }

  const handleRoomChange=(e:any)=>{
    setRoom(e.target.value)
  }

  const handleJoinRoom=()=>{
    const num=parseInt(room)
    socket?.send(JSON.stringify({
      message: text,
      user: "John",
      command: "JOIN",
      roomID: num
    }))
  }
  return (
    <div>
      <input onChange={handleChange} />
      <button onClick={handleClick}>Send</button>

      <p>Join</p>
      <input onChange={handleRoomChange}/>
      <button onClick={handleJoinRoom}>Join</button>
    </div>
  )
}