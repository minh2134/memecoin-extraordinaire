//@author Dinh Le Hoang Anh - 105508318
//@author Pham Vu Minh - 105110564
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

const Wallet = ({ isOpen, onClose }) => {
  const [selectedWallet, setSelectedWallet] = useState(null);
  const [isLoading, setIsLoading] = useState(false);
  const { assignWallet } = useWallet();

  if (!isOpen) return null;

  const handleWalletClick = (walletName) => {
    setSelectedWallet(walletName);
  };

  const handleQRCodeClick = async () => {
    if (!selectedWallet) {
      alert('Please select a wallet first!');
      return;
    }

    try {
      setIsLoading(true);
      console.log(`Connecting ${selectedWallet} wallet...`);
      
      // Get a random prefunded wallet from the backend
      const walletInfo = await assignWallet();
      console.log('Kurtosis wallet assigned:', walletInfo);
      
      alert(`Connected to wallet: ${walletInfo.address}\nBalance: ${walletInfo.balance} ETH`);
      onClose();
    } catch (error) {
      console.error('Failed to connect wallet:', error);
      alert(`Failed to connect wallet: ${error.message}`);
    } finally {
      setIsLoading(false);
    }
  };

  const handleModalClick = (e) => {
    if (e.target === e.currentTarget) {
      onClose();
      setSelectedWallet(null);
    }
  };

  return (
    <div 
      className="fixed inset-0 bg-black/70 backdrop-blur-sm z-50 flex items-center justify-center"
      onClick={handleModalClick}
    >
      <div className="bg-[#1E1E1E] rounded-2xl p-8 max-w-4xl w-full mx-4">
        <h2 className="font-heading text-3xl text-white mb-4">Wallet Connect</h2>
        
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
            {isLoading ? (
              <div className="w-48 h-48 mx-auto mb-6 flex items-center justify-center">
                <p className="text-white">Connecting to Kurtosis...</p>
              </div>
            ) : (
              <img 
                src={qrCode} 
                alt="QR Code"
                className={`w-48 h-48 mx-auto mb-6 cursor-pointer ${selectedWallet ? 'hover:opacity-80' : 'opacity-50'}`}
                onClick={handleQRCodeClick}
              />
            )}
            <div className="space-y-2">
              <p className="text-gray-300">1. Select your wallet from the left</p>
              <p className="text-gray-300">2. Click the QR code to connect</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Wallet; 