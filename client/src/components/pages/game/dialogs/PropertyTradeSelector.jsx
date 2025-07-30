import { Button } from "../../../ui/button";

const PropertyTradeSelector = ({
  ownedIndices,
  properties,
  selectedIndices,
  onToggle
}) => {
  return (
    <div className="flex flex-wrap gap-2">
      {ownedIndices.map((propertyIdx) => {
        const isSelected = selectedIndices.includes(propertyIdx);
        const name = properties[propertyIdx]?.name;

        return (
          <Button
            key={propertyIdx}
            className={`px-2 py-1 border rounded ${
              isSelected ? "bg-blue-500 text-white" : "bg-gray-200"
            }`}
            onClick={() => onToggle(propertyIdx)}
          >
            {name}
          </Button>
        );
      })}
    </div>
  );
};

export default PropertyTradeSelector;