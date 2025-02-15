import React, { useState } from 'react';
import Navbar from './components/Navbar';
import Hero from './components/Hero';
import Features from './components/Features';
import TradingStats from './components/TradingStats';
import Wallet from './components/Wallet';
import './App.css';

function App() {
  const [walletAddress, setWalletAddress] = useState('');
  const [isWalletOpen, setIsWalletOpen] = useState(false);
  const [isDropdownOpen, setIsDropdownOpen] = useState(false);

  const handleDisconnect = () => {
    setWalletAddress('');
    setIsDropdownOpen(false);
  };

  return (
    <div className="min-h-screen bg-mystery-dark">
      <Navbar 
        walletAddress={walletAddress} 
        setIsWalletOpen={setIsWalletOpen}
        isDropdownOpen={isDropdownOpen}
        setIsDropdownOpen={setIsDropdownOpen}
        handleDisconnect={handleDisconnect}
      />
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

      <Wallet 
        isOpen={isWalletOpen}
        onClose={() => setIsWalletOpen(false)}
        setWalletAddress={setWalletAddress}
      />
    </div>
  );
}

export default App;
