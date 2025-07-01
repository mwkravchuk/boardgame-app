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

const BuyPropertyDialog = ({ open, onClose, property, onBuy, onAuction }) => {
  return (
    <Dialog open={open} onOpenChange={onClose}>
      <DialogPortal>
        <DialogOverlay/>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>{property?.name}</DialogTitle>
            <DialogDescription>
              This property costs ${property?.price}. Would you like to buy or auction it?
            </DialogDescription>
          </DialogHeader>
          <DialogFooter>
            <Button onClick={onBuy}>Buy</Button>
            <Button variant="outline" onClick={onAuction}>Auction</Button>
            <DialogClose asChild>
              <Button variant="ghost">Cancel</Button>
            </DialogClose>
          </DialogFooter>
        </DialogContent>
      </DialogPortal>
    </Dialog>
  );
};

export default BuyPropertyDialog;