import { useNavigate } from "react-router";
import { useLocation } from "react-router-dom";

const Lobby = () => {

  const location = useLocation();
  const roomCode = location.state.roomCode;

  const navigate = useNavigate();

  const handleStartGame = () => {
    navigate("/game");
  };

  return (
    <div>
      <h2>GAME LOBBY</h2>
      <p>ROOM CODE IS: {roomCode}</p>
      <button onClick={handleStartGame}>START GAME</button>
    </div>
  );
};

export default Lobby;