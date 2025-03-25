import WebSocketComponent from "./websocket/WebSocketComponent";

import './App.css'

function App() {
  return (
    <div>
      <h1>Board Game</h1>
      <WebSocketComponent /> {/* Handles websocket communication */}
    </div>
  )
}

export default App
