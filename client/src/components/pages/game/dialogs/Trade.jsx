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

const TradeDialog = ({ open, close, prompt, sendMessage }) => {
  const [step, setStep] = useState(1);
  const [targetId, setTargetId] = useState(null);
  const playerId = prompt.data.playerId;

  const players = prompt.data.gameState.players;
  const otherPlayers = Object.values(players).filter(p => p.id !== playerId);
  const selfPlayer = players[playerId];
  const otherPlayer = otherPlayers.find((p) => p.id === targetId);

  const selfMoney = selfPlayer?.money;
  const theirMoney = otherPlayer?.money;

  const handlePlayerSelect = (id) => {
    setTargetId(id);
    setStep(2);
  };

  const handlePropose = () => {
    sendMessage("propose_trade", {
      targetId,
      myOfferMoney,
      theirOfferMoney,
    });
    close();
  };

    // Step 2 state
  const [myOfferMoney, setMyOfferMoney] = useState(0);
  const [theirOfferMoney, setTheirOfferMoney] = useState(0);

  return (
    <Dialog open={open} onOpenChange={close}>
      <DialogPortal>
        <DialogOverlay/>
        <DialogContent>
          {step === 1 && (
            <>
              <DialogHeader>
                <DialogTitle>Choose a player to trade with</DialogTitle>
                <DialogDescription>
                  {otherPlayers.map(player => (
                    <Button key={player.id} onClick={() => handlePlayerSelect(player.id)}>
                      {player.displayName}
                    </Button>
                  ))}
                </DialogDescription>
              </DialogHeader>
              <DialogFooter>
                <DialogClose asChild>
                  <Button variant="ghost">Cancel</Button>
                </DialogClose>
              </DialogFooter>
            </>
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

              {/* YOUR MONEY */}
              <div className="mt-4">
                <label>Your money offer: ${myOfferMoney}</label>
                <input
                  type="range"
                  min="0"
                  max={selfMoney}
                  value={myOfferMoney}
                  onChange={(e) => setMyOfferMoney(Number(e.target.value))}
                />
              </div>

              {/* THEIR MONEY */}
              <div className="mt-4">
                <label>Their money offer: ${theirOfferMoney}</label>
                <input
                  type="range"
                  min="0"
                  max={theirMoney}
                  value={theirOfferMoney}
                  onChange={(e) => setTheirOfferMoney(Number(e.target.value))}
                />
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