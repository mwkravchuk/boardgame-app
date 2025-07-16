import { createContext, useContext, useState, useEffect } from "react";
import { useWebSocket } from "./WebSocketProvider";

const PlayerContext = createContext();

export const PlayerProvider = ({ children }) => {

  const { addListener, removeListener } = useWebSocket();
  const [playerId, setPlayerId] = useState("");

  useEffect(() => {
    const updatePlayerId = (message) => {
      console.log("player id: ", message.data);
      setPlayerId(message.data);
    };

    addListener("new_id", updatePlayerId);

    return () => {
      removeListener("new_id", updatePlayerId);
    };

  }, [addListener, removeListener]);

  return (
    <PlayerContext.Provider value={{ playerId }}>
      { children }
    </PlayerContext.Provider>
  );
};

export const usePlayer = () => useContext(PlayerContext);