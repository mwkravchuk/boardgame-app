import { useEffect } from "react";
import { useNavigate } from "react-router";
import { useLocation } from "react-router-dom";

import { useWebSocket } from "../../contexts/WebSocketProvider";

import ChatBox from "../ChatBox";

const Lobby = () => {

  const location = useLocation();
  const navigate = useNavigate();
  const { sendMessage, addListener, removeListener } = useWebSocket();

  const roomCode = location.state.roomCode;

  useEffect(() => {
    const handleGameStarted = () => {
      console.log("Game started. Go to game board mate");
      navigate("/game");
    };

    addListener("game_started", handleGameStarted);

    return () => {
      removeListener("game_started", handleGameStarted);
    };
  }, [addListener, removeListener, navigate]);

  const handleStartGame = () => {
    sendMessage("start_game", "");
  };

  return (
    <div className="h-full flex flex-row justify-self-center gap-4 px-10 py-5 p-4 bg-amber-100 border-solid border-3 border-amber-300">
      <div className="flex flex-col gap-1 justify-center">
        <p>ROOM CODE IS: {roomCode}</p>
        <button className="btn-primary" onClick={handleStartGame}>START GAME</button>
      </div>
      <ChatBox />
    </div>
  );
};

export default Lobby;