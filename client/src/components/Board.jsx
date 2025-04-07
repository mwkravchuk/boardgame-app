const Board = ({ ws }) => {

  const handleRollDice = () => {
    const diceMessage = {
      type: "roll_dice"
    };

    ws.send(JSON.stringify(diceMessage));

  };

  return (
    <div>
      <button onClick={handleRollDice}>ROLL DICE</button>
    </div>
  );
};

export default Board;