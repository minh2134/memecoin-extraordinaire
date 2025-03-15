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

  const assignWallet = async () => {
    setLoading(true);
    setError(null);
    
    try {
      const { address, balance } = await web3Service.assignWalletToUser();
      setWallet(address);
      setBalance(balance);
      return address;
    } catch (err) {
      setError('Failed to assign wallet: ' + err.message);
      throw err;
    } finally {
      setLoading(false);
    }
  };

  const value = {
    wallet,
    balance,
    loading,
    error,
    assignWallet
  };

  return (
    <WalletContext.Provider value={value}>
      {children}
    </WalletContext.Provider>
  );
} 