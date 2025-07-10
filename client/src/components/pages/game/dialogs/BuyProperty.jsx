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

const BuyPropertyDialog = ({ open, close, prompt, sendMessage }) => {

  const property = prompt.data.property;

  const handleBuy = () => {
    sendMessage("buy_property", null);
    close();
  };

  const handleAuction = () => {
    sendMessage("auction_property", null);
    close();
  };

  return (
    <Dialog open={open} onOpenChange={close}>
      <DialogPortal>
        <DialogOverlay/>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>{property.name}</DialogTitle>
            <DialogDescription>
              This property costs ${property.price}. Would you like to buy or auction it?
            </DialogDescription>
          </DialogHeader>
          <DialogFooter>
            <Button onClick={handleBuy}>Buy</Button>
            <Button variant="outline" onClick={handleAuction}>Auction</Button>
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