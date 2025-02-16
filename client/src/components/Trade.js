import React, { useState } from 'react'
import { cryptoList } from './Market'

const Trade = ({ 
  isDropdownTradeFromOpen, 
  setIsDropdownTradeFromOpen,
  isDropdownTradeToOpen,
  setIsDropdownTradeToOpen
}) => {
  
  const [fromState, setFromState] = useState('BTC');
  const [toState, setToState] = useState('ETH');
  const buttonClasses = "font-heading px-6 py-2 rounded-full bg-gradient-to-r from-[#503BEE] to-[#8A2BE2] hover:from-[#8A2BE2] hover:to-[#503BEE] text-white transition-all duration-300 shadow-lg hover:shadow-mystery-accent/50 flex items-center";
  
  return (
    <div className="min-h-screen bg-gradient-to-b from-[#0B041A] to-mystery-dark">
      <div className="container mx-auto px-4 py-24">
        
          <div className="flex items-center justify-center gap-12">
          
          {/* the trading price window */}
          <div className="relative w-full max-w-[509px] h-screen max-h-[314px] bg-[#1E1E1E] rounded-[20px] flex-col divide-y divide-[#7730E6]"> 
            
            {/* top section, from a currency */}
            <div className="shrink w-full h-1/2 px-4 py-2">
              <div className="text-gray-400">
                From
              </div>
              <div className="flex flex-row w-full max-w-[472px] h-full max-h-[91px] rounded-[30px] bg-[#372F47] justify-between items-center">
                <div>
                  <div
                    className='text-2xl font-heading px-4'
                    onClick={() => setIsDropdownTradeFromOpen(true)}
                  > 
                    {fromState}
                  </div>
                  {isDropdownTradeFromOpen && (
                    <div 
                      className="absolute left-0 mt-2 w-48 bg-[#1E1E1E] border border-mystery-accent/30 rounded-xl shadow-xl backdrop-blur-sm z-10 overflow-hidden"
                      onMouseLeave={() => setIsDropdownTradeFromOpen(false)}
                    >
                      <ul className="font-heading">
                        { cryptoList.map(item => {
                          return (
                            <li 
                              className="px-4 py-3 text-white hover:bg-mystery-accent/20 cursor-pointer transition-colors duration-200"
                              onClick = {() => setFromState( item.symbol )}
                            >  
                              { item.symbol }
                            </li>
                          )
                        })}
                      </ul>
                    </div>
                  )}
                </div>
                <div className="text-2xl font-heading px-4">
                  1.00
                </div>
              </div>
            </div>
            
            {/* bottom section, to a currency */}
            <div className="shrink w-full h-1/2 px-4 py-2">
              <div className="text-gray-400">
                To
              </div>
              <div className="flex flex-row w-full max-w-[472px] h-full max-h-[91px] rounded-[30px] bg-[#372F47] justify-between items-center">
                <div>
                  <div
                    className='text-2xl font-heading px-4'
                    onClick={() => setIsDropdownTradeToOpen(true)}
                  > 
                    { toState }
                  </div>
                  {isDropdownTradeToOpen && (
                    <div 
                      className="absolute left-0 mt-2 w-48 bg-[#1E1E1E] border border-mystery-accent/30 rounded-xl shadow-xl backdrop-blur-sm z-10 overflow-hidden"
                      onMouseLeave={() => setIsDropdownTradeToOpen(false)}
                    >
                      <ul className="font-heading">
                        { cryptoList.map(item => {
                          return (
                            <li 
                              className="px-4 py-3 text-white hover:bg-mystery-accent/20 cursor-pointer transition-colors duration-200" 
                              onClick={ () => setToState( item.symbol ) }
                            >
                              { item.symbol }
                            </li>
                          )
                        })}
                      </ul>
                    </div>
                  )}
                </div>
                 <div className="text-2xl font-heading px-4">
                  35.99
                </div>
              </div>
            </div>
            <button className={buttonClasses}>Swap</button>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Trade;
