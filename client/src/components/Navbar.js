import React from 'react';
import { useNavigate, Link } from 'react-router-dom';
import memelogo from '../assets/memecoin.png';
import { FaWallet, FaArrowDown } from 'react-icons/fa';

const Navbar = ({ walletAddress, setIsWalletOpen, isDropdownOpen, setIsDropdownOpen, handleDisconnect }) => {
  const navigate = useNavigate();
  const buttonClasses = "font-heading px-6 py-2 rounded-full bg-gradient-to-r from-[#503BEE] to-[#8A2BE2] hover:from-[#8A2BE2] hover:to-[#503BEE] text-white transition-all duration-300 shadow-lg hover:shadow-mystery-accent/50 flex items-center";

  const NavLink = ({ text, onClick, highlight }) => (
    <button
      onClick={onClick}
      className={`font-heading relative group ${
        highlight 
          ? 'text-2xl font-bold mx-4' 
          : 'text-white text-lg mx-4'
      } transition-colors`}
    >
      {highlight ? (
        <span className="text-[#FFE600] group-hover:text-white group-hover:border-[#FFE600] transition-all duration-300">
          {text}
        </span>
      ) : (
        <>
          {text}
          <span className="absolute -bottom-1 left-0 w-0 h-0.5 bg-mystery-accent transition-all duration-300 group-hover:w-full"></span>
        </>
      )}
    </button>
  );

  return (
    <nav className="bg-gradient-to-r from-[#0B041A] to-[#361480] px-6 py-0 shadow-lg">
      <div className="container mx-auto flex items-center justify-between">
        <div className="flex items-center">
          <img 
            src={memelogo} 
            alt="Memecoin Logo" 
            className="h-[120px] w-auto cursor-pointer hover:scale-105 transition-transform duration-300" 
            onClick={() => navigate('/')}
          />
        </div>
        
        <div className="hidden md:flex items-center space-x-10">
          <NavLink text="Home" onClick={() => navigate('/')} />
          <NavLink text="Trade" highlight={true} />
          <NavLink text="Market" onClick={() => navigate('/market')} />
          <NavLink text="Developers" onClick={() => navigate('/developers')} />
          <NavLink text="About" onClick={() => navigate('/about')} />
        
          <div className="relative ml-4">
            {!walletAddress ? (
              <button 
                className={buttonClasses}
                onClick={() => setIsWalletOpen(true)}
              >
                <span>Link Wallet</span>
                <FaWallet className="ml-2 animate-pulse" />
              </button>
            ) : (
              <div>
                <button 
                  className={buttonClasses}
                  onClick={() => setIsDropdownOpen(true)}
                >
                  <FaArrowDown className="mr-2" />
                  <span>{walletAddress}</span>
                  <FaWallet className="ml-2" />
                </button>
                {isDropdownOpen && (
                  <div 
                    className="absolute right-0 mt-2 w-48 bg-[#1E1E1E] border border-mystery-accent/30 rounded-xl shadow-xl backdrop-blur-sm z-10 overflow-hidden"
                    onMouseLeave={() => setIsDropdownOpen(false)}
                  >
                    <ul className="font-heading">
                      <li className="px-4 py-3 text-white hover:bg-mystery-accent/20 cursor-pointer transition-colors duration-200">
                        Recent Transactions
                      </li>
                      <li 
                        className="px-4 py-3 text-white hover:bg-mystery-accent/20 cursor-pointer transition-colors duration-200"
                        onClick={handleDisconnect}
                      >
                        Disconnect
                      </li>
                    </ul>
                  </div>
                )}
              </div>
            )}
          </div>
        </div>
      </div>
    </nav>
  );
};

export default Navbar; 
