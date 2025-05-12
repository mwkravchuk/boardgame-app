import styles from "./Header.module.css";

import { usePlayer } from "../../contexts/PlayerProvider";

const Header = () => {

  const { playerId } = usePlayer();

  return (
    <header className={styles.header}>
      <h2>Catanopoly</h2>
      <div>playerid is {playerId}</div>
    </header>
  )

};

export default Header;