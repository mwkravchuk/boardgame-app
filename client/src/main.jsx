import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.jsx'

import { WebSocketProvider } from "./contexts/WebSocketProvider.jsx";
import { GameProvider } from "./contexts/GameProvider.jsx";
import { PlayerProvider } from "./contexts/PlayerProvider.jsx";

createRoot(document.getElementById('root')).render(
  // <StrictMode> JUST FOR DEVELOPMENT
    <WebSocketProvider>
      <GameProvider>
        <PlayerProvider>
          <App />
        </PlayerProvider>
      </GameProvider>
    </WebSocketProvider>
  // </StrictMode>
)
