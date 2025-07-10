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
      setPlayerId(message.data);
    };

    const updateCurrentTurnId = (message) => {
      console.log("game_state msg:", message.data);
      console.log("current turn:", message.data.currentTurn)
      const currentTurnId = message.data.turnOrder[message.data.currentTurn];
      setCurrentTurnId(currentTurnId);
    };

    addListener("new_id", updatePlayerId);
    addListener("game_state", updateCurrentTurnId);

    return () => {
      removeListener("new_id", updatePlayerId);
      removeListener("game_state", updateCurrentTurnId);
    };

  }, [addListener, removeListener]);

  return (
    <PlayerContext.Provider value={{ playerId, isMyTurn, currentTurnId }}>
      { children }
    </PlayerContext.Provider>
  );
};

export const usePlayer = () => useContext(PlayerContext);