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
} from "../../../../ui/dialog"
import { Button } from "../../../../ui/button";

import PropertyManager from "./PropertyManager";

const ManagePropertiesDialog = ({ open, close, gameState, prompt, sendMessage }) => {
  const { playerId } = prompt.data;
  const player = gameState.players[playerId];
  const properties = gameState.properties;

  return (
    <Dialog open={open} onOpenChange={close}>
      <DialogPortal>
        <DialogOverlay/>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Manage Properties</DialogTitle>
            <DialogDescription>Buy/sell houses or mortgage your properties.</DialogDescription>
          </DialogHeader>

          <div className="flex flex-col gap-2 max-h-80 overflow-y-auto">
            {player.properties.map((propIdx) => (
              <PropertyManager
                key={propIdx}
                property={properties[propIdx]}
                propertyIdx={propIdx}
                sendMessage={sendMessage}
              />
            ))}
          </div>

          <DialogFooter>
            <DialogClose asChild>
              <Button variant="ghost">Cancel</Button>
            </DialogClose>
          </DialogFooter>
        </DialogContent>
      </DialogPortal>
    </Dialog>
  );
};

export default ManagePropertiesDialog;