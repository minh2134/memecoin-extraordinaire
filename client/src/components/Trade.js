import React, { useState, useEffect } from 'react'
import tradeIcon from '../assets/trade-icon.png'
import { getSimpleTokenList } from '../utils/cryptoData';
import { Decimal } from 'decimal.js';
import { API_BASE_URL } from '../utils/api.js';

const Trade = ({ 
  isDropdownTradeFromOpen, 
  setIsDropdownTradeFromOpen,
  isDropdownTradeToOpen,
  setIsDropdownTradeToOpen,
  onTransaction,
  wallet
}) => {
  // Add new state for tab selection
  const [activeTab, setActiveTab] = useState('swap'); // 'swap' or 'limit'
  
  const [fromState, setFromState] = useState('BTC');
  const [toState, setToState] = useState('DOGE');
  const [fromAmount, setFromAmount] = useState('');
  const [toAmount, setToAmount] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [transactionStatus, setTransactionStatus] = useState(null);
  const [slippage, setSlippage] = useState('5'); // Only used for swap
  
  // New state for limit orders
  const [customRate, setCustomRate] = useState('');
  const [marketRate, setMarketRate] = useState(null);
  
  const tokens = getSimpleTokenList();

  // Fetch the market rate for the selected token pair
  useEffect(() => {
    const fetchMarketRate = async () => {
      if (fromState && toState && fromState !== toState) {
        try {
          const response = await fetch(
            `${API_BASE_URL}/trade/rate?from=${fromState}&to=${toState}`
          );
          
          if (!response.ok) throw new Error('Failed to fetch market rate');
          
          const data = await response.json();
          setMarketRate(new Decimal(data.rate));
          
          // If customRate is empty, initialize it with the market rate
          if (!customRate) {
            setCustomRate(data.rate);
          }
        } catch (error) {
          console.error('Failed to fetch market rate:', error);
        }
      }
    };
    
    fetchMarketRate();
  }, [fromState, toState]);

  // Calculate the expected amount based on the active tab and rates
  useEffect(() => {
    if (!fromAmount) {
      setToAmount('');
      return;
    }
    
    if (activeTab === 'swap' && marketRate) {
      // For swap, use market rate
      const calculatedAmount = new Decimal(fromAmount).times(marketRate).toFixed(6);
      setToAmount(calculatedAmount);
    } else if (activeTab === 'limit' && customRate) {
      // For limit, use custom rate
      const calculatedAmount = new Decimal(fromAmount).times(customRate).toFixed(6);
      setToAmount(calculatedAmount);
    }
  }, [fromAmount, marketRate, customRate, activeTab]);

  const buttonClasses = `font-heading px-6 py-3 w-full rounded-[30px] bg-gradient-to-r 
    ${(!fromAmount || (activeTab === 'limit' && !customRate)) 
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
    
    // Also swap the rates for limit orders
    if (activeTab === 'limit' && marketRate) {
      setCustomRate(new Decimal(1).div(new Decimal(customRate)).toFixed(6));
    }
  };

  const handleAmountChange = (e, setter) => {
    const value = e.target.value;
    if (value === '' || (Number(value) >= 0 && !value.includes('e'))) {
      setter(value);
    }
  };

  // Handle custom rate change for limit orders
  const handleRateChange = (e) => {
    const value = e.target.value;
    
    // Validate input
    if (value === '' || (Number(value) > 0 && !value.includes('e'))) {
      setCustomRate(value);
      
      // Update to amount based on new rate if from amount exists
      if (fromAmount && value !== '') {
        try {
          // Convert both values to Decimal objects to avoid errors
          const decimalFromAmount = new Decimal(fromAmount);
          const decimalRate = new Decimal(value);
          
          // Calculate the new amount
          const calculatedAmount = decimalFromAmount.times(decimalRate).toFixed(6);
          setToAmount(calculatedAmount);
        } catch (error) {
          console.error('Error calculating amount:', error);
          // Don't update the to amount if calculation fails
        }
      } else {
        // Clear the to amount if from amount is empty or rate is empty
        setToAmount('');
      }
    }
  };

  // Handle slippage change with validation
  const handleSlippageChange = (e) => {
    const value = e.target.value;
    if (value === '' || (Number(value) >= 0 && Number(value) <= 100 && !value.includes('e'))) {
      setSlippage(value);
    }
  };

  // Set preset slippage values
  const handleSlippagePreset = (value) => {
    setSlippage(value);
  };

  const closeModal = () => {
    setIsLoading(false);
    setTransactionStatus(null);
  };

  // Swap execution (market order)
  const handleSwap = async () => {
    if (!fromAmount || !marketRate) return;
    
    setIsLoading(true);
    setTransactionStatus(null);
    
    try {
      // Calculate slippage value (as a decimal for the API)
      const slippageValue = (slippage === '' ? 5 : Number(slippage)) / 100;
      
      // Create the swap request in the format expected by the backend
      const swapRequest = {
        SourceCurr: fromState,
        TargetCurr: toState,
        SourceAmount: fromAmount,
        Rate: marketRate,
        Slippage: new Decimal(slippageValue)
      };
      
      console.log('Sending swap request:', swapRequest);
      
      // Call the API
      const response = await fetch(`${API_BASE_URL}/trade/swap`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(swapRequest)
      });
      
      if (!response.ok) {
        throw new Error('Failed to find a matching offer');
      }
      
      const result = await response.json();
      console.log('Swap result:', result);
      
      // Format the transaction for the Transactions component
      const transactionData = {
        tradedAddress: result.TradedAddress,
        tradedAmount: result.TradedAmount,
        receivedAmount: result.ReceivedAmount,
        fromCurr: result.FromCurr,
        toCurr: result.ToCurr,
        status: 'Success',
        orderType: 'Market',
        date: new Date().toLocaleString()
      };
      
      onTransaction(transactionData);
      setTransactionStatus('success');
      setFromAmount('');
      setToAmount('');
      
    } catch (error) {
      console.error('Swap failed:', error);
      
      // Format failed transaction for the Transactions component
      const failedTransaction = {
        tradedAmount: fromAmount,
        receivedAmount: '0',
        fromCurr: fromState,
        toCurr: toState,
        status: 'Failed',
        orderType: 'Market',
        date: new Date().toLocaleString(),
        errorMessage: error.message
      };
      
      onTransaction(failedTransaction);
      setTransactionStatus('failed');
    } finally {
      setIsLoading(false);
    }
  };

  // Limit order placement
  const handleLimitOrder = async () => {
    if (!fromAmount || !customRate) return;
    
    setIsLoading(true);
    setTransactionStatus(null);
    
    try {
      // Create the limit order request
      const limitRequest = {
        SourceCurr: fromState,
        TargetCurr: toState,
        SourceAmount: fromAmount,
        Rate: new Decimal(customRate)
      };
      
      console.log('Sending limit order request:', limitRequest);
      
      // Call the API
      const response = await fetch(`${API_BASE_URL}/trade/limit`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(limitRequest)
      });
      
      if (!response.ok) {
        throw new Error('Failed to place limit order');
      }
      
      const result = await response.json();
      console.log('Limit order result:', result);
      
      // If the order was matched immediately
      if (result.IsMatched) {
        const transactionData = {
          tradedAddress: result.SwapDetails.TradedAddress,
          tradedAmount: result.SwapDetails.TradedAmount,
          receivedAmount: result.SwapDetails.ReceivedAmount,
          fromCurr: result.SwapDetails.FromCurr,
          toCurr: result.SwapDetails.ToCurr,
          rate: customRate,
          status: 'Success',
          orderType: 'Limit (Filled)',
          date: new Date().toLocaleString()
        };
        
        onTransaction(transactionData);
      } else {
        // Order was placed but not matched yet
        const pendingOrderData = {
          tradedAmount: fromAmount,
          receivedAmount: toAmount,
          fromCurr: fromState,
          toCurr: toState,
          rate: customRate,
          status: 'Pending',
          orderType: 'Limit (Open)',
          date: new Date().toLocaleString()
        };
        
        onTransaction(pendingOrderData);
      }
      
      setTransactionStatus('success');
      setFromAmount('');
      setToAmount('');
      setCustomRate('');
      
    } catch (error) {
      console.error('Limit order failed:', error);
      
      const failedTransaction = {
        tradedAmount: fromAmount,
        receivedAmount: toAmount,
        fromCurr: fromState,
        toCurr: toState,
        rate: customRate,
        status: 'Failed',
        orderType: 'Limit',
        date: new Date().toLocaleString(),
        errorMessage: error.message
      };
      
      onTransaction(failedTransaction);
      setTransactionStatus('failed');
    } finally {
      setIsLoading(false);
    }
  };

  // Update the validation functions
  const canExecuteSwap = () => {
    return fromAmount && 
           parseFloat(fromAmount) > 0 && 
           fromState !== toState;
  };

  const canExecuteLimitOrder = () => {
    return fromAmount && 
           parseFloat(fromAmount) > 0 && 
           customRate && 
           parseFloat(customRate) > 0 && 
           fromState !== toState;
  };

  return (
    <div className="min-h-screen bg-gradient-to-b from-[#0B041A] to-mystery-dark">
      <div className="container mx-auto px-4 py-24">
        <div className="flex items-center justify-center">
          <div className="relative w-full max-w-[509px] bg-[#1E1E1E] rounded-[20px] p-6 space-y-6">
            {/* Tab Selector */}
            <div className="flex rounded-xl overflow-hidden border border-mystery-accent/20 mb-6">
              <button
                className={`flex-1 py-3 px-6 font-heading text-lg ${
                  activeTab === 'swap'
                    ? 'bg-mystery-accent text-white'
                    : 'bg-[#372F47] text-gray-300 hover:bg-[#4a3f5e] transition-colors'
                }`}
                onClick={() => setActiveTab('swap')}
              >
                Swap
              </button>
              <button
                className={`flex-1 py-3 px-6 font-heading text-lg ${
                  activeTab === 'limit'
                    ? 'bg-mystery-accent text-white'
                    : 'bg-[#372F47] text-gray-300 hover:bg-[#4a3f5e] transition-colors'
                }`}
                onClick={() => setActiveTab('limit')}
              >
                Limit
              </button>
            </div>

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
                  disabled={activeTab === 'limit'} // Disable direct input in limit mode
                  onChange={(e) => handleAmountChange(e, setToAmount)}
                  placeholder="0.00"
                  className={`text-2xl font-heading bg-transparent text-right w-32 focus:outline-none placeholder-gray-500 
                    [appearance:textfield] [&::-webkit-outer-spin-button]:appearance-none [&::-webkit-inner-spin-button]:appearance-none
                    ${activeTab === 'limit' ? 'opacity-70 cursor-not-allowed' : ''}`}
                />
              </div>
            </div>

            {/* Market Rate Display */}
            {marketRate && (
              <div className="mb-2 text-sm text-gray-400">
                <div className="flex justify-between items-center">
                  <span>Market Rate: 1 {fromState} = {marketRate.toFixed(6)} {toState}</span>
                </div>
              </div>
            )}

            {/* Conditional UI based on active tab */}
            {activeTab === 'swap' ? (
              // Slippage Tolerance Section (only for swap)
              <div className="w-full">
                <div className="text-gray-400 mb-2">Slippage Tolerance</div>
                <div className="flex flex-col gap-3">
                  {/* Presets for common slippage values */}
                  <div className="flex flex-row gap-2">
                    {['0.5', '1', '3', '5'].map((value) => (
                      <button
                        key={value}
                        onClick={() => handleSlippagePreset(value)}
                        className={`px-3 py-2 rounded-lg ${
                          slippage === value
                            ? 'bg-mystery-accent text-white'
                            : 'bg-[#372F47] text-gray-300 hover:bg-[#4a3f5e] transition-colors'
                        }`}
                      >
                        {value}%
                      </button>
                    ))}
                  </div>
                  
                  {/* Custom slippage input */}
                  <div className="flex flex-row rounded-lg bg-[#372F47] overflow-hidden">
                    <input
                      type="number"
                      value={slippage}
                      onChange={handleSlippageChange}
                      placeholder="5"
                      className="flex-grow p-3 bg-transparent focus:outline-none text-white [appearance:textfield] [&::-webkit-outer-spin-button]:appearance-none [&::-webkit-inner-spin-button]:appearance-none"
                    />
                    <div className="bg-[#2e273a] px-4 flex items-center text-gray-300">%</div>
                  </div>
                </div>
              </div>
            ) : (
              // Limit Rate Section (only for limit)
              <div className="w-full">
                <div className="text-gray-400 mb-2">Limit Rate</div>
                <div className="flex flex-col gap-3">
                  {/* Custom rate input */}
                  <div className="flex flex-row rounded-lg bg-[#372F47] overflow-hidden">
                    <div className="bg-[#2e273a] px-4 flex items-center text-gray-300">1 {fromState} =</div>
                    <input
                      type="number"
                      value={customRate}
                      onChange={handleRateChange}
                      placeholder={marketRate ? marketRate.toString() : "0.00"}
                      className="flex-grow p-3 bg-transparent focus:outline-none text-white [appearance:textfield] [&::-webkit-outer-spin-button]:appearance-none [&::-webkit-inner-spin-button]:appearance-none"
                    />
                    <div className="bg-[#2e273a] px-4 flex items-center text-gray-300">{toState}</div>
                  </div>
                  
                  {/* Rate comparison */}
                  {marketRate && customRate && (
                    <div className="text-sm">
                      {new Decimal(customRate).greaterThan(marketRate) ? (
                        <span className="text-green-400">
                          +{new Decimal(customRate).minus(marketRate).div(marketRate).times(100).toFixed(2)}% above market
                        </span>
                      ) : (
                        <span className="text-red-400">
                          {new Decimal(customRate).minus(marketRate).div(marketRate).times(100).toFixed(2)}% below market
                        </span>
                      )}
                    </div>
                  )}
                  
                  {/* Rate explanation */}
                  <div className="text-sm text-gray-400">
                    Your order will execute when the market rate reaches your limit rate.
                  </div>
                </div>
              </div>
            )}

            {/* Action Button based on active tab */}
            {activeTab === 'swap' ? (
              <button 
                className={buttonClasses}
                onClick={handleSwap}
                disabled={!fromAmount || !toAmount}
              >
                {!fromAmount || !toAmount ? 'Enter an Amount' : 'Swap'}
              </button>
            ) : (
              <button 
                className={buttonClasses}
                onClick={handleLimitOrder}
                disabled={!fromAmount || !customRate}
              >
                {!fromAmount ? 'Enter an Amount' : !customRate ? 'Enter a Rate' : 'Place Limit Order'}
              </button>
            )}
          </div>
        </div>
      </div>

      {/* Loading Modal */}
      {isLoading && !transactionStatus && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
          <div className="bg-[#1E1E1E] p-8 rounded-xl text-center">
            <div className="animate-spin w-12 h-12 border-4 border-mystery-accent border-t-transparent rounded-full mx-auto mb-4"></div>
            <p className="text-white font-heading text-lg">
              {activeTab === 'swap' 
                ? 'Please wait while our Smart Contract finds a suitable offer for you!' 
                : 'Placing your limit order...'}
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
                ? activeTab === 'swap'
                  ? 'Transaction completed!' 
                  : 'Limit order placed successfully!'
                : activeTab === 'swap'
                  ? 'Sorry! We cannot find you a suitable deal'
                  : 'Failed to place limit order'}
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
