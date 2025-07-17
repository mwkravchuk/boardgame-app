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

const InitiateTradeDialog = ({ open, close, prompt, sendMessage }) => {

  const handleInitiateTrade = () => {
    sendMessage("initiate_trade", null);
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
              yo
            </DialogDescription>
          </DialogHeader>
          <DialogFooter>
            <Button onClick={handleInitiateTrade}>trade</Button>
            <DialogClose asChild>
              <Button variant="ghost">Cancel</Button>
            </DialogClose>
          </DialogFooter>
        </DialogContent>
      </DialogPortal>
    </Dialog>
  );
};

export default InitiateTradeDialog;