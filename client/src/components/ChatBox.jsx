import { useState, useEffect } from "react";
import { useWebSocket } from "../contexts/WebSocketProvider";

import styles from "./ChatBox.module.css";

const ChatBox = () => {

  const { sendMessage, addListener, removeListener } = useWebSocket();

  const [messages, setMessages] = useState([]);
  const [messageToSend, setMessageToSend] = useState("");

  // Wrap in useEffect so that it just adds it as a listener
  // once on mount
  useEffect(() => {
    // Define the function
    
    const updateChatMessages = (message) => {
      console.log(message)
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
    <div className={styles.chatBox}>
      <div className={styles.chatHistory}>
        <h2>chat history</h2>
        <ul className={styles.messagesList}>
          {messages.map((msg, i) => (
            <li key={i}>{msg.sender} : {msg.data}</li>
          ))}
          {console.log(messages)}
        </ul>
      </div>
      <div className={styles.chatInput}>
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