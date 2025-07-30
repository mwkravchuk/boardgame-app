import { DialogHeader,
         DialogTitle,
         DialogDescription,
         DialogFooter,
         DialogClose } from "../../../../ui/dialog";
import { Button } from "../../../../ui/button";

const PlayerSelectStep = ({ otherPlayers, onSelect, onCancel }) => {
  return (
    <>
      <DialogHeader>
        <DialogTitle>Choose a player to trade with</DialogTitle>
        <DialogDescription>
          {otherPlayers.map(player => (
            <Button key={player.id} onClick={() => onSelect(player.id)}>
              {player.displayName}
            </Button>
          ))}
        </DialogDescription>
      </DialogHeader>
      <DialogFooter>
        <DialogClose asChild>
          <Button variant="ghost" onClick={onCancel}>Cancel</Button>
        </DialogClose>
      </DialogFooter>
    </>
  );
};

export default PlayerSelectStep;