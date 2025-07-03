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
} from "../../../../components/ui/dialog"
import { Button } from "../../../../components/ui/button";

const OweRentDialog = ({ open, onClose, property, playerName, onPayRent }) => {
  return (
    <Dialog open={open} onOpenChange={onClose}>
      <DialogPortal>
        <DialogOverlay/>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>{playerName}</DialogTitle>
            <DialogDescription>
              This property is owned. You must pay ${property.rent}.
            </DialogDescription>
          </DialogHeader>
          <DialogFooter>
            <Button onClick={onPayRent}>Pay</Button>
            <DialogClose asChild>
              <Button variant="ghost">Cancel</Button>
            </DialogClose>
          </DialogFooter>
        </DialogContent>
      </DialogPortal>
    </Dialog>
  );
};

export default OweRentDialog;