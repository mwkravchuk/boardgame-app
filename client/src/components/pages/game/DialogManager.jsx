import { useState, useEffect, useRef } from "react";
import { useWebSocket } from "../../../contexts/WebSocketProvider";

import BuyPropertyDialog from "./dialogs/BuyProperty";
import OweRentDialog from "./dialogs/OweRent";

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

    if (tile && tile.isProperty && tile.isOwned && tile.ownerId !== playerId) {
      console.log("we are on a property owned by someone else. pay rent");
      setPrompt({
        type: "owe_rent",
        data: { property: tile, player },
      })
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

  const handlePayRent = () => {
    sendMessage("pay_rent", null);
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

      {prompt?.type === "owe_rent" && (
        <OweRentDialog
          open={true}
          onClose={closePrompt}
          property={prompt.data.property}
          playerName={gameState.players?.[playerId].displayName}
          onPayRent={handlePayRent}
        />
      )}
    </>
  );
};

export default DialogManager;