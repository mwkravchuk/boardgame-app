const PlayerInfo = ({ gameState }) => {

  if (!gameState?.players || !gameState?.properties) return null;
  const properties = gameState.properties;

  return (
    <div>
      <div>
        {Object.values(gameState.players).map((player) => (
          <div
            key={player.id}
            className="flex flex-col"
          >
            <div className="flex">
              <span>{player.displayName} {player.money}</span>
            </div>
            <div className="flex flex-col text-xs">
                {player.properties?.map((propertyIndex) => (
                  <div key={propertyIndex} className="">
                    {properties[propertyIndex].name}
                  </div>
                ))}
              </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default PlayerInfo;