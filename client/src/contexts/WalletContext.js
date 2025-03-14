import React, { createContext, useState, useContext, useEffect } from 'react';
import web3Service from '../utils/web3Service';

const WalletContext = createContext();

export function useWallet() {
  return useContext(WalletContext);
}

export function WalletProvider({ children }) {
  const [wallet, setWallet] = useState(null);
  const [balances, setBalances] = useState({});
  const [transactions, setTransactions] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  
  const currencies = ["MEME", "ETH", "BTC", "USDT"];

  // Check if wallet exists
  const hasWallet = !!wallet;

  async function createWallet() {
    setLoading(true);
    setError(null);
    
    try {
      const newWallet = await web3Service.createMockWallet();
      setWallet(newWallet);
      await refreshBalances(newWallet.address);
      return newWallet;
    } catch (err) {
      setError('Failed to create wallet: ' + err.message);
      throw err;
    } finally {
      setLoading(false);
    }
  }

  async function connectWallet() {
    setLoading(true);
    setError(null);
    
    try {
      const address = await web3Service.connectWallet();
      setWallet({ address });
      await refreshBalances(address);
      return address;
    } catch (err) {
      setError('Failed to connect wallet: ' + err.message);
      throw err;
    } finally {
      setLoading(false);
    }
  }

  async function refreshBalances(address = null) {
    if (!address && !wallet) return;
    
    setLoading(true);
    const walletAddress = address || wallet.address;
    
    try {
      const newBalances = {};
      for (const currency of currencies) {
        const balance = await web3Service.getBalance(walletAddress, currency);
        newBalances[currency] = balance;
      }
      setBalances(newBalances);
    } catch (err) {
      setError('Failed to refresh balances: ' + err.message);
    } finally {
      setLoading(false);
    }
  }

  async function addFunds(currency, amount) {
    if (!wallet) throw new Error('No wallet connected');
    
    setLoading(true);
    try {
      await web3Service.mintTokens(wallet.address, currency, amount);
      await refreshBalances();
    } catch (err) {
      setError('Failed to add funds: ' + err.message);
      throw err;
    } finally {
      setLoading(false);
    }
  }

  async function executeTrade(targetAddress, sourceCurr, targetCurr, sourceAmount, targetAmount) {
    if (!wallet) throw new Error('No wallet connected');
    
    setLoading(true);
    try {
      const transaction = await web3Service.executeTrade(
        targetAddress,
        sourceCurr,
        targetCurr,
        sourceAmount,
        targetAmount
      );
      
      // Add transaction to list
      setTransactions(prev => [transaction, ...prev]);
      
      // Refresh balances after trade
      await refreshBalances();
      
      return transaction;
    } catch (err) {
      setError('Trade execution failed: ' + err.message);
      throw err;
    } finally {
      setLoading(false);
    }
  }

  const value = {
    wallet,
    balances,
    transactions,
    loading,
    error,
    hasWallet,
    currencies,
    createWallet,
    connectWallet,
    refreshBalances,
    addFunds,
    executeTrade
  };

  return (
    <WalletContext.Provider value={value}>
      {children}
    </WalletContext.Provider>
  );
} 