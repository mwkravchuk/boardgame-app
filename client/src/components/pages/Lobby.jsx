import { useState, useEffect } from "react";
import { useNavigate } from "react-router";
import { useLocation } from "react-router-dom";

import { useWebSocket } from "../../contexts/WebSocketProvider";

import ChatBox from "../ChatBox";
import ColorSelector from "../ColorSelector";
import { Button } from "../ui/button";
import { Alert, AlertTitle, AlertDescription } from "../ui/alert";
import { AlertCircleIcon } from "lucide-react";

const Lobby = () => {

  const location = useLocation();
  const navigate = useNavigate();
  const { sendMessage, addListener, removeListener } = useWebSocket();

  const [selectedColor, setSelectedColor] = useState(null);
  const [error, setError] = useState(null);

  const roomCode = location.state.roomCode;

  const handleColorSelect = (color) => {
    setSelectedColor(color);
    sendMessage("color_selected", color);
  };

  const handleStartGame = () => {
    sendMessage("start_game", "");
  };

  useEffect(() => {
    const handleGameStarted = () => {
      setError(false);
      navigate("/game");
    };

    const handleGameStartedFail = () => {
      setError(true)
    };

    addListener("game_started", handleGameStarted);
    addListener("game_started_fail", handleGameStartedFail);

    return () => {
      removeListener("game_started", handleGameStarted);
      removeListener("game_started_fail", handleGameStartedFail);
    };
  }, [addListener, removeListener, sendMessage, selectedColor, navigate]);

  return (
    <div className="h-full flex flex-row justify-self-center gap-4 px-10 py-5 p-4 bg-amber-100">
      <div className="flex flex-col gap-1 justify-center">
        <p>ROOM CODE IS: {roomCode}</p>
        <ColorSelector selectedColor={selectedColor} handleColorSelect={handleColorSelect}/>
        <Button onClick={handleStartGame}>START GAME</Button>
        {error && (
          <Alert variant="destructive" className="mt-1">
            <AlertCircleIcon />
            <AlertTitle>Unable to start game.</AlertTitle>
            <AlertDescription>
              <p>Please verify the following:</p>
              <ul className="list-inside list-disc text-sm">
                <li>Room creator starts game</li>
                <li>At least two players</li>
                <li>No overlapping colors</li>
              </ul>
            </AlertDescription>
          </Alert>
        )}
      </div>
      <ChatBox />
    </div>
  );
};

export default Lobby;