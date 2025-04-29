import styles from "./Board.module.css";

const Board = () => {
  return (
    <div>
      <h2>Board</h2>
      <div className={styles.square}></div>
    </div>
  );
};

export default Board;