import Board from "./board/Board";
import ChatBox from "../../ChatBox";
import PlayerInfo from "./sidebar/PlayerInfo";
import Console from "./sidebar/Console";
import DialogManager from "./dialogs/DialogManager";

import { Separator } from "../../ui/separator";

import { useGame } from "../../../contexts/GameProvider";
import { usePlayer } from "../../..//contexts/PlayerProvider";
import { useState } from "react";

const Game = () => {

  const { gameState } = useGame();
  console.log("gamestate in game component: ", gameState);

  const { playerId } = usePlayer();
  const [prompt, setPrompt] = useState(null);

  return (
    <div className="flex flex-col gap-4">
      <div className="flex gap-4 flex-row">
        <Board gameState={gameState} setPrompt={setPrompt}/>
        <div className="flex flex-col p-4 gap-4 w-80 bg-amber-100 border-solid border-3 border-amber-300">
          <Console />
          <Separator className="bg-amber-300"/>
          <ChatBox />
          <Separator className="bg-amber-300"/>
          <PlayerInfo gameState={gameState} />
        </div>
      </div>
      <DialogManager gameState={gameState} playerId={playerId} prompt={prompt} setPrompt={setPrompt}/>
    </div>
  );
};

export default Game;