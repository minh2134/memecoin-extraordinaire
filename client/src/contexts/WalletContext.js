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