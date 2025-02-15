import React from 'react';
import memelogo from '../assets/memecoin.png';
import { FaWallet, FaArrowDown } from 'react-icons/fa';

const Navbar = ({ walletAddress, setIsWalletOpen, isDropdownOpen, setIsDropdownOpen, handleDisconnect }) => {
  const buttonClasses = "font-heading px-6 py-2 rounded-full bg-transparent border-2 border-gradient-to-r from-[#E5E5E5] to-[#8A2BE2] text-white hover:opacity-90 transition-opacity flex items-center";

  return (
    <nav className="bg-gradient-to-r from-[#0B041A] to-[#361480] px-6 py-2">
      <div className="container mx-auto flex items-center justify-between">
        <div className="flex items-center">
          <img src={memelogo} alt="Memecoin Logo" className="h-[93px] w-auto" />
        </div>
        
        <div className="flex items-center space-x-8">
          <NavLink text="Home" />
          <NavLink text="Trade" highlight={true} />
          <NavLink text="Market" />
          <NavLink text="About" />
          <NavLink text="Developers" />
          
          <div className="relative">
            {!walletAddress ? (
              <button 
                className={buttonClasses}
                onClick={() => setIsWalletOpen(true)}
              >
                <span>Link Wallet</span>
                <FaWallet className="ml-2" />
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
                    className="absolute right-0 mt-2 w-48 bg-[#1E1E1E] border-2 border-gradient-to-r from-[#E5E5E5] to-[#8A2BE2] rounded-lg shadow-lg z-10"
                    onMouseLeave={() => setIsDropdownOpen(false)}
                  >
                    <ul className="font-heading">
                      <li className="px-4 py-2 text-white hover:bg-mystery-accent/20 cursor-pointer">
                        Recent Transactions
                      </li>
                      <li 
                        className="px-4 py-2 text-white hover:bg-mystery-accent/20 cursor-pointer"
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

const NavLink = ({ text, highlight = false }) => (
  <a
    href="#"
    className={`font-heading text-lg ${
      highlight 
        ? 'text-mystery-highlight' 
        : 'text-white hover:text-mystery-accent'
    } transition-colors`}
  >
    {text}
  </a>
);

export default Navbar; 