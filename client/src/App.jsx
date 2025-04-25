import Header from "./components/structure/Header";
import ChatBox from "./components/ChatBox";
import { WebSocketProvider } from "./contexts/WebSocketProvider";

import './App.css'

function App() {

  return (
    <div>
      <WebSocketProvider>
        <Header />
        <ChatBox />
      </WebSocketProvider>
    </div>
  )
}

export default App
