import { useState, useEffect } from "react";
import ChatBox from "../ChatBox";
import Board from "../Board";

const WebSocketComponent = ({ setChatHistory }) => {

  const [ws, setWs] = useState(null);

  // Just connect to the websocket once (on render)
  // note: if u see it twice, it's because of react's strictmode. wont happen in production.
  useEffect(() => {
    const ws = new WebSocket("ws://localhost:8080/ws");
    console.log("Attemping websocket connection");
    ws.onopen = () => {
      console.log("successfully connected");
      setWs(ws);
    }

    ws.onclose = (event) => {
      console.log("socket closed connection: ", event);
    }

    // Remember that our server just reads and then writes the message back. So when it writes
    // it back, this is our client receiving it, and then just logging it. Our server echoes messages
    ws.onmessage = (msg) => {
      console.log("Received message: ", msg);
      const parsedData = JSON.parse(msg.data); // Parse JSON message
      console.log("Received message:", parsedData);

      if (msg.data) {
        if (parsedData.type === "chat") {
          setChatHistory(prevHistory => [
            ...prevHistory,
            `${parsedData.sender}: ${parsedData.data}`
          ]);
        } else if (parsedData.type === "roll_dice") {
          setChatHistory(prevHistory => [
            ...prevHistory,
            `${parsedData.sender}rolled: ${parsedData.data.dice1} ${parsedData.data.dice2}`
          ]);
        }
      }

    }

    ws.onerror = (error) => {
      console.log("socket error: ", error);
    }

    return () => { // clean up when unmounts
      ws.close();
    }

  }, [setChatHistory]);

  return (
    <div>
      {ws && <ChatBox ws={ws}/>}
      {ws && <Board ws={ws}/>}
    </div>
  )
};

export default WebSocketComponent;