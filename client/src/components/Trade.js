import React from 'react'

const Trade = ({ 
  isDropdownTradeFromOpen, 
  setIsDropdownTradeFromOpen,
  isDropdownTradeToOpen,
  setIsDropdownTradeToOpen
}) => {
  
  const buttonClasses = "font-heading px-6 py-2 rounded-full bg-gradient-to-r from-[#503BEE] to-[#8A2BE2] hover:from-[#8A2BE2] hover:to-[#503BEE] text-white transition-all duration-300 shadow-lg hover:shadow-mystery-accent/50 flex items-center";
  
  return (
    <div className="min-h-screen bg-gradient-to-b from-[#0B041A] to-mystery-dark">
      <div className="container mx-auto px-4 py-24">
        <div>
          <button
            className={buttonClasses}
            onClick={() => setIsDropdownTradeFromOpen(true)}
          > what
          </button>
          {isDropdownTradeFromOpen && (
            <div 
              className="absolute right-0 mt-2 w-48 bg-[#1E1E1E] border border-mystery-accent/30 rounded-xl shadow-xl backdrop-blur-sm z-10 overflow-hidden"
              onMouseLeave={() => setIsDropdownTradeFromOpen(false)}
            >
              <ul className="font-heading">
                <li className="px-4 py-3 text-white hover:bg-mystery-accent/20 cursor-pointer transition-colors duration-200">
                  Recent Transactions
                </li>
              </ul>
            </div>
          )}
        </div>
          <div className="flex items-center justify-center gap-12">
          
          {/* the trading price window */}
          <div className="relative w-full max-w-[509px] h-screen max-h-[314px] bg-[#1E1E1E] rounded-[20px] flex-col divide-y divide-[#7730E6]"> 
            
            {/* top section, from a currency */}
            <div className="shrink w-full h-1/2 px-4 py-2">
              <div className="text-gray-400">
                From
              </div>
              <div className="flex flex-row w-full max-w-[472px] h-full max-h-[91px] rounded-[30px] bg-[#372F47] justify-between items-center">
                <div className="text-2xl font-heading px-4">
                  BTC
                </div>
                 <div className="text-2xl font-heading px-4">
                  0.01
                </div>
              </div>
            </div>
            
            {/* bottom section, to a currency */}
            <div className="shrink w-full h-1/2 px-4 py-2">
              <div className="text-gray-400">
                To
              </div>
              <div className="flex flex-row w-full max-w-[472px] h-full max-h-[91px] rounded-[30px] bg-[#372F47] justify-between items-center">
                <div className="text-2xl font-heading px-4">
                  ETH
                </div>
                 <div className="text-2xl font-heading px-4">
                  0.01
                </div>
              </div>

            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Trade;
