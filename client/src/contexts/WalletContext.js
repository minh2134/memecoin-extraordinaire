import React, { createContext, useState, useContext, useEffect } from 'react';
import { createWallet, getBalance, addFunds } from '../services/walletService';

// Create context
const WalletContext = createContext();

// Hook for using the wallet context
export const useWallet = () => useContext(WalletContext);

// Wallet provider component
export const WalletProvider = ({ children }) => {
  const [wallet, setWallet] = useState(null);
  const [balances, setBalances] = useState({});
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  // Initialize wallet from localStorage on component mount
  useEffect(() => {
    const storedWallet = localStorage.getItem('wallet');
    if (storedWallet) {
      const parsedWallet = JSON.parse(storedWallet);
      setWallet(parsedWallet);
      
      // Fetch current balances
      fetchBalances(parsedWallet.address);
    }
  }, []);

  // Fetch wallet balances
  const fetchBalances = async (address) => {
    if (!address) return;
    
    setLoading(true);
    setError(null);
    
    try {
      const response = await getBalance(address);
      setBalances(response.balances || {});
    } catch (err) {
      setError('Failed to fetch balances');
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  // Create a new wallet
  const createNewWallet = async (name = "Default User") => {
    setLoading(true);
    setError(null);
    
    try {
      const newWallet = await createWallet(name);
      setWallet(newWallet);
      
      // Store wallet in localStorage
      localStorage.setItem('wallet', JSON.stringify(newWallet));
      
      // Initialize with empty balances
      setBalances({});
      
      return newWallet;
    } catch (err) {
      setError('Failed to create wallet');
      console.error(err);
      throw err;
    } finally {
      setLoading(false);
    }
  };

  // Add funds to wallet (for testing)
  const addFundsToWallet = async (currency, amount) => {
    if (!wallet || !wallet.address) {
      setError('No wallet found');
      return;
    }
    
    setLoading(true);
    setError(null);
    
    try {
      await addFunds(wallet.address, currency, amount);
      
      // Update balances after adding funds
      await fetchBalances(wallet.address);
      
      return true;
    } catch (err) {
      setError('Failed to add funds');
      console.error(err);
      throw err;
    } finally {
      setLoading(false);
    }
  };

  // Context value
  const value = {
    wallet,
    balances,
    loading,
    error,
    createWallet: createNewWallet,
    addFunds: addFundsToWallet,
    refreshBalances: () => wallet && fetchBalances(wallet.address),
    hasWallet: !!wallet,
  };

  return (
    <WalletContext.Provider value={value}>
      {children}
    </WalletContext.Provider>
  );
}; 