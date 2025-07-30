import { useState } from "react";

import {
  Dialog,
  DialogPortal,
  DialogOverlay,
  DialogContent,
  DialogHeader,
  DialogFooter,
  DialogTitle,
  DialogDescription,
  DialogClose,
} from "../../../ui/dialog"
import { Button } from "../../../ui/button";

import PropertyTradeSelector from "./trading/PropertyTradeSelector";
import PlayerSelectStep from "./trading/PlayerSelectStep";

const TradeDialog = ({ open, close, prompt, sendMessage }) => {
  const [step, setStep] = useState(1);
  const [targetId, setTargetId] = useState(null);
  const playerId = prompt.data.playerId;

  const players = prompt.data.gameState.players;
  const properties = prompt.data.gameState.properties;
  const otherPlayers = Object.values(players).filter(p => p.id !== playerId);
  const selfPlayer = players[playerId];
  const otherPlayer = otherPlayers.find((p) => p.id === targetId);

  console.log("self player: ", selfPlayer);
  console.log("other player: ", otherPlayer);

  const selfMoney = selfPlayer?.money;
  const theirMoney = otherPlayer?.money;
  const selfProps = selfPlayer?.properties;
  const theirProps = otherPlayer?.properties;

  const handlePropose = () => {
    sendMessage("propose_trade", {
      targetId,
      myOfferMoney,
      theirOfferMoney,
      myOfferProps,
      theirOfferProps,
    });
    close();
  };

    // Step 2 state
  const [myOfferMoney, setMyOfferMoney] = useState(0);
  const [theirOfferMoney, setTheirOfferMoney] = useState(0);
  const [myOfferProps, setMyOfferProps] = useState([]);
  const [theirOfferProps, setTheirOfferProps] = useState([]);

  const toggleMyOfferProps = (propertyIdx) => {
    setMyOfferProps((prev) =>
      prev.includes(propertyIdx)
        ? prev.filter((i) => i !== propertyIdx)
        : [...prev, propertyIdx]
    );
  };

  const toggleTheirOfferProps = (propertyIdx) => {
    setTheirOfferProps((prev) =>
      prev.includes(propertyIdx)
        ? prev.filter((i) => i !== propertyIdx)
        : [...prev, propertyIdx]
    );
  };

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
            <>
              <DialogHeader>
                <DialogTitle>
                  Trading with {otherPlayer.displayName}
                </DialogTitle>
                <DialogDescription>
                  Select properties and money to offer/request.
                </DialogDescription>
              </DialogHeader>

              <div className="flex flex-row">
                {/* YOUR OFFER */}
                <div className="mt-4">
                  <label>Your money offer: ${myOfferMoney}</label>
                  <input
                    type="range"
                    min="0"
                    max={selfMoney}
                    value={myOfferMoney}
                    onChange={(e) => setMyOfferMoney(Number(e.target.value))}/>
                  <PropertyTradeSelector
                    ownedIndices={selfProps}
                    properties={properties}
                    selectedIndices={myOfferProps}
                    onToggle={toggleMyOfferProps}/>
                </div>

                {/* THEIR STUFF YOU WANT */}
                <div className="mt-4">
                  <label>Their money offer: ${theirOfferMoney}</label>
                  <input
                    type="range"
                    min="0"
                    max={theirMoney}
                    value={theirOfferMoney}
                    onChange={(e) => setTheirOfferMoney(Number(e.target.value))}/>
                    <PropertyTradeSelector
                    ownedIndices={theirProps}
                    properties={properties}
                    selectedIndices={theirOfferProps}
                    onToggle={toggleTheirOfferProps}/>
                </div>
              </div>

              <DialogFooter>
                <Button onClick={handlePropose}>Propose Trade</Button>
                <DialogClose asChild>
                  <Button variant="ghost">Cancel</Button>
                </DialogClose>
              </DialogFooter>
            </>
          )}
        </DialogContent>
      </DialogPortal>
    </Dialog>
  );
};

export default TradeDialog;