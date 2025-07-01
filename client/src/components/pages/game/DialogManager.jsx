import { useState, useEffect, useRef } from "react";
import { useWebSocket } from "../../../contexts/WebSocketProvider";

import BuyPropertyDialog from "./dialogs/BuyProperty";

const DialogManager = ({ gameState, playerId, isMyTurn }) => {

  const { sendMessage } = useWebSocket();

  const [prompt, setPrompt] = useState(null);
  const lastPromptedTileIndexRef = useRef(null);

  useEffect(() => {
    console.log("gamestate: ", gameState)
    if (!gameState || !playerId || !isMyTurn) return;
    const player = gameState.players?.[playerId];
    if (!player) return;

    const tile = gameState.properties?.[player.position];
    console.log("tile: ", tile);
    if (tile && tile.isProperty && !tile.isOwned && player.position !== lastPromptedTileIndexRef.current) {
      console.log("we are on a property, set prompt.");
      setPrompt({
        type: "buy_property",
        data: { property: tile },
      });
      lastPromptedTileIndexRef.current = player.position;
    }
  }, [gameState, playerId, isMyTurn]);


  const closePrompt = () => setPrompt(null);

  const handleBuy = () => {
    sendMessage("buy_property", prompt.data.property.name);
    closePrompt();
  };

  const handleAuction = () => {
    sendMessage("auction_property", prompt.data.property.name);
    closePrompt();
  };

  return (
    <>
      {prompt?.type === "buy_property" && (
        <BuyPropertyDialog
          open={true}
          onClose={closePrompt}
          property={prompt.data.property}
          onBuy={handleBuy}
          onAuction={handleAuction}
        />
      )}

    </>
  );
};

export default DialogManager;