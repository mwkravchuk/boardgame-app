import { useEffect } from "react";
import { useNavigate } from "react-router";
import { useWebSocket } from "../../contexts/WebSocketProvider";
import styles from "./Room.module.css";

const Room = () => {
  const navigate = useNavigate();
  const { sendMessage, addListener, removeListener } = useWebSocket();

  useEffect(() => {

    const handleRoomCreated = (message) => {
      console.log("message: ", message);
      navigate("/lobby", { state: { roomCode: message.data }});
    };

    addListener("room_created", handleRoomCreated);

    return () => {
      removeListener("room_created", handleRoomCreated);
    };
  }, [addListener, removeListener, navigate]);

  const handleCreateRoom = () => {
    sendMessage("create_room", "");
  };

  const handleJoinRoom = () => {
    sendMessage("join_room", "")
  };

  return (
    <div className={styles.roomContainer}>
      <button onClick={handleCreateRoom}>Create Room</button>
      <button onClick={handleJoinRoom}>Join Room</button>  
    </div>
  );
};

export default Room;