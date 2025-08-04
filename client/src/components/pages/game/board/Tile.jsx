const playerColorMap = {
  "red": "bg-red-500",
  "orange": "bg-orange-500",
  "yellow": "bg-yellow-500",
  "green": "bg-green-500",
  "blue": "bg-blue-500",
  "purple": "bg-purple-500",
};
const propertyColorMap = {
  "brown": "border-orange-900",
  "black": "border-neutral-900",
  "light blue": "border-blue-300",
  "pink": "border-fuchsia-500",
  "orange": "border-orange-400",
  "red": "border-red-600",
  "yellow": "border-yellow-300",
  "green": "border-green-600",
  "blue": "border-blue-700",
};

const Tile = ({ index, row, col, property, playersOnTile }) => {
  return (
    <div key={index}
         className={`absolute w-[72px] h-[72px] border-t-[8px] border ${propertyColorMap[property?.color] || "border-slate-300"} flex text-sm flex-col justify-between text-center ${property?.isMortgaged ? "bg-slate-500": ""}`}
         style={{
            top: `${row * 72}px`,
            left: `${col * 72}px`,}}>
      {/* Property Name */}
      {property?.name && (<div className="text-[10px] font-bold mx-0.5">{property.name}</div>)}
      {/* Players on tile */}
      <div className={"flex flex-row gap-1"}>
        {playersOnTile.map((player) => (
          <div
            key={player.id}
            className={`w-3 h-3 rounded-full mt-1 ${playerColorMap[player.color] || "bg-gray-500"}`}>
          </div>
        ))}
      </div>
      {/* Property Price */}
      {property?.price && (<div className="text-[10px]">{property.price}</div>)}
    </div>
  );
};

export default Tile;