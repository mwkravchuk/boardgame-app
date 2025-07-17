const Tile = ({ index, row, col, property, playersOnTile, propertyColorMap, playerColorMap }) => {
  return (
    <div key={index}
         className={`absolute w-[70px] h-[70px] border-t-[10px] border ${propertyColorMap[property?.color] || "border-slate-300"} flex text-sm flex-col text-center`}
         style={{
            top: `${row * 70}px`,
            left: `${col * 70}px`,}}>
      {/* Property Name */}
      {property?.name && (<div className="text-[8px] font-bold w-full">{property.name}</div>)}
      {/* Property Price */}
      {property?.price && (<div className="text-[8px]">{property.price}</div>)}
      {/* Players on tile */}
      <div className={"flex flex-row gap-1"}>
        {playersOnTile.map((player) => (
          <div
            key={player.id}
            className={`w-3 h-3 rounded-full mt-1 ${playerColorMap[player.color] || "bg-gray-500"}`}>
          </div>
        ))}
      </div>
    </div>
  );
};

export default Tile;