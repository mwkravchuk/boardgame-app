const PlayerInfo = ({ gameState }) => {

  if (!gameState?.players) return null;

  return (
    <div>
      <h2>player info</h2>
      <div>
        {Object.values(gameState.players).map((player) => (
          <div
            key={player.id}
            className="flex flex-col"
          >
            <div className="flex">
              <span>{player.id}</span>
              <span>{player.color}</span>
            </div>
            <div>
              <span>{player.money}</span>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default PlayerInfo;