import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import './index.css'

ReactDOM.createRoot(document.getElementById('root')!).render(
  //useEffectが2回呼ばれないように一時的にStrictMode解除
  // <React.StrictMode>
    <App />
  // </React.StrictMode>,
)
