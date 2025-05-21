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
    <div>
      <h2>GAME LOBBY</h2>
      <p>ROOM CODE IS: {roomCode}</p>
      <button onClick={handleStartGame}>START GAME</button>
      <ChatBox />
    </div>
  );
};

export default Lobby;