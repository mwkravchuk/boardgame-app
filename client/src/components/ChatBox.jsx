import { useState } from "react";

const ChatBox = ({ ws }) => {

  const [message, setMessage] = useState("");

  const handleSendMessage = () => {
    if (message.trim() !== "") {
      const chatMessage = {
        type: "chat",
        data: message,
      };

      ws.send(JSON.stringify(chatMessage));
      setMessage("");

    }
  };

  return (
    <div>
      <input
        type="text"
        value={message}
        onChange={(e) => setMessage(e.target.value)}
        placeholder="Type a message..."
      />
      <button onClick={handleSendMessage}>SEND</button>
    </div>
  );
};

export default ChatBox;