import Board from "./board/Board";
import Controls from "./board/Controls";
import ChatBox from "../../ChatBox";
import PlayerInfo from "./sidebar/PlayerInfo";
import Console from "./sidebar/Console";

import { Separator } from "../../ui/separator";
import { useGame } from "../../../contexts/GameProvider";

const Game = () => {

  const { gameState } = useGame();
  console.log("gamestate in game component: ", gameState);

  return (
    <div className="flex flex-col gap-4">
      <div className="flex gap-4 flex-row">
        <div className="bg-amber-100 border-solid border-3 border-amber-300 p-4">
          <Board gameState={gameState}/>
          <Controls gameState={gameState}/>
        </div>
        <div className="flex flex-col p-4 gap-4 w-80 bg-amber-100 border-solid border-3 border-amber-300">
          <Console />
          <Separator className="bg-amber-300"/>
          <ChatBox />
          <Separator className="bg-amber-300"/>
          <PlayerInfo gameState={gameState} />
        </div>
      </div>
    </div>
  );
};

export default Game;