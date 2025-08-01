import { useState, useEffect, useRef } from "react";
import { useWebSocket } from "../contexts/WebSocketProvider";

import { ScrollArea } from "../components/ui/scroll-area";
import { Separator } from "../components/ui/separator";
import { Input } from "../components/ui/input";

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

  const messagesEndRef = useRef(null);
  useEffect(() => {
    if (messagesEndRef.current) {
      messagesEndRef.current.scrollTop = messagesEndRef.current.scrollHeight;
    }
  }, [messages]);

  const handleSendMessage = () => {
    sendMessage("chat", messageToSend);
    setMessageToSend("");
  };

  return (
    <div>
        <div>
          <ul className="overflow-auto h-64 w-64 pr-2" ref={messagesEndRef}>
            {messages.map((msg, i) => (
              <li key={i}>{msg.sender} : {msg.data}</li>
            ))}
          </ul>
        </div>
      <Input type="text"
              value={messageToSend}
              onChange={(e) => setMessageToSend(e.target.value)}
              placeholder="Type a message..."
              onKeyDown={(e) => {
              if (e.key === "Enter") {
                handleSendMessage();
              }
            }}/>
    </div>
  );
};

export default ChatBox;