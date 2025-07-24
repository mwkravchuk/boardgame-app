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
  const [animationCompleted, setAnimationCompleted] = useState(false);

  return (
    <div>
      <div className="flex flex-row gap-4">
        <Board gameState={gameState} setPrompt={setPrompt} setAnimationCompleted={setAnimationCompleted}/>
        <div className="flex flex-col gap-4 p-4 w-80 bg-amber-100">
          <Console />
          <Separator className="bg-amber-300"/>
          <ChatBox />
          <Separator className="bg-amber-300"/>
          <PlayerInfo gameState={gameState} />
        </div>
      </div>
      <DialogManager gameState={gameState} playerId={playerId} prompt={prompt} setPrompt={setPrompt} animationCompleted={animationCompleted}/>
    </div>
  );
};

export default Game;