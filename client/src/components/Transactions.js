import React from 'react';
import { useNavigate } from 'react-router-dom';

const Transactions = ({ transactions = [] }) => {
  const navigate = useNavigate();

  return (
    <div className="min-h-screen bg-gradient-to-b from-[#0B041A] to-mystery-dark py-16">
      <div className="container mx-auto px-4">
        <div className="max-w-4xl mx-auto">
          <div className="flex items-center justify-between mb-8">
            <h2 className="font-heading text-3xl text-white">Recent Transactions</h2>
            <button 
              onClick={() => navigate('/trade')}
              className="font-heading px-6 py-2 rounded-full bg-mystery-accent text-white hover:opacity-90 transition-opacity"
            >
              Make a Trade
            </button>
          </div>

          {transactions.length === 0 ? (
            <div className="bg-[#1E1E1E]/50 rounded-xl p-8 text-center">
              <p className="font-heading text-xl text-gray-300">No recent transactions</p>
              <p className="text-gray-400 mt-2">Your trading history will appear here</p>
            </div>
          ) : (
            <div className="space-y-4">
              {transactions.map((tx, index) => (
                <div 
                  key={index}
                  className={`bg-[#1E1E1E]/50 rounded-xl p-6 border ${
                    tx.status === 'Failed' 
                      ? 'border-red-500/30 hover:border-red-500' 
                      : 'border-mystery-accent/30 hover:border-mystery-accent'
                  } transition-colors duration-300`}
                >
                  <div className="flex justify-between items-start mb-4">
                    <div>
                      <div className="flex items-center gap-2">
                        <p className="font-heading text-lg text-white">
                          {tx.tradingToken} → {tx.receivingToken}
                        </p>
                        {tx.status === 'Failed' && (
                          <span className="text-sm text-red-500 font-heading">Failed</span>
                        )}
                      </div>
                      <p className="text-gray-400 text-sm">{tx.date}</p>
                    </div>
                    <div className="text-right">
                      <p className="font-heading text-lg text-mystery-highlight">
                        {tx.tradingAmount} {tx.tradingToken}
                      </p>
                      <p className="font-heading text-lg text-mystery-highlight">
                        {tx.receivingAmount} {tx.receivingToken}
                      </p>
                    </div>
                  </div>
                  <div className="text-sm text-gray-400">
                    <p>From: {tx.senderWallet}</p>
                    {tx.receiverWallet && <p>To: {tx.receiverWallet}</p>}
                  </div>
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