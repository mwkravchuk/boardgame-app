import { useState, useEffect } from "react";
import { useWebSocket } from "../../../contexts/WebSocketProvider";
import { usePlayer } from "../../../contexts/PlayerProvider";

const Controls = () => {
    const { sendMessage, addListener, removeListener } = useWebSocket();
    const { isMyTurn } = usePlayer();
  
    const [hasRolled, setHasRolled] = useState(false);

    useEffect(() => {
      const updateBoardFromRoll = (message) => {
        setHasRolled(message.data)
      };
  
      const resetRollButton = (message) => {
        //console.log("Reset Roll callback:", message)
        setHasRolled(message.data)
      };

      addListener("roll_dice", updateBoardFromRoll);
      addListener("reset_roll_button", resetRollButton);
  
      return () => {
        removeListener("roll_dice", updateBoardFromRoll);
        removeListener("reset_roll_button", resetRollButton);
      };
    }, [addListener, removeListener]);
    
  
    const handleRollDice = () => {
      sendMessage("roll_dice", "");
    };

    const handleEndTurn = () => {
      sendMessage("new_turn", "");
    };
  
    return (
      <div className="flex flex-row gap-4 bg-amber-200">
        <button className="btn-primary" disabled={!isMyTurn || hasRolled} onClick={handleRollDice}>ROLL DICE</button>
        <button className="btn-primary" disabled={!isMyTurn || !hasRolled} onClick={handleEndTurn}>END TURN</button>
      </div>
    );
};

export default Controls;