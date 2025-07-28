import { useEffect, useRef } from "react";
import { useWebSocket } from "../../../../contexts/WebSocketProvider";

import BuyPropertyDialog from "./BuyProperty";
import OweRentDialog from "./OweRent";
import TradeDialog from "./Trade";
import BankruptDialog from "./Bankrupt";
import TradeOfferedDialog from "./TradeOffered";

const DialogManager = ({ gameState, playerId, prompt, setPrompt, animationCompleted }) => {

  const { addListener, removeListener, sendMessage } = useWebSocket();

  const lastPromptedTileIndexRef = useRef(null);

  const currentPlayerId = gameState.turnOrder[gameState.currentTurn];
  const isMyTurn = playerId === currentPlayerId;

  useEffect(() => {
    if (!gameState || !playerId || !isMyTurn) return;
    const player = gameState.players?.[playerId];
    if (!player) return;
    const tile = gameState.properties?.[player.position];

    // UNDER THESE CONDITIONS:
    // I CAN BUY THE PROPERTY
    if (tile && tile.isProperty && !tile.isOwned && player.position !== lastPromptedTileIndexRef.current && animationCompleted) {
      console.log("we are on a property, set prompt.");
      setPrompt({
        type: "buy_property",
        data: { property: tile },
      });
      lastPromptedTileIndexRef.current = player.position;
    }

    // I HAVE TO OWE RENT
    if (tile && tile.isProperty && tile.isOwned && tile.ownerId !== playerId && player.position !== lastPromptedTileIndexRef.current && animationCompleted) {
      console.log("we are on a property owned by someone else. pay rent");
      setPrompt({
        type: "owe_rent",
        data: { property: tile, displayName: player.displayName },
      })
      lastPromptedTileIndexRef.current = player.position;
    }

  }, [gameState, playerId, isMyTurn, setPrompt, animationCompleted]);

  // A TRADE OFFER WAS SENT TO ME
  useEffect(() => {
    const handleTradeReceived = (message) => {
      setPrompt({
        type: "trade_offered",
        data: message,
      });
    };

    addListener("trade_offered", handleTradeReceived);

    return () => {
      removeListener("trade_offered", handleTradeReceived);
    }
  }, [addListener, removeListener, setPrompt]);

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
    case "trade":
      return (
        <TradeDialog {...dialogProps} />
      )
    case "bankrupt":
      return (
        <BankruptDialog {...dialogProps}/>
      )
    case "trade_offered":
      return (
        <TradeOfferedDialog {...dialogProps}/>
      )
    default:
      return null;
  }
};

export default DialogManager;