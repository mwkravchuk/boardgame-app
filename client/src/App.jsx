import { useState } from "react";

import Header from "./components/structure/Header";
import ChatHistory from "./components/ChatHistory";
import WebSocketComponent from "./components/websocket/WebSocketComponent";

import './App.css'

function App() {

  const [chatHistory, setChatHistory] = useState([]);

  return (
    <div>
      <Header />
      <ChatHistory chatHistory={chatHistory} />
      <WebSocketComponent setChatHistory={setChatHistory}/> {/* Handles websocket communication */}
    </div>
  )
}

export default App
