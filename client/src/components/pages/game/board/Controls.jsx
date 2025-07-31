import { useState, useEffect } from "react";
import { useWebSocket } from "../../../../contexts/WebSocketProvider";
import { usePlayer } from "../../../../contexts/PlayerProvider";

import { Button } from "../../../ui/button";
import Dice from "../../../Dice";
import DebugControls from "./DebugControls";

const Controls = ({ gameState, setPrompt, setAnimationCompleted }) => {
  const { addListener, removeListener, sendMessage } = useWebSocket();
  const { playerId } = usePlayer();
  const player = gameState.players?.[playerId];

  const [diceValues, setDiceValues] = useState([]);

  const currentPlayerId = gameState.turnOrder[gameState.currentTurn];
  const isMyTurn = playerId === currentPlayerId;
  const hasRolled = gameState.players[playerId].hasRolled;

  const handleRollDice = () => {
    sendMessage("roll_dice", "");
  };

  const handleEndTurn = () => {
    sendMessage("new_turn", "");
  };

  const handleTrade = () => {
    setPrompt({
      type: "trade",
      data: { displayName: player.displayName, gameState, playerId },
    })
  };

  const handleBankrupt = () => {
    setPrompt({
      type: "bankrupt",
      data: { displayName: player.displayName },
    })
  };

  useEffect(() => {
    const updateDiceValues = (message) => {
      console.log("dice rolled message: ", message);
      setDiceValues(message.data);
    };

    addListener("dice_rolled", updateDiceValues);

    return () => {
      removeListener("dice_rolled", updateDiceValues);
    }
  }, [addListener, removeListener]);

  return (
    <div>
      <Dice values={diceValues} onClick={handleRollDice} setAnimationCompleted={setAnimationCompleted}/>
      <div className="flex flex-row gap-1 justify-center">
        <Button disabled={!isMyTurn || hasRolled} onClick={handleRollDice}>ROLL DICE</Button>
        <Button disabled={!isMyTurn || !hasRolled} onClick={handleEndTurn}>END TURN</Button>
        <Button disabled={!isMyTurn} onClick={handleTrade}>TRADE</Button>
        <Button disabled={!isMyTurn} onClick={handleBankrupt}>BANKRUPT</Button>
      </div>
      <DebugControls />
    </div>
  );
};

export default Controls;