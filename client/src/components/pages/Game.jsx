import { usePlayer } from "../../contexts/PlayerProvider";

import Board from "../Board";
import Controls from "../Controls";
import ChatBox from "../ChatBox";

const Game = () => {

  const { playerId, currentTurnId } = usePlayer();

  return (
    <>
      <div>
        <p>playerID:{playerId}</p>
        <p>currID:{currentTurnId}</p>
      </div>
      <div>
        <div>
          <Board />
          <Controls />
        </div>
        <div>
          <ChatBox />
        </div>
      </div>
    </>
  );
};

export default Game;