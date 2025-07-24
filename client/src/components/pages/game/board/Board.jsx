const BOARD_SIZE = 11; // 11 tiles per side

import Controls from "./Controls";
import Tile from "./Tile";

const Board = ({ gameState, setPrompt, setAnimationCompleted }) => {

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
    <div className="relative grid grid-cols-11 grid-rows-11 w-[792px] h-[792px] bg-amber-50">
      {tiles.map((_, index) => {
        const { row, col } = getTilePosition(index);
        const playersOnTile = gameState?.players ? Object.values(gameState.players).filter(p => p.position === index) : [];
        const property = gameState?.properties?.[index];
        return (
          <Tile key={index}
                index={index}
                row={row}
                col={col}
                property={property}
                playersOnTile={playersOnTile}/> 
        )
      })}
      {/* Controls centered on the board */}
      <div className="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 z-10">
        <Controls gameState={gameState} setPrompt={setPrompt} setAnimationCompleted={setAnimationCompleted}/>
      </div>
    </div>
  );
};

export default Board;