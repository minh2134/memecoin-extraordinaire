import React, { useState } from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Navbar from './components/Navbar';
import Hero from './components/Hero';
import Features from './components/Features';
import TradingStats from './components/TradingStats';
import NotFound from './components/404';
import Developers from './components/Developers';
import Wallet from './components/Wallet';
import About from './components/About'
import Market from './components/Market';
import Trade from './components/Trade';
import './App.css';

function App() {
  const [walletAddress, setWalletAddress] = useState('');
  const [isWalletOpen, setIsWalletOpen] = useState(false);
  const [isDropdownNavbarOpen, setIsDropdownNavbarOpen] = useState(false);
  const [isDropdownTradeFromOpen, setIsDropdownTradeFromOpen] = useState(false);
  const [isDropdownTradeToOpen, setIsDropdownTradeToOpen] = useState(false);

  const handleDisconnect = () => {
    setWalletAddress('');
    setIsDropdownNavbarOpen(false);
    setIsDropdownTradeFromOpen(false);
  };

  return (
    <Router>
      <div className="min-h-screen bg-mystery-dark">
        <Navbar 
          walletAddress={walletAddress} 
          setIsWalletOpen={setIsWalletOpen}
          isDropdownOpen={isDropdownNavbarOpen}
          setIsDropdownOpen={setIsDropdownNavbarOpen}
          handleDisconnect={handleDisconnect}
        />
        
        <Routes>
          <Route path="/" element={
            <>
              <Hero 
                walletAddress={walletAddress} 
                setIsWalletOpen={setIsWalletOpen}
              />
              <div className="bg-gradient-to-b from-[#0B041A] to-mystery-dark">
                <main className="container mx-auto px-4">
                  <Features />
                  <TradingStats />
                </main>
              </div>
            </>
          } />
          <Route path="/404" element={<NotFound />} />
          <Route path="/developers" element={<Developers />} />
	        <Route path="/about" element={<About />} />
          <Route path="/market" element={<Market />} />
          <Route path="/trade" element={
            <Trade 
              isDropdownTradeFromOpen={isDropdownTradeFromOpen}
              setIsDropdownOpen={setIsDropdownTradeFromOpen}
              isDropdownTradeToOpen={isDropdownTradeToOpen}
              setIsDropdownTradeToOpen={setIsDropdownTradeToOpen}
            />} />
          <Route path="*" element={<NotFound />} />
        </Routes>

        <Wallet 
          isOpen={isWalletOpen}
          onClose={() => setIsWalletOpen(false)}
          setWalletAddress={setWalletAddress}
        />
      </div>
    </Router>
  );
}

export default App;
