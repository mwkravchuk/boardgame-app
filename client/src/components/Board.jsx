import { useState, useEffect} from "react";
import { useWebSocket } from "../contexts/WebSocketProvider";


const Board = () => {
  const { sendMessage, addListener, removeListener } = useWebSocket();

  const [dice1, setDice1] = useState(0);
  const [dice2, setDice2] = useState(0);

  useEffect(() => {
    const updateBoardFromRoll = (message) => {
        console.log("message: ", message)
        setDice1(message.dice1);
        setDice2(message.dice2);
    };

    addListener("roll_dice", updateBoardFromRoll);

    return () => {
      removeListener("roll_dice", updateBoardFromRoll);
    };
  }, [addListener, removeListener, dice1, dice2, sendMessage]);
  

  const handleRollDice = () => {
    sendMessage("roll_dice", "");
  };


  return (
    <div>
      <button onClick={handleRollDice}>ROLL DICE</button>
      <p>🎲 Dice 1: {dice1}</p>
      <p>🎲 Dice 2: {dice2}</p>
    </div>
  );
};

export default Board;