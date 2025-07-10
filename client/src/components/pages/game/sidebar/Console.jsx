import { useState, useEffect, useRef } from "react";
import { useWebSocket } from "../../../../contexts/WebSocketProvider";

const Console = () => {
  const { addListener, removeListener } = useWebSocket();
  const [messages, setMessages] = useState([]);

  // Wrap in useEffect so that it just adds it as a listener
  // once on mount
  useEffect(() => {
    // Define the function
    
    const updateConsoleMessages = (message) => {
      setMessages((prev) => [...prev, message]);
    };
    // Make it a listener whenever we receive a chat message
    addListener("console", updateConsoleMessages);

    // Remove it as listener on unmount
    return () => {
      removeListener("console", updateConsoleMessages);
    };
  }, [addListener, removeListener]);

  const messagesEndRef = useRef(null);
  useEffect(() => {
    if (messagesEndRef.current) {
      messagesEndRef.current.scrollTop = messagesEndRef.current.scrollHeight;
    }
  }, [messages]);

  return (
    <div>
      <ul className="overflow-auto h-32 w-64 pr-2" ref={messagesEndRef}>
        {messages.map((msg, i) => (
          <li key={i}>{msg.sender} {msg.data}</li>
        ))}
      </ul>
    </div>
  );
};

export default Console;