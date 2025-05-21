import { usePlayer } from "../../contexts/PlayerProvider";

import Board from "../Board";
import Controls from "../Controls";
import ChatBox from "../ChatBox";

import styles from "./Game.module.css";

const Game = () => {

  const { playerId, currentTurnId } = usePlayer();

  return (
    <>
      <div>
        <p>playerID:{playerId}</p>
        <p>currID:{currentTurnId}</p>
      </div>
      <div className={styles.gameContainer}>
        <div className={styles.leftCol}>
          <Board />
          <Controls />
        </div>
        <div className={styles.rightCol}>
          <ChatBox />
        </div>
      </div>
    </>
  );
};

export default Game;