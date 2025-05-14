import styles from "./Header.module.css";

import { usePlayer } from "../../contexts/PlayerProvider";

const Header = () => {

  const { playerId, currentTurnId } = usePlayer();

  return (
    <header className={styles.header}>
      <h2>Catanopoly</h2>
      <div>playerid is {playerId}</div>
      <div>current turn id is {currentTurnId}</div>
    </header>
  )

};

export default Header;