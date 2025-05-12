import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.jsx'

import { WebSocketProvider } from "./contexts/WebSocketProvider.jsx";
import { PlayerProvider } from "./contexts/PlayerProvider.jsx";

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <WebSocketProvider>
      <PlayerProvider>
        <App />
      </PlayerProvider>
    </WebSocketProvider>
  </StrictMode>
)
