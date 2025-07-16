import { useState, useEffect } from "react";
import { useNavigate } from "react-router";
import { useLocation } from "react-router-dom";

import { useWebSocket } from "../../contexts/WebSocketProvider";

import ChatBox from "../ChatBox";
import ColorSelector from "../ColorSelector";

const Lobby = () => {

  const location = useLocation();
  const navigate = useNavigate();
  const { sendMessage, addListener, removeListener } = useWebSocket();

  const roomCode = location.state.roomCode;
  const [selectedColor, setSelectedColor] = useState(null);

  const handleColorSelect = (color) => {
    setSelectedColor(color);
    sendMessage("color_selected", color);
  };

  const handleStartGame = () => {
    sendMessage("start_game", "");
  };

  useEffect(() => {
    const handleGameStarted = () => {
      navigate("/game");
    };

    addListener("game_started", handleGameStarted);

    return () => {
      removeListener("game_started", handleGameStarted);
    };
  }, [addListener, removeListener, sendMessage, selectedColor, navigate]);

  return (
    <div className="h-full flex flex-row justify-self-center gap-4 px-10 py-5 p-4 bg-amber-100 border-solid border-3 border-amber-300">
      <div className="flex flex-col gap-1 justify-center">
        <p>ROOM CODE IS: {roomCode}</p>
        <button className="btn-primary" onClick={handleStartGame}>START GAME</button>
        <ColorSelector selectedColor={selectedColor} handleColorSelect={handleColorSelect}/>
      </div>
      <ChatBox />
    </div>
  );
};

export default Lobby;