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
import Transactions from './components/Transactions';
import './App.css';

function App() {
  const [walletAddress, setWalletAddress] = useState('');
  const [isWalletOpen, setIsWalletOpen] = useState(false);
  const [isDropdownNavbarOpen, setIsDropdownNavbarOpen] = useState(false);
  const [isDropdownTradeFromOpen, setIsDropdownTradeFromOpen] = useState(false);
  const [isDropdownTradeToOpen, setIsDropdownTradeToOpen] = useState(false);
  const [transactions, setTransactions] = useState([]);

  const handleDisconnect = () => {
    setWalletAddress('');
    setIsDropdownNavbarOpen(false);
    setIsDropdownTradeFromOpen(false);
  };

  const generateRandomWallet = () => {
    const chars = '0123456789abcdef';
    let wallet = '0x';
    for (let i = 0; i < 40; i++) {
      wallet += chars[Math.floor(Math.random() * chars.length)];
    }
    return wallet;
  };

  const handleNewTransaction = (fromToken, fromAmount, toToken, toAmount, failed = false) => {
    const newTransaction = {
      senderWallet: walletAddress,
      receiverWallet: failed ? null : generateRandomWallet(),
      date: new Date().toLocaleString(),
      tradingToken: fromToken,
      tradingAmount: fromAmount,
      receivingToken: toToken,
      receivingAmount: toAmount,
      status: failed ? 'Failed' : 'Success'
    };
    setTransactions([newTransaction, ...transactions]);
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
          <Route path="/transactions" element={<Transactions transactions={transactions} />} />
          <Route path="/trade" element={
            <Trade 
              isDropdownTradeFromOpen={isDropdownTradeFromOpen}
              setIsDropdownTradeFromOpen={setIsDropdownTradeFromOpen}
              isDropdownTradeToOpen={isDropdownTradeToOpen}
              setIsDropdownTradeToOpen={setIsDropdownTradeToOpen}
              onTransaction={handleNewTransaction}
            />
          } />
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
