import { createContext, useContext, useState, useEffect } from "react";
import { useWebSocket } from "./WebSocketProvider";

const GameContext = createContext();

export const GameProvider = ({ children }) => {
  const { addListener, removeListener } = useWebSocket();
  const [gameState, setGameState] = useState(null);

  useEffect(() => {
    const handleGameStateUpdate = (message) => {
      console.log("gamestate update before game start: ", message.data)
      setGameState(message.data);
    };

    addListener("game_state", handleGameStateUpdate);

    return () => {
      removeListener("game_state", handleGameStateUpdate);
    };
  }, [addListener, removeListener]);

  return (
    <GameContext.Provider value={{ gameState }}>
      { children }
    </GameContext.Provider>
  );
};

export const useGame = () => useContext(GameContext);