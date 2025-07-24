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

  const handlePlayerSelect = (id) => {
    setTargetId(id);
    setStep(2);
  };

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
                  Trading with {otherPlayers.find((p) => p.id === targetId)?.displayName}
                </DialogTitle>
                <DialogDescription>
                  Select properties and money to offer/request.
                </DialogDescription>
              </DialogHeader>
            </>
          )}

        </DialogContent>
      </DialogPortal>
    </Dialog>
  );
};

export default TradeDialog;