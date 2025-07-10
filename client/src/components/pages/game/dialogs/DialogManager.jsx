import { useState, useEffect, useRef } from "react";
import { useWebSocket } from "../../../../contexts/WebSocketProvider";

import BuyPropertyDialog from "./BuyProperty";
import OweRentDialog from "./OweRent";

const DialogManager = ({ gameState, playerId, isMyTurn }) => {

  const { sendMessage } = useWebSocket();

  const [prompt, setPrompt] = useState(null);
  const lastPromptedTileIndexRef = useRef(null);

  useEffect(() => {
    if (!gameState || !playerId || !isMyTurn) return;
    const player = gameState.players?.[playerId];
    if (!player) return;
    const tile = gameState.properties?.[player.position];

    // UNDER THESE CONDITIONS:
    // BUY PROPERTY 
    if (tile && tile.isProperty && !tile.isOwned && player.position !== lastPromptedTileIndexRef.current) {
      console.log("we are on a property, set prompt.");
      setPrompt({
        type: "buy_property",
        data: { property: tile },
      });
      lastPromptedTileIndexRef.current = player.position;
    }

    // OWE RENT
    if (tile && tile.isProperty && tile.isOwned && tile.ownerId !== playerId && player.position !== lastPromptedTileIndexRef.current) {
      console.log("we are on a property owned by someone else. pay rent");
      setPrompt({
        type: "owe_rent",
        data: { property: tile, displayName: player.displayName },
      })
      lastPromptedTileIndexRef.current = player.position;
    }

  }, [gameState, playerId, isMyTurn]);

  const closePrompt = () => setPrompt(null);

  const dialogProps = {
    open: true,
    close: closePrompt,
    prompt,
    sendMessage,
  };

  switch (prompt?.type) {
    case "buy_property":
      return (
        <BuyPropertyDialog {...dialogProps}/>
      );
    case "owe_rent":
      return (
        <OweRentDialog {...dialogProps} />
      );

    default:
      return null;
  }
};

export default DialogManager;