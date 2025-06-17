const BOARD_SIZE = 11; // 11 tiles per side
const colorMap = {
  "red": "bg-red-500",
  "orange": "bg-orange-500",
  "yellow": "bg-yellow-500",
  "green": "bg-green-500",
  "blue": "bg-blue-500",
  "purple": "bg-purple-500",
}

const Board = ({ gameState }) => {

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
    <div className="grid grid-cols-11 grid-rows-11 w-[660px] h-[660px] bg-amber-50 relative">
      {tiles.map((tile, index) => {
        const { row, col } = getTilePosition(index);
        const playerOnTile = gameState?.players ? Object.values(gameState.players).filter(p => p.position === index) : [];
        return (
          <div key={index}
               className="absolute w-[60px] h-[60px] border border-amber-600 flex items-center justify-center text-sm"
               style={{
                top: `${row * 60}px`,
                left: `${col * 60}px`,
               }}>
            {playerOnTile.map((player) => (
              <div
                key={player.id}
                className={`w-3 h-3 rounded-full mt-1 ${colorMap[player.color] || "bg-gray-500"}`}>
              </div>
            ))}
          </div>)
      })}
    </div>
  );
};

export default Board;