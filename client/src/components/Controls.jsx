import { useState, useEffect } from "react";
import { useWebSocket } from "../contexts/WebSocketProvider";
import { usePlayer } from "../contexts/PlayerProvider";

import styles from "./Controls.module.css";

const Controls = () => {
    const { sendMessage, addListener, removeListener } = useWebSocket();
    const { isMyTurn } = usePlayer();
  
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

    const handleEndTurn = () => {
      sendMessage("new_turn", "");
    };
  
    return (
      <div className={styles.controls}>
        <button onClick={handleRollDice}>ROLL DICE</button>
        <button disabled={!isMyTurn} onClick={handleEndTurn}>END TURN</button>
      </div>
    );
};

export default Controls;