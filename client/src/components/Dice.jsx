import { useEffect, useRef } from "react";
import DiceBox from '@3d-dice/dice-box';

export default function Dice({ values }) {
  const boxRef = useRef(null);

  useEffect(() => {
    const init = async () => {
      const diceBox = new DiceBox('#dice-box', {
        assetPath: '/assets/',
        scale: 5,
        gravity: 9.8,
        startingHeight: 8,
      });
      await diceBox.init();

      boxRef.current = diceBox;

      if (values && values.length) {
        diceBox.roll("2d6");
      }
    };
    init();
  }, [values]);

  useEffect(() => {
    if (boxRef.current && values?.length) {
      boxRef.current.roll(values.map(v => `d6-${v}`));
    };
  }, [values]);

  return (
    <div
      id="dice-box"
      style={{
        width: '300px',
        height: '300px',
        border: `1px solid #ccc`,
        position: 'relative',
      }}>
    </div>
  );
};