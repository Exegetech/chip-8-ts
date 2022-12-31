import React from 'react'
import ReactDOM from 'react-dom/client'

import { Display } from './devices/display'
import { Sound } from './devices/sound'
import { Keyboard } from './devices/keyboard'
import App from './App'
import './index.css'

global.Display = Display
global.Sound = Sound
global.Keyboard = Keyboard

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
)
