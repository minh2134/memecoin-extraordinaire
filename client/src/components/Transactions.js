import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { Decimal } from 'decimal.js';

const Transactions = ({ transactions = [] }) => {
  const navigate = useNavigate();
  const [expandedTransaction, setExpandedTransaction] = useState(null);

  // Function to format addresses by shortening them
  const shortenAddress = (address) => {
    if (!address) return '';
    return `${address.substring(0, 6)}...${address.substring(address.length - 4)}`;
  };

  // Toggle transaction details
  const toggleDetails = (index) => {
    if (expandedTransaction === index) {
      setExpandedTransaction(null);
    } else {
      setExpandedTransaction(index);
    }
  };

  return (
    <div className="min-h-screen bg-gradient-to-b from-[#0B041A] to-mystery-dark py-16">
      <div className="container mx-auto px-4">
        <div className="max-w-4xl mx-auto">
          <div className="flex items-center justify-between mb-8">
            <h2 className="font-heading text-3xl text-white">Transaction History</h2>
            <button 
              onClick={() => navigate('/trade')}
              className="font-heading px-6 py-2 rounded-full bg-mystery-accent text-white hover:opacity-90 transition-opacity"
            >
              New Trade
            </button>
          </div>

          {transactions.length === 0 ? (
            <div className="bg-[#1E1E1E]/50 rounded-xl p-8 text-center">
              <p className="font-heading text-xl text-gray-300">No transactions yet</p>
              <p className="text-gray-400 mt-2">Complete a swap to see your trading history</p>
            </div>
          ) : (
            <div className="space-y-4">
              {transactions.map((tx, index) => (
                <div 
                  key={index}
                  className={`bg-[#1E1E1E]/50 rounded-xl p-6 border 
                    ${tx.status === 'Failed' 
                      ? 'border-red-500/30 hover:border-red-500' 
                      : 'border-mystery-accent/30 hover:border-mystery-accent'}
                    transition-colors duration-300 cursor-pointer`}
                  onClick={() => toggleDetails(index)}
                >
                  <div className="flex justify-between items-start">
                    <div>
                      <div className="flex items-center gap-2">
                        <p className="font-heading text-lg text-white">
                          {tx.fromCurr || tx.tradingToken} → {tx.toCurr || tx.receivingToken}
                        </p>
                        <span className={`text-sm font-heading px-2 py-0.5 rounded-full ${
                          tx.status === 'Failed' 
                            ? 'bg-red-500/20 text-red-400' 
                            : 'bg-green-500/20 text-green-400'
                        }`}>
                          {tx.status || 'Success'}
                        </span>
                      </div>
                      <p className="text-gray-400 text-sm mt-1">{tx.date}</p>
                    </div>
                    <div className="text-right">
                      <p className="font-heading text-lg text-mystery-highlight">
                        {tx.tradedAmount || tx.tradingAmount} {tx.fromCurr || tx.tradingToken}
                      </p>
                      <p className="font-heading text-lg text-mystery-highlight">
                        {tx.receivedAmount || tx.receivingAmount} {tx.toCurr || tx.receivingToken}
                      </p>
                    </div>
                  </div>
                  
                  {/* Expanded transaction details */}
                  {expandedTransaction === index && (
                    <div className="mt-4 pt-4 border-t border-gray-700/50 text-sm">
                      <div className="grid grid-cols-1 md:grid-cols-2 gap-3">
                        <div>
                          <p className="text-gray-400">Rate</p>
                          <p className="text-white">
                            1 {tx.fromCurr || tx.tradingToken} = {' '}
                            {new Decimal(tx.receivedAmount || tx.receivingAmount)
                              .div(new Decimal(tx.tradedAmount || tx.tradingAmount))
                              .toFixed(6)}{' '}
                            {tx.toCurr || tx.receivingToken}
                          </p>
                        </div>
                        
                        {tx.sourceAddress && (
                          <div>
                            <p className="text-gray-400">Your Address</p>
                            <p className="text-white font-mono">{shortenAddress(tx.sourceAddress)}</p>
                          </div>
                        )}
                        
                        {tx.tradedAddress && (
                          <div>
                            <p className="text-gray-400">Counterparty</p>
                            <p className="text-white font-mono">{shortenAddress(tx.tradedAddress)}</p>
                          </div>
                        )}
                        
                        {tx.slippage && (
                          <div>
                            <p className="text-gray-400">Slippage</p>
                            <p className="text-white">{new Decimal(tx.slippage).times(100).toString()}%</p>
                          </div>
                        )}
                      </div>
                      
                      {/* If there's an error message available */}
                      {tx.errorMessage && (
                        <div className="mt-3 p-3 bg-red-500/10 rounded-lg border border-red-500/30">
                          <p className="text-red-400">{tx.errorMessage}</p>
                        </div>
                      )}
                    </div>
                  )}
                </div>
              ))}
            </div>
          )}
        </div>
      </div>
    </div>
  );
};

export default Transactions; 