import { useWebSocket } from "../../../../contexts/WebSocketProvider";
import { usePlayer } from "../../../../contexts/PlayerProvider";

const Controls = ({ gameState }) => {
    const { sendMessage } = useWebSocket();
    const { playerId } = usePlayer();

    const isMyTurn = playerId === gameState.turnOrder[gameState.currentTurn];
    const hasRolled = false;    
  
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