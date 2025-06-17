import { useState, useEffect } from "react";

import Board from "./Board";
import Controls from "./Controls";
import ChatBox from "../../ChatBox";
import PlayerInfo from "./PlayerInfo";
import Console from "./Console";

import { Separator } from "../../ui/separator";

import { useWebSocket } from "../../../contexts/WebSocketProvider";

const Game = () => {

  const { addListener, removeListener } = useWebSocket();
  const [gamestate, setGamestate] = useState(null);

  useEffect(() => {
    const updateGamestate = (message) => {
      console.log("game_state msg data:", message.data);
      setGamestate(message.data);
    };

    addListener("game_state", updateGamestate);

    return () => {
      removeListener("game_state", updateGamestate);
    };
  }, [addListener, removeListener]);

  return (
    <div className="flex flex-col gap-4">
      <div className="flex gap-4 flex-row">
        <div className="bg-amber-100 border-solid border-3 border-amber-300 p-4">
          <Board gameState={gamestate}/>
          <Controls gameState={gamestate}/>
        </div>
        <div className="flex flex-col p-4 gap-4 w-80 bg-amber-100 border-solid border-3 border-amber-300">
          <Console />
          <Separator className="bg-amber-300"/>
          <ChatBox />
          <Separator className="bg-amber-300"/>
          <PlayerInfo />
        </div>
      </div>
    </div>
  );
};

export default Game;