import { useEffect, useRef } from "react";
import DiceBox from '@3d-dice/dice-box-threejs';

export default function Dice({ values, setAnimationCompleted }) {
  const boxRef = useRef(null);

  useEffect(() => {
    const init = async () => {
      const diceBox = new DiceBox('#dice-box', {
        assetPath: '/assets/',
        scale: 5,
        gravity: 9.8,
        startingHeight: 4,
        onRollComplete: () => {
          setAnimationCompleted(true);
        }
      });
      await diceBox.initialize();

      boxRef.current = diceBox;
    };
    init();
  }, [setAnimationCompleted]);

  useEffect(() => {
    if (boxRef.current && values?.length) {
      setAnimationCompleted(false);
      boxRef.current.roll(`2d6@${values[0]},${values[1]}`);
    };
  }, [values, setAnimationCompleted]);

  return (
    <div
      id="dice-box"
      style={{
        width: '500px',
        height: '500px',
        position: 'relative',
      }}>
    </div>
  );
};