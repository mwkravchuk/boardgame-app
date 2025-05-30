import { useState, useEffect } from "react";
import { useNavigate } from "react-router";
import { useWebSocket } from "../../contexts/WebSocketProvider";

import { Input } from "../ui/input";
import { Button } from "../ui/button";

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
      <Button variant="outline" onClick={handleCreateRoom}>CREATE ROOM</Button>
      {/* Form to join a room */}
      <form className="flex gap-2" onSubmit={handleJoinRoom}>
        <Input
          name="join"
          value={joinCode}
          onChange={(e) => setJoinCode(e.target.value)}
          placeholder="Enter room code"
          onKeyDown={(e) => {
            if (e.key === "Enter") {
              handleJoinRoom();
            }
          }}/>
        <Button variant="outline" type="submit">JOIN ROOM</Button>
      </form>
    </div>
  );
};

export default Room;