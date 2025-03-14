import React, { useState, useEffect } from 'react';
import { useWallet } from '../contexts/WalletContext';
import binance from '../assets/wallets/binance.png';
import edge from '../assets/wallets/edge.png';
import exodus from '../assets/wallets/exodus.png';
import metamask from '../assets/wallets/metamask.png';
import trust from '../assets/wallets/trust.png';
import zengo from '../assets/wallets/zengo.png';
import qrCode from '../assets/qr-code.png';
import { BsFillWalletFill, BsCurrencyExchange, BsPlusCircle } from 'react-icons/bs';
import web3Service from '../utils/web3Service';

const wallets = [
  { name: 'Binance', icon: binance },
  { name: 'Edge', icon: edge },
  { name: 'Exodus', icon: exodus },
  { name: 'MetaMask', icon: metamask },
  { name: 'Trust', icon: trust },
  { name: 'Zengo', icon: zengo },
];

const Wallet = ({ isOpen, onClose, setWalletAddress }) => {
  const [selectedWallet, setSelectedWallet] = useState(null);
  const [selectedOption, setSelectedOption] = useState('external'); // 'external', 'builtin', or 'dashboard'
  const { createWallet, wallet, balances, loading, error, addFunds, refreshBalances, hasWallet } = useWallet();
  const [newFundAmount, setNewFundAmount] = useState('');
  const [selectedCurrency, setSelectedCurrency] = useState('MEME');
  const [isAddingFunds, setIsAddingFunds] = useState(false);
  
  // Refresh balances when dashboard is selected
  useEffect(() => {
    if (selectedOption === 'dashboard' && hasWallet) {
      refreshBalances();
    }
  }, [selectedOption, hasWallet, refreshBalances]);
  
  if (!isOpen) return null;

  const handleWalletClick = async (walletName) => {
    setSelectedWallet(walletName);
    
    // Only attempt to connect if MetaMask is selected
    if (walletName === 'MetaMask') {
      try {
        // Initialize and connect using web3Service
        await web3Service.initialize();
        const address = await web3Service.connectWallet();
        
        if (address) {
          setWalletAddress(address);
          setSelectedOption('dashboard');
          onClose();
        }
      } catch (error) {
        console.error('Failed to connect wallet:', error);
        alert('Failed to connect to MetaMask. Please make sure it is installed and unlocked.');
      }
    }
  };

  const handleModalClick = (e) => {
    if (e.target === e.currentTarget) {
      onClose();
      setSelectedWallet(null);
    }
  };

  const handleQRCodeClick = async () => {
    if (selectedWallet === 'MetaMask') {
      try {
        const address = await web3Service.connectWallet();
        if (address) {
          setWalletAddress(address);
          onClose();
        }
      } catch (error) {
        console.error('Failed to connect wallet:', error);
        alert('Failed to connect to MetaMask. Please make sure it is installed and unlocked.');
      }
    } else if (selectedWallet) {
      // For other wallets, keep the existing mock behavior or implement similarly
      setWalletAddress('0x...BAa8');
      onClose();
      setSelectedWallet(null);
    } else {
      alert('Please select a wallet first!');
    }
  };
  
  const handleCreateWallet = async () => {
    try {
      const newWallet = await createWallet();
      setWalletAddress(newWallet.address);
      setSelectedOption('dashboard'); // Switch to dashboard after creation
    } catch (err) {
      console.error('Failed to create wallet:', err);
    }
  };
  
  const handleUseExistingWallet = () => {
    if (wallet && wallet.address) {
      setWalletAddress(wallet.address);
      setSelectedOption('dashboard'); // Switch to dashboard after connecting
    }
  };
  
  const handleAddFunds = async (e) => {
    e.preventDefault();
    
    if (!newFundAmount || isNaN(parseFloat(newFundAmount))) {
      return;
    }
    
    try {
      await addFunds(selectedCurrency, parseFloat(newFundAmount));
      setNewFundAmount('');
      setIsAddingFunds(false);
      refreshBalances(); // Refresh balances after adding funds
    } catch (err) {
      console.error('Failed to add funds:', err);
    }
  };
  
  // Convert raw balance to decimal format (backend stores as integer with 2 decimal places)
  const formatBalance = (amount) => {
    return (amount / 100).toFixed(2);
  };

  return (
    <div 
      className="fixed inset-0 bg-black/70 backdrop-blur-sm z-50 flex items-center justify-center"
      onClick={handleModalClick}
    >
      <div className="bg-[#1E1E1E] rounded-2xl p-8 max-w-4xl w-full mx-4">
        <h2 className="font-heading text-3xl text-white mb-4">Wallet Connect</h2>
        
        
        {/* External Wallet Option */}
        {selectedOption === 'external' && (
          <div className="flex gap-8">
            {/* Left side - Wallet selection */}
            <div className="flex-1">
                          
              <div className="grid grid-cols-2 gap-4">
                {wallets.map((wallet) => (
                  <div
                    key={wallet.name}
                    className={`p-4 border-2 border-mystery-accent/20 rounded-xl cursor-pointer transition-all duration-300 hover:border-mystery-accent/60 ${
                      selectedWallet === wallet.name ? 'scale-105 border-mystery-accent' : ''
                    }`}
                    onClick={() => handleWalletClick(wallet.name)}
                  >
                    <img 
                      src={wallet.icon} 
                      alt={wallet.name}
                      className="w-16 h-16 mx-auto mb-2 object-contain"
                    />
                    <p className="text-center text-white font-heading">{wallet.name}</p>
                  </div>
                ))}
              </div>
            </div>

            {/* Right side - QR code */}
            <div className="w-72 bg-[#2A2A2A] p-6 rounded-xl">
              <h3 className="font-heading text-xl text-white mb-6">How to connect to your wallet</h3>
              <img 
                src={qrCode} 
                alt="QR Code"
                className={`w-48 h-48 mx-auto mb-6 cursor-pointer ${selectedWallet ? 'hover:opacity-80' : 'opacity-50'}`}
                onClick={handleQRCodeClick}
              />
              <div className="space-y-2">
                <p className="text-gray-300">1. Select your wallet from the left</p>
                <p className="text-gray-300">2. Click the QR code to connect</p>
              </div>
            </div>
          </div>
        )}

        
        {/* Wallet Dashboard Option */}
        {selectedOption === 'dashboard' && (
          <div className="py-6">
            {loading ? (
              <div className="py-8 text-center text-white">Loading wallet information...</div>
            ) : error ? (
              <div className="py-8 text-center text-red-500">Error: {error}</div>
            ) : (
              <>
                <div className="bg-[#2A2A2A] rounded-xl p-6 mb-6">
                  <div className="flex flex-col md:flex-row items-start md:items-center justify-between mb-6">
                    <div>
                      <h3 className="text-xl font-heading text-white mb-2">Your Wallet</h3>
                      <p className="text-gray-400 font-mono text-sm">{wallet.address}</p>
                    </div>
                    <button 
                      onClick={() => setIsAddingFunds(true)}
                      className="mt-4 md:mt-0 flex items-center gap-2 px-4 py-2 bg-mystery-accent/20 hover:bg-mystery-accent/30 text-mystery-accent rounded-lg transition-colors"
                    >
                      <BsPlusCircle /> Add Funds
                    </button>
                  </div>
                  
                  <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
                    {Object.entries(balances).map(([currency, amount]) => (
                      <div key={currency} className="bg-[#1E1E1E] p-4 rounded-lg">
                        <h4 className="text-gray-400 mb-1 text-sm">Balance</h4>
                        <div className="flex items-baseline gap-2">
                          <span className="text-2xl font-bold text-white">{formatBalance(amount)}</span>
                          <span className="text-mystery-accent">{currency}</span>
                        </div>
                      </div>
                    ))}
                    
                    {Object.keys(balances).length === 0 && (
                      <div className="bg-[#1E1E1E] p-4 rounded-lg col-span-full">
                        <p className="text-gray-400 text-center">No balances available. Add funds to get started!</p>
                      </div>
                    )}
                  </div>
                </div>
                
                <div className="bg-[#2A2A2A] rounded-xl p-6">
                  <div className="flex items-center gap-3 mb-4">
                    <BsCurrencyExchange className="text-xl text-mystery-accent" />
                    <h3 className="text-xl font-heading text-white">Recent Activity</h3>
                  </div>
                  
                  <div className="text-gray-400 text-center py-6">
                    <p>Transaction history will appear here</p>
                    <p className="text-sm mt-2">Coming soon!</p>
                  </div>
                </div>
              </>
            )}
          </div>
        )}
        
        {/* Add Funds Dialog */}
        {isAddingFunds && (
          <div className="fixed inset-0 bg-black/70 backdrop-blur-sm z-[60] flex items-center justify-center">
            <div className="bg-[#1E1E1E] rounded-2xl p-6 max-w-md w-full mx-4">
              <h2 className="text-2xl font-heading text-white mb-4">Add Funds</h2>
              <form onSubmit={handleAddFunds}>
                <div className="mb-4">
                  <label htmlFor="currency" className="block text-gray-400 mb-2">Currency</label>
                  <select
                    id="currency"
                    value={selectedCurrency}
                    onChange={(e) => setSelectedCurrency(e.target.value)}
                    className="w-full bg-[#2A2A2A] text-white border border-gray-700 rounded-lg px-3 py-2"
                  >
                    <option value="MEME">MEME</option>
                    <option value="DOGE">DOGE</option>
                    <option value="PEPE">PEPE</option>
                    <option value="SHIB">SHIB</option>
                  </select>
                </div>
                
                <div className="mb-6">
                  <label htmlFor="amount" className="block text-gray-400 mb-2">Amount</label>
                  <input
                    id="amount"
                    type="number"
                    min="0.01"
                    step="0.01"
                    value={newFundAmount}
                    onChange={(e) => setNewFundAmount(e.target.value)}
                    className="w-full bg-[#2A2A2A] text-white border border-gray-700 rounded-lg px-3 py-2"
                    placeholder="0.00"
                    required
                  />
                </div>
                
                <div className="flex gap-4">
                  <button
                    type="button"
                    onClick={() => setIsAddingFunds(false)}
                    className="flex-1 px-4 py-2 border border-gray-700 text-white rounded-lg hover:bg-gray-800 transition-colors"
                  >
                    Cancel
                  </button>
                  <button
                    type="submit"
                    className="flex-1 px-4 py-2 bg-mystery-accent hover:bg-mystery-accent/80 text-black font-bold rounded-lg transition-colors"
                  >
                    Add Funds
                  </button>
                </div>
              </form>
            </div>
          </div>
        )}
      </div>
    </div>
  );
};

export default Wallet; 