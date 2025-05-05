import { createContext, useContext, useEffect, useRef } from "react";

const WebSocketContext = createContext();

export const WebSocketProvider = ({ children }) => {

  const webSocketRef = useRef(null);

  // Map: Key = msgType, Val = [functions to call]
  const listenersRef = useRef({});

  useEffect(() => {
    const ws = new WebSocket("ws://localhost:8080/ws");
    ws.onopen = () => {
      console.log("WebSocket opened");
    }

    ws.onmessage = (e) => {
      const msg = e.data;
      const parsedMsg = JSON.parse(msg);
      const { type, sender, data } = parsedMsg;
      //console.log("type: ", type);
      //console.log("sender: ", sender);
      //console.log("data: ", data);
    
      // For each function that has waited for this message
      // type to happen, call them.

      const message = { sender, data };
      
      const listeners = listenersRef.current[type] || [];
      listeners.forEach((cb) => cb(message));

      // Say listenersRef = {game_state: [x, y, z]}
      // if the msg type is game_state,
      // then we call x, y, z with data as argument
    }

    ws.onclose = () => {
      console.log("WebSocket closed");
    }

    webSocketRef.current = ws;

    return () => {
      ws.close();
    }
  }, []);

  const sendMessage = (type, data) => {
    if (webSocketRef.current && webSocketRef.current.readyState === WebSocket.OPEN ) {
      webSocketRef.current.send(JSON.stringify({ type, data }));
    } else {
      console.warn("WebSocket is not open. Cannot send.")
    };
  };

  const addListener = (type, cb) => {
    if (!listenersRef.current[type]) {
      listenersRef.current[type] = [];
    }
    listenersRef.current[type].push(cb);
  };

  const removeListener = (type, cb) => {
    if (!listenersRef.current[type]) {
      return;
    }
    // Keeps all functions that != cb
    listenersRef.current[type] = listenersRef.current[type].filter(
      (func) => func !== cb
    );
  };

  return (
    <WebSocketContext.Provider value={{sendMessage, addListener, removeListener}}>
      { children }
    </WebSocketContext.Provider>
  );
};

export const useWebSocket = () => useContext(WebSocketContext);