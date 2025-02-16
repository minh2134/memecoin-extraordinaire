import React from 'react';
import { useNavigate } from 'react-router-dom';
import cheems from '../assets/cheems.png';
import dogeWow from '../assets/doge-wow.png';

const Hero = ({ walletAddress, setIsWalletOpen }) => {
  const navigate = useNavigate();

  return (
    <div className="relative overflow-hidden">
      {/* Bright upper section */}
      <div className="bg-[#503BEE] bg-opacity-20 pt-16 pb-32">
        <div className="container mx-auto px-4">
          <div className="flex flex-col lg:flex-row items-center justify-between relative">
            {/* Left side content */}
            <div className="lg:w-1/2 z-10">
              <h1 className="font-heading text-5xl lg:text-7xl leading-tight bg-gradient-to-r from-[#503BEE] via-[#8A2BE2] to-[#361480] text-transparent bg-clip-text mb-6">
                Welcome to Memecoin Extraordinaire!
              </h1>
              <p className="font-body text-xl lg:text-2xl text-gray-300 mb-8">
                The ultimate trading platform for meme lovers and crypto enthusiasts!
              </p>
            </div>

            {/* Right side with memes */}
            <div className="lg:w-1/2 relative h-[400px]">
              {/* Main doge-wow in the middle-right */}
              <img 
                src={dogeWow} 
                alt="Doge" 
                className="absolute right-20 top-1/2 transform -translate-y-1/2 w-48 lg:w-64 animate-float z-20" 
              />
              
              {/* Additional meme decorations */}
              <div className="absolute top-0 right-0 w-32 h-32 bg-mystery-accent/10 rounded-full animate-pulse" />
              <div className="absolute bottom-20 right-40 w-24 h-24 bg-mystery-highlight/10 rounded-full animate-pulse delay-300" />
              
              {/* Floating emojis */}
              <span className="absolute top-10 right-10 text-4xl animate-float">🚀</span>
              <span className="absolute bottom-20 right-10 text-4xl animate-float-delayed">💎</span>
              <span className="absolute top-1/2 right-[300px] text-4xl animate-float">🌙</span>
            </div>
          </div>
        </div>
      </div>

      {/* Dark space gradient section */}
      <div className="bg-gradient-to-b from-[#0B041A] to-mystery-dark py-16">
        <div className="container mx-auto px-4">
          <div className="flex flex-col lg:flex-row items-center justify-between">
            <div className="lg:w-1/2 z-10">
              {!walletAddress ? (
                <div className="flex gap-4">
                  <button 
                    className="font-heading px-8 py-3 rounded-full bg-mystery-accent text-white hover:opacity-90 transition-opacity"
                    onClick={() => setIsWalletOpen(true)}
                  >
                    Link Wallet
                  </button>
                  <button 
                    className="font-heading px-8 py-3 rounded-full bg-transparent border-2 border-mystery-accent text-white hover:opacity-90 transition-opacity"
                    onClick={() => navigate('/404')}
                  >
                    Get a Wallet
                  </button>
                </div>
              ) : (
                <div className="text-center lg:text-left">
                  <h3 className="font-heading text-3xl bg-gradient-to-r from-[#FF6B6B] via-[#FFE66D] to-[#4ECDC4] text-transparent bg-clip-text mb-4">
                    Much Success! Very Connect! 🎉
                  </h3>
                  <p className="font-heading text-xl text-gray-300">
                    Wow! Such smart trader! Your wallet is now ready to do the big stonks! 
                    <span className="block mt-2">Time to trade like pro doge! To the moon! 🚀🌕</span>
                  </p>
                </div>
              )}
            </div>
            <div className={`${walletAddress ? 'w-full' : 'lg:w-1/2'} relative mt-8 lg:mt-0`}>
              <div className="flex flex-col items-center">
                <img 
                  src={cheems} 
                  alt="Cheems" 
                  className="w-48 lg:w-64 animate-float mb-4" 
                />
                <p className="font-heading text-2xl bg-gradient-to-r from-[#FF6B6B] via-[#FFE66D] to-[#4ECDC4] text-transparent bg-clip-text p-4 border-2 border-[#4ECDC4] rounded-xl shadow-[0_0_15px_rgba(78,205,196,0.5)] animate-pulse">
                  TRADE MEME! LIKE CHEEMS!
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Hero; 
