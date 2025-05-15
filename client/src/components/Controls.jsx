import { useState, useEffect } from "react";
import { useWebSocket } from "../contexts/WebSocketProvider";
import { usePlayer } from "../contexts/PlayerProvider";

import styles from "./Controls.module.css";

const Controls = () => {
    const { sendMessage, addListener, removeListener } = useWebSocket();
    const { isMyTurn } = usePlayer();
  
    //const [dice1, setDice1] = useState(0);
    //const [dice2, setDice2] = useState(0);
    const [hasRolled, setHasRolled] = useState(false);

    useEffect(() => {
      const updateBoardFromRoll = (message) => {
          //console.log("Roll dice callback:", message)
          //setDice1(message.dice1);
          //setDice2(message.dice2);
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
      <div className={styles.controls}>
        <button disabled={!isMyTurn || hasRolled} onClick={handleRollDice}>ROLL DICE</button>
        <button disabled={!isMyTurn || !hasRolled} onClick={handleEndTurn}>END TURN</button>
      </div>
    );
};

export default Controls;