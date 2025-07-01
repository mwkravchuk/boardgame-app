const BOARD_SIZE = 11; // 11 tiles per side
const playerColorMap = {
  "red": "bg-red-500",
  "orange": "bg-orange-500",
  "yellow": "bg-yellow-500",
  "green": "bg-green-500",
  "blue": "bg-blue-500",
  "purple": "bg-purple-500",
}
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
}

import { usePlayer } from "../../../contexts/PlayerProvider";
import DialogManager from "./DialogManager";

const Board = ({ gameState }) => {

  const { playerId, isMyTurn } = usePlayer();

  const numTotalTiles = BOARD_SIZE * 4 - 4;
  const tiles = Array.from({ length: numTotalTiles }, (_, i) => `Tile ${i + 1}`);

  const getTilePosition = (index) => {
    if (index < 11) {
      return { row: 0, col: index };
    } else if (index < 20) {
      return { row: index - 10, col: 10 };
    } else if (index < 31) {
      return { row: 10, col: 30 - index};
    } else {
      return {row: 40 - index, col: 0 };
    }
  };

  return (
    <div className="grid grid-cols-11 grid-rows-11 w-[770px] h-[770px] bg-amber-50 relative">
      {tiles.map((tile, index) => {
        const { row, col } = getTilePosition(index);
        const playerOnTile = gameState?.players ? Object.values(gameState.players).filter(p => p.position === index) : [];
        const property = gameState?.properties?.[index];
        return (
          <div key={index}>
            <div
                className={`absolute w-[70px] h-[70px] border-t-[10px] border ${propertyColorMap[property?.color] || "border-slate-300"} flex items-center justify-center text-sm flex-col text-center`}
                style={{
                  top: `${row * 70}px`,
                  left: `${col * 70}px`,
                }}>
              {/* Property Name */}
              {property?.name && (<div className="text-[10px] font-bold w-full">{property.name}</div>)}

              {/* Property Price */}
              {property?.price && (<div className="text-[10px]">{property.price}</div>)}

              {/* Players on tile */}
              {playerOnTile.map((player) => (
                <div
                  key={player.id}
                  className={`w-3 h-3 rounded-full mt-1 ${playerColorMap[player.color] || "bg-gray-500"}`}>
                </div>
              ))}
            </div>
          </div>
        )
      })}
      <DialogManager gameState={gameState} playerId={playerId} isMyTurn={isMyTurn}/>
    </div>
  );
};

export default Board;