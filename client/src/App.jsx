import { useState } from "react";

import Header from "./components/structure/Header";
import ChatBox from "./components/ChatBox";
import Board from "./components/Board";
import { WebSocketProvider } from "./contexts/WebSocketProvider";

import './App.css'

function App() {

  return (
    <div>
      <WebSocketProvider>
        <Header />
        <ChatBox />
        <Board />
      </WebSocketProvider>
    </div>
  )
}

export default App
