import React from 'react';
import logo from '../assets/memecoin-logo.png'
import logoMeme from '../assets/memecoin-about.png'

const About = () => {
  const { useState } = React;
  const [isHovered, setIsHovered] = useState(false);
  const handleHover = () => {
    setIsHovered(!isHovered);
  };

  return (
    <div className="min-h-screen bg-gradient-to-b from-[#0B041A] to-mystery-dark">
      <div className="container mx-auto px-4 py-24">
        <div className="flex flex-col lg:flex-row items-center justify-between gap-12">
          {/* Left side - ME logo */}
          <div className="lg:w-1/2">
            <div className="relative">
              <div className="absolute -inset-4 bg-gradient-to-r from-[#503BEE] via-[#8A2BE2] to-[#361480] opacity-30 blur-lg rounded-3xl"></div>
              <img
                src={isHovered ? logoMeme : logo }
                alt="do people even read this"
                className="relative size-[600px] mx-auto transform hover:scale-105 transition-transform duration-300"
                onMouseEnter={handleHover}
                onMouseLeave={handleHover}
              />
            </div>
          </div>

          {/* Right side - About text */}
          <div className="lg:w-1/2 text center lg:text-left lg-pl-8">
            <p className="font-heading text-2xl text-gray-300 mb-12">
              Dear fren, we're not like those Hawk Tuah projects that go poof into thin air!

              <span className="block mt-2">We're a community of genuine, meme connoisseurs, and blockchain enthusiasts who believe in the power of DogeGod! Our smart contract is audited by the most prestigious good bois in the industry, and our team is fully doxxed (you can find us in the Developers section!).</span>
              <span className="block mt-2">MUCH TRUST! VERY LEGIT! 🐕</span>
            </p>
          </div>
              
        </div>
      </div>
    </div>
  );
};

export default About;
