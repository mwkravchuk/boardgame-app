import { useEffect } from "react";

const WebSocketComponent = () => {

  // Just connect to the websocket once (on render)
  // note: if u see it twice, it's because of react's strictmode. wont happen in production.
  useEffect(() => {
    const ws = new WebSocket("ws://localhost:8080/ws");
    console.log("Attemping websocket connection");
    ws.onopen = () => {
      console.log("successfully connected");
      ws.send("hi from the client!");
    }

    ws.onclose = (event) => {
      console.log("socket closed connection: ", event);
    }

    // Remember that our server just reads and then writes the message back. So when it writes
    // it back, this is our client receiving it, and then just logging it. Our server echoes messages
    ws.onmessage = (msg) => {
      console.log(msg);
    }

    ws.onerror = (error) => {
      console.log("socket error: ", error);
    }
  }, []);

  return null; // Doesn't render anything, just handles data
};

export default WebSocketComponent;