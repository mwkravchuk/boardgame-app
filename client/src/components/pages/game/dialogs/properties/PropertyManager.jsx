import { Button } from "../../../../ui/button";

const PropertyManager = ({ property, sendMessage }) => {
  const { name, numHouses, isMortgaged } = property;

  const handleBuyHouse = () => {
    sendMessage("buy_house", property.index);
  };

  const handleSellHouse = () => {
    sendMessage("sell_house", property.index);
  };

  const handleToggleMortgage = () => {
    sendMessage("toggle_mortgage", property.index);
  };

  return (
    <div className="border p-2 rounded flex justify-between items-center">
      <div>
        <p className="font-bold">{name}</p>
        <p>Houses: {numHouses}</p>
        {isMortgaged && <p className="text-red-600">Mortgaged</p>}
      </div>
      <div className="flex gap-1">
        <Button onClick={handleBuyHouse} disabled={isMortgaged}>+ House</Button>
        <Button onClick={handleSellHouse} disabled={numHouses === 0}>- House</Button>
        <Button variant="outline" onClick={handleToggleMortgage}>
          {isMortgaged ? "Unmortgage" : "Mortgage"}
        </Button>
      </div>
    </div>
  );
};

export default PropertyManager;