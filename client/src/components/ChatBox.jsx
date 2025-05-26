import { useState, useEffect } from "react";
import { useWebSocket } from "../contexts/WebSocketProvider";

const ChatBox = () => {

  const { sendMessage, addListener, removeListener } = useWebSocket();

  const [messages, setMessages] = useState([]);
  const [messageToSend, setMessageToSend] = useState("");

  // Wrap in useEffect so that it just adds it as a listener
  // once on mount
  useEffect(() => {
    // Define the function
    
    const updateChatMessages = (message) => {
      setMessages((prev) => [...prev, message]);
    };
    // Make it a listener whenever we receive a chat message
    addListener("chat", updateChatMessages);

    // Remove it as listener on unmount
    return () => {
      removeListener("chat", updateChatMessages);
    };
  }, [addListener, removeListener]);

  const handleSendMessage = () => {
    sendMessage("chat", messageToSend);
    setMessageToSend("");
  };

  return (
    <div>
      <div>
        <h2>chat history</h2>
        <ul>
          {messages.map((msg, i) => (
            <li key={i}>{msg.sender} : {msg.data}</li>
          ))}
        </ul>
      </div>
      <div>
        <input
          type="text"
          value={messageToSend}
          onChange={(e) => setMessageToSend(e.target.value)}
          placeholder="Type a message..."
          onKeyDown={(e) => {
            if (e.key === "Enter") {
              handleSendMessage();
            }
          }}
        />
      </div>
    </div>
  );
};

export default ChatBox;