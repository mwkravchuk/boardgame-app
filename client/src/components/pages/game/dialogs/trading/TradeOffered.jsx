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

const TradeOfferedDialog = ({ open, close, prompt, sendMessage }) => {

  const { fromPlayer, offerMoney, requestMoney, offerProps, requestProps } = prompt.data;

  const handleAccept = () => {
    sendMessage("respond_to_trade", "accept");
    close();
  };

  const handleReject = () => {
    sendMessage("respond_to_trade", "reject");
    close();
  };

  return (
    <Dialog open={open} onOpenChange={close}>
      <DialogPortal>
        <DialogOverlay/>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>You have been offered a trade by {fromPlayer}</DialogTitle>
          </DialogHeader>
          YOU RECEIVE {offerMoney}, YOU RECEIVE {requestMoney}
          <DialogFooter>
            <Button onClick={handleAccept}>ACCEPT</Button>
            <Button onClick={handleReject}>REJECT</Button>
          </DialogFooter>
        </DialogContent>
      </DialogPortal>
    </Dialog>
  );
};

export default TradeOfferedDialog;