import { useState, useEffect } from "react";
import { useWebSocket } from "../../../../contexts/WebSocketProvider";

import BuyPropertyDialog from "./forced/BuyProperty";
import OweRentDialog from "./forced/OweRent";
import TradeDialog from "./trading/Trade";
import BankruptDialog from "./Bankrupt";
import TradeOfferedDialog from "./trading/TradeOffered";
import ManagePropertiesDialog from "./properties/ManageProperties";

const DialogManager = ({ gameState, prompt, setPrompt, animationCompleted }) => {

  const { addListener, removeListener, sendMessage } = useWebSocket();
  const [pendingMessage, setPendingMessage] = useState(null);

  // set dialog prompt based on state received from backend
  useEffect(() => {
    const handleOweRent = (message) => {
      setPendingMessage({
        type: "owe_rent",
        data: {
          rent: message.data.rent,
          displayName: message.data.displayName,
        },
      });
    };

    const handleCanBuyProperty = (message) => {
      setPendingMessage({
        type: "can_buy_property",
        data: {
          property: message.data.property,
        }
      });
    };

    const handleTradeReceived = (message) => {
      /* setPrompt over setPendingmessage
         since it is weird behavior if the dice hasnt been rolled */
      setPrompt({
        type: "trade_offered",
        data: {
          fromPlayer:   message.data.fromPlayer,
          offerMoney:   message.data.offerMoney,
          requestMoney: message.data.requestMoney,
          offerProps:   message.data.offerProps,
          requestProps: message.data.requestProps,
        },
      });
    };

    addListener("owe_rent", handleOweRent);
    addListener("can_buy_property", handleCanBuyProperty);
    addListener("trade_offered", handleTradeReceived);

    return () => {
      removeListener("owe_rent", handleOweRent);
      removeListener("can_buy_property", handleCanBuyProperty);
      removeListener("trade_offered", handleTradeReceived);
    }
  }, [addListener, removeListener, setPrompt]);

  // Show FORCED DIALOGS when animation completes
  useEffect(() => {
    if (animationCompleted && pendingMessage) {
      setPrompt(pendingMessage);
      setPendingMessage(null);
    }
  }, [animationCompleted, pendingMessage, setPrompt]);

  const closePrompt = () => setPrompt(null);

  const dialogProps = {
    open: true,
    close: closePrompt,
    gameState,
    prompt,
    addListener,
    removeListener,
    sendMessage,
  };

  switch (prompt?.type) {
    case "can_buy_property":
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
    case "manage_properties":
      return (
        <ManagePropertiesDialog {...dialogProps}/>
      )
    default:
      return null;
  }
};

export default DialogManager;