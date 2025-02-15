import React from 'react';
import memelogo from '../assets/memecoin.png';

const Navbar = () => {
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
          <button className="font-heading px-6 py-2 rounded-full bg-transparent border-2 border-gradient-to-r from-[#E5E5E5] to-[#8A2BE2] text-white hover:opacity-90 transition-opacity">
            Link Wallet
          </button>
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