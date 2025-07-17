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

const BankruptDialog = ({ open, close, prompt, sendMessage }) => {

  const handleBankrupt = () => {
    sendMessage("bankrupt", null);
    close();
  };

  return (
    <Dialog open={open} onOpenChange={close}>
      <DialogPortal>
        <DialogOverlay/>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>{prompt.data.displayName}</DialogTitle>
            <DialogDescription>
              Are you sure you want to bankrupt?
            </DialogDescription>
          </DialogHeader>
          <DialogFooter>
            <Button onClick={handleBankrupt}>Bankrupt</Button>
            <DialogClose asChild>
              <Button variant="ghost">Cancel</Button>
            </DialogClose>
          </DialogFooter>
        </DialogContent>
      </DialogPortal>
    </Dialog>
  );
};

export default BankruptDialog;