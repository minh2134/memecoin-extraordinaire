import React, { useState, useEffect } from 'react';
import { useWallet } from '../contexts/WalletContext';
import { BsFillWalletFill, BsCurrencyExchange, BsPlusCircle } from 'react-icons/bs';
import web3Service from '../utils/web3Service';

const WalletDashboard = () => {
  const { wallet, balances, loading, error, createWallet, addFunds, refreshBalances, hasWallet } = useWallet();
  const [newFundAmount, setNewFundAmount] = useState('');
  const [selectedCurrency, setSelectedCurrency] = useState('MEME');
  const [isAddingFunds, setIsAddingFunds] = useState(false);
  
  // Refresh balances on component mount
  useEffect(() => {
    if (hasWallet) {
      refreshBalances();
    }
  }, [hasWallet, refreshBalances]);
  
  const handleCreateWallet = async () => {
    try {
      await createWallet();
    } catch (err) {
      console.error('Failed to create wallet:', err);
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
    } catch (err) {
      console.error('Failed to add funds:', err);
    }
  };
  
  // Convert raw balance to decimal format (backend stores as integer with 2 decimal places)
  const formatBalance = (amount) => {
    return (amount / 100).toFixed(2);
  };
  
  if (loading) {
    return <div className="py-12 text-center text-white">Loading wallet information...</div>;
  }
  
  if (error) {
    return <div className="py-12 text-center text-red-500">Error: {error}</div>;
  }
  
  if (!hasWallet) {
    return (
      <div className="py-16 max-w-md mx-auto text-center">
        <BsFillWalletFill className="mx-auto text-6xl text-mystery-accent mb-4" />
        <h2 className="text-3xl font-heading text-white mb-4">No Wallet Found</h2>
        <p className="text-gray-300 mb-6">Create a wallet to store and trade your memecoins!</p>
        <button 
          onClick={handleCreateWallet}
          className="px-6 py-3 bg-mystery-accent hover:bg-mystery-accent/80 text-black font-bold rounded-lg transition-colors"
        >
          Create Wallet
        </button>
      </div>
    );
  }
  
  return (
    <div className="max-w-4xl mx-auto py-8 px-4">
      <div className="bg-[#1E1E1E] rounded-2xl p-8 mb-8">
        <div className="flex flex-col md:flex-row items-start md:items-center justify-between mb-6">
          <div>
            <h2 className="text-3xl font-heading text-white mb-2">Your Wallet</h2>
            <p className="text-gray-400 font-mono text-sm">{wallet.address}</p>
          </div>
          <button 
            onClick={() => setIsAddingFunds(true)}
            className="mt-4 md:mt-0 flex items-center gap-2 px-4 py-2 bg-mystery-accent/20 hover:bg-mystery-accent/30 text-mystery-accent rounded-lg transition-colors"
          >
            <BsPlusCircle /> Add Funds
          </button>
        </div>
        
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
          {Object.entries(balances).map(([currency, amount]) => (
            <div key={currency} className="bg-[#2A2A2A] p-4 rounded-xl">
              <h3 className="text-gray-400 mb-1 text-sm">Balance</h3>
              <div className="flex items-baseline gap-2">
                <span className="text-2xl font-bold text-white">{formatBalance(amount)}</span>
                <span className="text-mystery-accent">{currency}</span>
              </div>
            </div>
          ))}
          
          {Object.keys(balances).length === 0 && (
            <div className="bg-[#2A2A2A] p-4 rounded-xl col-span-full">
              <p className="text-gray-400 text-center">No balances available. Add funds to get started!</p>
            </div>
          )}
        </div>
      </div>
      
      {/* Add Funds Dialog */}
      {isAddingFunds && (
        <div className="fixed inset-0 bg-black/70 backdrop-blur-sm z-50 flex items-center justify-center">
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
      
      {/* Recent Transactions Section */}
      <div className="bg-[#1E1E1E] rounded-2xl p-8">
        <div className="flex items-center gap-3 mb-6">
          <BsCurrencyExchange className="text-2xl text-mystery-accent" />
          <h2 className="text-2xl font-heading text-white">Recent Transactions</h2>
        </div>
        
        <div className="text-gray-400 text-center py-8">
          <p>Transaction history will appear here</p>
          <p className="text-sm mt-2">Coming soon!</p>
        </div>
      </div>
    </div>
  );
};

export default WalletDashboard; 