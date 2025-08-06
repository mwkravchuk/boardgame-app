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

const OweRentDialog = ({ open, close, prompt, sendMessage }) => {

  const { rent, displayName } = prompt.data;

  const handlePayRent = () => {
    sendMessage("pay_rent", null);
    close();
  };

  return (
    <Dialog open={open} onOpenChange={close}>
      <DialogPortal>
        <DialogOverlay/>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>{displayName}</DialogTitle>
            <DialogDescription>
              This property is owned. You must pay ${rent}.
            </DialogDescription>
          </DialogHeader>
          <DialogFooter>
            <Button onClick={handlePayRent}>Pay</Button>
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