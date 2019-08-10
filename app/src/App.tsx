import React from 'react';
import { BrowserRouter } from 'react-router-dom'
import './App.css';
import Routes from './Routes'
import NavDrawer from './components/NavDrawer';

const App: React.FC = () => {
  return (
    <BrowserRouter>
      <NavDrawer>
        <Routes />
      </NavDrawer>
    </BrowserRouter>
  )
}

export default App;
