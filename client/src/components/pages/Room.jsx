import { useState, useEffect } from "react";
import { useNavigate } from "react-router";
import { useWebSocket } from "../../contexts/WebSocketProvider";
import styles from "./Room.module.css";

const Room = () => {
  const navigate = useNavigate();
  const { sendMessage, addListener, removeListener } = useWebSocket();

  const [joinCode, setJoinCode] = useState("");

  useEffect(() => {

    const handleRoomJoined = (message) => {
      console.log("message: ", message);
      navigate("/lobby", { state: { roomCode: message.data }});
    };

    addListener("room_joined", handleRoomJoined);

    return () => {
      removeListener("room_", handleRoomJoined);
    };
  }, [addListener, removeListener, navigate]);

  const handleCreateRoom = () => {
    sendMessage("create_room", "");
  };

  const handleJoinRoom = (e) => {
    e.preventDefault();
    if (joinCode.trim()) {
      sendMessage("join_room", joinCode.trim());
    }
  };

  return (
    <div className={styles.roomContainer}>
      <button onClick={handleCreateRoom}>Create Room</button>

      {/* Form to join a room */}
      <form onSubmit={handleJoinRoom}>
        <input
          name="join"
          value={joinCode}
          onChange={(e) => setJoinCode(e.target.value)}
          placeholder="Enter room code"/>
        <button type="submit">Join Room</button>  
      </form>
    </div>
  );
};

export default Room;