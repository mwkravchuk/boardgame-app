import { createContext, useContext, useState, useEffect } from "react";
import { useWebSocket } from "./WebSocketProvider";

const PlayerContext = createContext();

export const PlayerProvider = ({ children }) => {

  const { addListener, removeListener } = useWebSocket();
  const [playerId, setPlayerId] = useState("");
  const [currentTurnId, setCurrentTurnId] = useState("");

  const isMyTurn = playerId === currentTurnId;

  useEffect(() => {
    
    const updatePlayerId = (message) => {
      console.log("newid msg:", message);
      setPlayerId(message.data);
    };

    const updateCurrentTurnId = (message) => {
      console.log("newturn msg:", message);
      setCurrentTurnId(message.data);
    };

    addListener("new_id", updatePlayerId);
    addListener("new_turn", updateCurrentTurnId);

    return () => {
      removeListener("new_id", updatePlayerId);
      removeListener("new_turn", updateCurrentTurnId);
    };

  }, [addListener, removeListener]);

  return (
    <PlayerContext.Provider value={{ playerId, isMyTurn, currentTurnId }}>
      { children }
    </PlayerContext.Provider>
  );
};

export const usePlayer = () => useContext(PlayerContext);