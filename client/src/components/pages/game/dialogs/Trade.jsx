import { useState, useEffect } from "react";
import { Dialog, DialogPortal, DialogOverlay, DialogContent } from "../../../ui/dialog"

import PlayerSelectStep from "./trading/PlayerSelectStep";
import TradeFormStep from "./trading/TradeFormStep";

const TradeDialog = ({ open, close, prompt, addListener, removeListener, sendMessage }) => {
  const [step, setStep] = useState(1);
  const [targetId, setTargetId] = useState(null);
  const playerId = prompt.data.playerId;

  const players = prompt.data.gameState.players;
  const properties = prompt.data.gameState.properties;
  const otherPlayers = Object.values(players).filter(p => p.id !== playerId);
  const selfPlayer = players[playerId];
  const otherPlayer = otherPlayers.find((p) => p.id === targetId);

  console.log("step: ", step)

  useEffect(() => {
    const handleAccepted = () => {
      // do stuff maybe
      close();
    };
    const handleRejected = () => {
      // do stuff maybe
      close();
    };

    addListener("trade_accepted", handleAccepted);
    addListener("trade_rejected", handleRejected);

  }, [addListener, removeListener, close]);


  return (
    <Dialog open={open} onOpenChange={close}>
      <DialogPortal>
        <DialogOverlay/>
        <DialogContent>
          {step === 1 && (
            <PlayerSelectStep otherPlayers={otherPlayers}
                              onSelect={(id) => {
                                setTargetId(id);
                                setStep(2);
                              }}
                              onCancel={close} />
          )}
          {step === 2 && (
            <TradeFormStep selfPlayer={selfPlayer}
                           otherPlayer={otherPlayer}
                           targetId={targetId}
                           setStep={setStep}
                           properties={properties}
                           close={close}
                           sendMessage={sendMessage}
                           />
          )}
          {step === 3 && (
            <div>
              <span>Waiting on a response!</span>
              <div className="w-8 h-8 border-4 boder-gray-300 border-t-blue-500 rounded-full animate-spin"></div>            
            </div>
          )}
        </DialogContent>
      </DialogPortal>
    </Dialog>
  );
};

export default TradeDialog;