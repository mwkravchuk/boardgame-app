import { useState, useEffect } from "react";
import { useNavigate } from "react-router";
import { useWebSocket } from "../../contexts/WebSocketProvider";

import { Input } from "../ui/input";
import { Button } from "../ui/button";
import { Alert, AlertTitle, AlertDescription } from "../ui/alert";
import { AlertCircleIcon } from "lucide-react"

const Room = () => {
  const navigate = useNavigate();
  const { sendMessage, addListener, removeListener } = useWebSocket();

  const [joinCode, setJoinCode] = useState("");
  const [displayName, setDisplayName] = useState("");
  const [error, setError] = useState("");

  useEffect(() => {

    const handleRoomJoined = (message) => {
      setError(false);
      navigate("/lobby", { state: { roomCode: message.data }});
    };

    const handleRoomJoinedFail = (message) => {
      console.log("join fail: ", message);
      setError(true);
    };

    addListener("room_joined", handleRoomJoined);
    addListener("room_joined_fail", handleRoomJoinedFail);

    return () => {
      removeListener("room_joined", handleRoomJoined);
      removeListener("room_joined_fail", handleRoomJoinedFail);
    };
  }, [addListener, removeListener, navigate]);

  const handleCreateRoom = () => {
    sendMessage("create_room", { displayName });
  };

  const handleJoinRoom = (e) => {
    e.preventDefault();
    if (joinCode.trim()) {
      sendMessage("join_room", { code: joinCode.trim(), displayName });
    }
  };

  return (
    <div className="flex flex-col justify-self-center self-center gap-4 h-full">
      {/* Enter username */}
      <div>
        <h2 className="text-2xl">NAME</h2>
        <Input
          name="displayName"
          value={displayName}
          onChange={(e) => setDisplayName(e.target.value)}
          placeholder="Enter name"
        />
      </div>
      {/* Create a room */}
      <div className="flex flex-col">
        <h2 className="text-2xl">HOST</h2>
        <Button className="py-5 px-8" onClick={handleCreateRoom}>CREATE ROOM</Button>
      </div>
      {/* Join a room */}
      <div className="flex flex-col">
        <h2 className="text-2xl">PRIVATE</h2>
        <form className="flex" onSubmit={handleJoinRoom}>
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
          <Button className="px-3" type="submit">JOIN ROOM</Button>
        </form>
        {error && (
          <Alert variant="destructive" className="mt-1">
            <AlertCircleIcon />
            <AlertTitle>Unable to join room.</AlertTitle>
            <AlertDescription>
              <p>Please verify room code and try again.</p>
            </AlertDescription>
          </Alert>
        )}
      </div>
    </div>
  );
};

export default Room;