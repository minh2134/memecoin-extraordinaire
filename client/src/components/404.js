import React from 'react';
import { useNavigate } from 'react-router-dom';
import memelogo from '../assets/memecoin-logo.png';

const NotFound = () => {
  const navigate = useNavigate();

  return (
    <div className="container mx-auto px-4 py-16">
      <div className="flex flex-col lg:flex-row items-center justify-between gap-8">
        {/* Left side - Large logo */}
        <div className="lg:w-1/2">
          <img 
            src={memelogo} 
            alt="Memecoin Logo" 
            className="w-full max-w-[500px] mx-auto"
          />
        </div>
        
        {/* Right side - 404 message */}
        <div className="lg:w-1/2 text-center lg:text-left">
          <h1 className="font-heading text-8xl bg-gradient-to-r from-[#FF6B6B] via-[#FFE66D] to-[#4ECDC4] text-transparent bg-clip-text mb-6">
            404
          </h1>
          <h2 className="font-heading text-4xl text-white mb-4">
            MUCH EMPTY! VERY LOST! 🐕
          </h2>
          <p className="font-heading text-2xl text-gray-300 mb-8">
            WOW! Such wrong path! The memes you seek are not here, fren!
          </p>
          <button 
            onClick={() => navigate('/')}
            className="font-heading px-8 py-3 rounded-full bg-mystery-accent text-white hover:opacity-90 transition-opacity"
          >
            Return to Meme Heaven
          </button>
        </div>
      </div>
    </div>
  );
};

export default NotFound; 