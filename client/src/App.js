import React from 'react';
import Navbar from './components/Navbar';
import Hero from './components/Hero';
import Features from './components/Features';
import TradingStats from './components/TradingStats';
import './App.css';

function App() {
  return (
    <div className="min-h-screen bg-mystery-dark">
      <Navbar />
      <Hero />
      <div className="bg-gradient-to-b from-[#0B041A] to-mystery-dark">
        <main className="container mx-auto px-4">
          <Features />
          <TradingStats />
        </main>
      </div>
    </div>
  );
}

export default App;
