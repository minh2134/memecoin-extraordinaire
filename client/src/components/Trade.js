import React, { useState } from 'react'
import tradeIcon from '../assets/trade-icon.png'

const Trade = ({ 
  isDropdownTradeFromOpen, 
  setIsDropdownTradeFromOpen,
  isDropdownTradeToOpen,
  setIsDropdownTradeToOpen,
  onTransaction
}) => {
  const [fromState, setFromState] = useState('BTC');
  const [toState, setToState] = useState('DOGE');
  const [fromAmount, setFromAmount] = useState('');
  const [toAmount, setToAmount] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [transactionStatus, setTransactionStatus] = useState(null); // 'success' or 'failed'
  
  const tokens = [
    { name: 'Bitcoin', symbol: 'BTC' },
    { name: 'Dogecoin', symbol: 'DOGE' },
    { name: 'Shiba Inu', symbol: 'SHIB' },
    { name: 'Pepe', symbol: 'PEPE' },
    { name: 'Bonk', symbol: 'BONK' }
  ];

  const buttonClasses = `font-heading px-6 py-3 w-full rounded-[30px] bg-gradient-to-r 
    ${(!fromAmount || !toAmount) 
      ? 'from-gray-500 to-gray-600 cursor-not-allowed' 
      : 'from-[#503BEE] to-[#8A2BE2] hover:from-[#8A2BE2] hover:to-[#503BEE]'} 
    text-white transition-all duration-300 shadow-lg hover:shadow-mystery-accent/50 flex items-center justify-center`;

  const handleFromTokenSelect = (symbol) => {
    if (symbol === toState) {
      setToState(fromState);
    }
    setFromState(symbol);
    setIsDropdownTradeFromOpen(false);
  };

  const handleToTokenSelect = (symbol) => {
    if (symbol === fromState) {
      setFromState(toState);
    }
    setToState(symbol);
    setIsDropdownTradeToOpen(false);
  };

  const handleSwapPositions = () => {
    const tempFrom = fromState;
    const tempFromAmount = fromAmount;
    setFromState(toState);
    setFromAmount(toAmount);
    setToState(tempFrom);
    setToAmount(tempFromAmount);
  };

  const handleAmountChange = (e, setter) => {
    const value = e.target.value;
    if (value === '' || (Number(value) >= 0 && !value.includes('e'))) {
      setter(value);
    }
  };

  const closeModal = () => {
    setIsLoading(false);
    setTransactionStatus(null);
  };

  const handleSwap = async () => {
    if (!fromAmount || !toAmount) return;
    
    setIsLoading(true);
    setTransactionStatus(null);
    
    await new Promise(resolve => setTimeout(resolve, 2000));
    
    const isSuccess = Math.random() >= 0.5;
    
    if (isSuccess) {
      onTransaction(fromState, fromAmount, toState, toAmount);
      setTransactionStatus('success');
    } else {
      onTransaction(fromState, fromAmount, toState, toAmount, true);
      setTransactionStatus('failed');
    }
    
    setFromAmount('');
    setToAmount('');
  };

  return (
    <div className="min-h-screen bg-gradient-to-b from-[#0B041A] to-mystery-dark">
      <div className="container mx-auto px-4 py-24">
        <div className="flex items-center justify-center">
          <div className="relative w-full max-w-[509px] bg-[#1E1E1E] rounded-[20px] p-6 space-y-6">
            {/* From section */}
            <div className="w-full">
              <div className="text-gray-400 mb-2">From</div>
              <div className="flex flex-row w-full rounded-[30px] bg-[#372F47] justify-between items-center p-4">
                <div className="flex items-center gap-4">
                  <div
                    className="text-2xl font-heading cursor-pointer hover:text-mystery-accent transition-colors"
                    onClick={() => setIsDropdownTradeFromOpen(true)}
                  > 
                    {fromState}
                  </div>
                  {isDropdownTradeFromOpen && (
                    <div 
                      className="absolute left-0 mt-2 w-48 bg-[#1E1E1E] border border-mystery-accent/30 rounded-xl shadow-xl z-10"
                      onMouseLeave={() => setIsDropdownTradeFromOpen(false)}
                    >
                      {tokens.map(token => (
                        <div 
                          key={token.symbol}
                          className="px-4 py-3 text-white hover:bg-mystery-accent/20 cursor-pointer transition-colors"
                          onClick={() => handleFromTokenSelect(token.symbol)}
                        >
                          {token.symbol}
                        </div>
                      ))}
                    </div>
                  )}
                </div>
                <input
                  type="number"
                  value={fromAmount}
                  onChange={(e) => handleAmountChange(e, setFromAmount)}
                  placeholder="0.00"
                  className="text-2xl font-heading bg-transparent text-right w-32 focus:outline-none placeholder-gray-500 [appearance:textfield] [&::-webkit-outer-spin-button]:appearance-none [&::-webkit-inner-spin-button]:appearance-none"
                />
              </div>
            </div>

            {/* Trade Icon */}
            <div className="flex justify-center -my-2">
              <img 
                src={tradeIcon} 
                alt="Trade" 
                className="w-12 h-12 cursor-pointer hover:scale-110 transition-transform"
                onClick={handleSwapPositions}
              />
            </div>

            {/* To section */}
            <div className="w-full">
              <div className="text-gray-400 mb-2">To</div>
              <div className="flex flex-row w-full rounded-[30px] bg-[#372F47] justify-between items-center p-4">
                <div className="flex items-center gap-4">
                  <div
                    className="text-2xl font-heading cursor-pointer hover:text-mystery-accent transition-colors"
                    onClick={() => setIsDropdownTradeToOpen(true)}
                  > 
                    {toState}
                  </div>
                  {isDropdownTradeToOpen && (
                    <div 
                      className="absolute left-0 mt-2 w-48 bg-[#1E1E1E] border border-mystery-accent/30 rounded-xl shadow-xl z-10"
                      onMouseLeave={() => setIsDropdownTradeToOpen(false)}
                    >
                      {tokens.map(token => (
                        <div 
                          key={token.symbol}
                          className="px-4 py-3 text-white hover:bg-mystery-accent/20 cursor-pointer transition-colors"
                          onClick={() => handleToTokenSelect(token.symbol)}
                        >
                          {token.symbol}
                        </div>
                      ))}
                    </div>
                  )}
                </div>
                <input
                  type="number"
                  value={toAmount}
                  onChange={(e) => handleAmountChange(e, setToAmount)}
                  placeholder="0.00"
                  className="text-2xl font-heading bg-transparent text-right w-32 focus:outline-none placeholder-gray-500 [appearance:textfield] [&::-webkit-outer-spin-button]:appearance-none [&::-webkit-inner-spin-button]:appearance-none"
                />
              </div>
            </div>

            {/* Swap Button */}
            <button 
              className={buttonClasses}
              onClick={handleSwap}
              disabled={!fromAmount || !toAmount}
            >
              {!fromAmount || !toAmount ? 'Enter an Amount' : 'Swap'}
            </button>
          </div>
        </div>
      </div>

      {/* Loading Modal */}
      {isLoading && !transactionStatus && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
          <div className="bg-[#1E1E1E] p-8 rounded-xl text-center">
            <div className="animate-spin w-12 h-12 border-4 border-mystery-accent border-t-transparent rounded-full mx-auto mb-4"></div>
            <p className="text-white font-heading text-lg">
              Please wait while our Smart Contract find a suitable offer for you!
            </p>
          </div>
        </div>
      )}

      {/* Success/Failed Modal */}
      {transactionStatus && (
        <div 
          className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50"
          onClick={closeModal}
        >
          <div className="bg-[#1E1E1E] p-8 rounded-xl text-center" onClick={e => e.stopPropagation()}>
            <p className="text-white font-heading text-xl mb-6">
              {transactionStatus === 'success' 
                ? 'Transaction completed!' 
                : 'Sorry! We cannot find you a suitable deal'}
            </p>
            <button
              className="font-heading px-6 py-2 rounded-full bg-mystery-accent text-white hover:opacity-90 transition-opacity"
              onClick={closeModal}
            >
              Close
            </button>
          </div>
        </div>
      )}
    </div>
  );
}

export default Trade;
