import React from 'react';
import dogeComrade from '../assets/doge-comrade.png';

const Developers = () => {
  return (
    <div className="min-h-screen bg-gradient-to-b from-[#0B041A] to-mystery-dark">
      <div className="container mx-auto px-4 py-24">
        <div className="flex flex-col lg:flex-row items-center justify-between gap-12">
          {/* Left side - Doge Comrade image */}
          <div className="lg:w-1/2">
            <div className="relative">
              <div className="absolute -inset-4 bg-gradient-to-r from-[#503BEE] via-[#8A2BE2] to-[#361480] opacity-30 blur-lg rounded-3xl"></div>
              <img 
                src={dogeComrade} 
                alt="Doge Comrade" 
                className="relative w-full max-w-[600px] mx-auto transform hover:scale-105 transition-transform duration-300"
              />
            </div>
          </div>
          
          {/* Right side - Catchphrase and developer list */}
          <div className="lg:w-1/2 text-center lg:text-left lg:pl-8">
            <p className="font-heading text-2xl text-gray-300 mb-12">
              Two comrades joined forces to create this amazing platform for fellow meme lovers! 
              <span className="block mt-2">Such wow! Many features! 🚀</span>
            </p>
            <div className="space-y-6">
              <div className="p-6 bg-[#1E1E1E]/50 rounded-xl border-2 border-mystery-accent/30 hover:border-mystery-accent transition-colors duration-300">
                <h3 className="font-heading text-2xl text-mystery-highlight mb-2">PHAM VU MINH</h3>
                <p className="font-heading text-gray-300">Software Engineering</p>
              </div>
              <div className="p-6 bg-[#1E1E1E]/50 rounded-xl border-2 border-mystery-accent/30 hover:border-mystery-accent transition-colors duration-300">
                <h3 className="font-heading text-2xl text-mystery-highlight mb-2">DINH LE HOANG ANH</h3>
                <p className="font-heading text-gray-300">Data Analyst</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Developers; 