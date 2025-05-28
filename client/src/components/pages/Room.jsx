import { useState, useEffect } from "react";
import { useNavigate } from "react-router";
import { useWebSocket } from "../../contexts/WebSocketProvider";

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
    <div className="flex flex-col justify-self-center gap-4 bg-blue-300 h-full">
      <button className="btn-primary" onClick={handleCreateRoom}>Create Room</button>

      {/* Form to join a room */}
      <form className="bg-blue-400" onSubmit={handleJoinRoom}>
        <input
          className="px-4"
          name="join"
          value={joinCode}
          onChange={(e) => setJoinCode(e.target.value)}
          placeholder="Enter room code"/>
        <button className="btn-primary" type="submit">Join Room</button>  
      </form>
    </div>
  );
};

export default Room;