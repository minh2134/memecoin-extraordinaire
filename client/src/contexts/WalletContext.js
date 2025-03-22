//@author Dinh Le Hoang Anh - 105508318
//@author Pham Vu Minh - 105110564
import React, { createContext, useState, useContext, useEffect } from 'react';
import web3Service from '../utils/web3Service';

const WalletContext = createContext();

export function useWallet() {
  return useContext(WalletContext);
}

export function WalletProvider({ children }) {
  const [wallet, setWallet] = useState(null);
  const [balance, setBalance] = useState(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [tokenBalances, setTokenBalances] = useState({});

  const assignWallet = async () => {
    setLoading(true);
    setError(null);
    
    try {
      const walletInfo = await web3Service.assignWalletToUser();
      console.log("Wallet assigned:", walletInfo);
      
      setWallet(walletInfo.address);
      setBalance(walletInfo.balance);
      
      // Return the full wallet info including address and balance
      return walletInfo;
    } catch (err) {
      console.error("Error assigning wallet:", err);
      setError('Failed to assign wallet: ' + err.message);
      throw err;
    } finally {
      setLoading(false);
    }
  };

  const fetchTokenBalances = async (address) => {
    if (!address) return;
    
    const tokens = ['BTC', 'DAI', 'SNX']; // Only these three coins
    const balances = {};
    
    try {
      for (const token of tokens) {
        const response = await fetch(`http://localhost:8080/account/curr/${token}`);
        if (response.ok) {
          const data = await response.json();
          balances[token] = data.amount;
        } else {
          balances[token] = '0';
        }
      }
      
      setTokenBalances(balances);
    } catch (error) {
      console.error('Error fetching token balances:', error);
    }
  };

  const value = {
    wallet,
    balance,
    loading,
    error,
    assignWallet,
    tokenBalances,
    fetchTokenBalances
  };

  return (
    <WalletContext.Provider value={value}>
      {children}
    </WalletContext.Provider>
  );
} 