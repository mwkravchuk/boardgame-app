import { useState } from "react";
import { DialogHeader, DialogTitle, DialogDescription, DialogFooter, DialogClose } from "../../../../ui/dialog";
import { Button } from "../../../../ui/button";

import PropertyTradeSelector from "./PropertyTradeSelector";


const TradeFormStep = ({ selfPlayer, otherPlayer, targetId, setStep, properties, close, sendMessage }) => {
  const [myOfferMoney, setMyOfferMoney] = useState(0);
  const [theirOfferMoney, setTheirOfferMoney] = useState(0);
  const [myOfferProps, setMyOfferProps] = useState([]);
  const [theirOfferProps, setTheirOfferProps] = useState([]);

  // Toggle properties in trade offer (for either my own props or their props)
  const toggle = (setter) => (propertyIdx) => {
    setter((prev) => prev.includes(propertyIdx) ? prev.filter(i => i !== propertyIdx) : [...prev, propertyIdx]);
  };

  const handlePropose = () => {
    sendMessage("propose_trade", {
      targetId,
      myOfferMoney,
      theirOfferMoney,
      myOfferProps,
      theirOfferProps,
    });
    setStep(3);
  };

  return (
    <>
      <DialogHeader>
        <DialogTitle>
          Trading with {otherPlayer.displayName}
        </DialogTitle>
        <DialogDescription>
          Select properties and money to offer/request.
        </DialogDescription>
      </DialogHeader>

      <div className="flex flex-row">
        {/* WHAT YOU OFFER */}
        <div className="mt-4">
          <label>Your money offer: ${myOfferMoney}</label>
          <input
            type="range"
            min="0"
            max={selfPlayer?.money}
            value={myOfferMoney}
            onChange={(e) => setMyOfferMoney(Number(e.target.value))}/>
          <PropertyTradeSelector
            ownedIndices={selfPlayer?.properties || []}
            properties={properties}
            selectedIndices={myOfferProps}
            onToggle={toggle(setMyOfferProps)}/>
        </div>

        {/* WHAT THEY OFFER IN RETURN */}
        <div className="mt-4">
          <label>Their money offer: ${theirOfferMoney}</label>
          <input
            type="range"
            min="0"
            max={otherPlayer?.money}
            value={theirOfferMoney}
            onChange={(e) => setTheirOfferMoney(Number(e.target.value))}/>
            <PropertyTradeSelector
              ownedIndices={otherPlayer?.properties || []}
              properties={properties}
              selectedIndices={theirOfferProps}
              onToggle={toggle(setTheirOfferProps)}
            />
        </div>
      </div>

      <DialogFooter>
        <Button onClick={handlePropose}>Propose Trade</Button>
        <DialogClose asChild>
          <Button variant="ghost" onClick={close}>Cancel</Button>
        </DialogClose>
      </DialogFooter>
    </>
  );
};

export default TradeFormStep;