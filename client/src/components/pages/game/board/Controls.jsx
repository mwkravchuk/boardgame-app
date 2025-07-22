import { useWebSocket } from "../../../../contexts/WebSocketProvider";
import { usePlayer } from "../../../../contexts/PlayerProvider";

import { Button } from "../../../ui/button";

const Controls = ({ gameState, setPrompt }) => {
  const { sendMessage } = useWebSocket();
  const { playerId } = usePlayer();
  const player = gameState.players?.[playerId];

  const currentPlayerId = gameState.turnOrder[gameState.currentTurn];
  console.log("current player id: ", currentPlayerId);
  const isMyTurn = playerId === currentPlayerId;
  console.log("ismyturn", isMyTurn);
  const hasRolled = gameState.players[playerId].hasRolled;
  console.log("has rolled: ", hasRolled);

  const handleRollDice = () => {
    sendMessage("roll_dice", "");
  };

  const handleEndTurn = () => {
    sendMessage("new_turn", "");
  };

  const handleTrade = () => {
    setPrompt({
      type: "initiate_trade",
      data: { displayName: player.displayName },
    })
  };

  const handleBankrupt = () => {
    setPrompt({
      type: "bankrupt",
      data: { displayName: player.displayName },
    })
  };

  return (
    <div className="flex flex-row gap-1">
      <Button disabled={!isMyTurn || hasRolled} onClick={handleRollDice}>ROLL DICE</Button>
      <Button disabled={!isMyTurn || !hasRolled} onClick={handleEndTurn}>END TURN</Button>
      <Button disabled={!isMyTurn} onClick={handleTrade}>TRADE</Button>
      <Button disabled={!isMyTurn} onClick={handleBankrupt}>BANKRUPT</Button>
    </div>
  );
};

export default Controls;