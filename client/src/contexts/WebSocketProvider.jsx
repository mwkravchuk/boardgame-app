import { createContext, useContext, useRef } from "react";

const WebSocketContext = createContext();

const WebSocketProvider = ({ children }) => {

  




  <WebSocketContext.Provider value={}>
    { children }
  </WebSocketContext.Provider>
};

export const useWebSocket = () => useContext(WebSocketContext);