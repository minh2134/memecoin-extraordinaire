import React, { useState } from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { WalletProvider } from './contexts/WalletContext';
import Navbar from './components/Navbar';
import Hero from './components/Hero';
import Features from './components/Features';
import TradingStats from './components/TradingStats';
import NotFound from './components/404';
import Developers from './components/Developers';
import About from './components/About';
import Market from './components/Market';
import Trade from './components/Trade';
import Transactions from './components/Transactions';
import Login from './components/Login';
import Wallet from './components/Wallet';
import './App.css';

function App() {
  const [isDropdownNavbarOpen, setIsDropdownNavbarOpen] = useState(false);
  const [isDropdownTradeFromOpen, setIsDropdownTradeFromOpen] = useState(false);
  const [isDropdownTradeToOpen, setIsDropdownTradeToOpen] = useState(false);
  const [transactions, setTransactions] = useState([]);
  const [isWalletOpen, setIsWalletOpen] = useState(false);

  const handleDisconnect = () => {
    // Clear the wallet from context instead of local state
    window.location.reload(); // Simple way to reset the app state
  };

  return (
    <WalletProvider>
      <Router>
        <div className="min-h-screen bg-mystery-dark">
          <Navbar 
            setIsDropdownOpen={setIsDropdownNavbarOpen}
            isDropdownOpen={isDropdownNavbarOpen}
            handleDisconnect={handleDisconnect}
            setIsWalletOpen={setIsWalletOpen}
          />
          
          {isWalletOpen && (
            <Wallet 
              isOpen={isWalletOpen}
              onClose={() => setIsWalletOpen(false)}
            />
          )}

          <Routes>
            <Route path="/login" element={<Login />} />
            <Route path="/" element={<Hero setIsWalletOpen={setIsWalletOpen} />} />
            <Route path="/404" element={<NotFound />} />
            <Route path="/developers" element={<Developers />} />
            <Route path="/about" element={<About />} />
            <Route path="/market" element={<Market />} />
            <Route path="/transactions" element={<Transactions transactions={transactions} />} />
            <Route path="/trade" element={
              <Trade 
                isDropdownTradeFromOpen={isDropdownTradeFromOpen}
                setIsDropdownTradeFromOpen={setIsDropdownTradeFromOpen}
                isDropdownTradeToOpen={isDropdownTradeToOpen}
                setIsDropdownTradeToOpen={setIsDropdownTradeToOpen}
              />
            } />
            <Route path="*" element={<NotFound />} />
          </Routes>
        </div>
      </Router>
    </WalletProvider>
  );
}

export default App;
