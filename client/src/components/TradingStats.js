import React from 'react';

const TradingStats = () => {
  return (
    <section className="py-20">
      <div className="bg-mystery-accent/5 rounded-3xl border border-mystery-accent/20 p-16">
        <h2 className="font-heading text-5xl text-white text-center mb-16 bg-gradient-to-r from-[#503BEE] via-[#8A2BE2] to-[#361480] text-transparent bg-clip-text">
          Trading Statistics
        </h2>
        <div className="grid grid-cols-1 md:grid-cols-4 gap-12">
          {stats.map((stat, index) => (
            <div key={index} className="text-center group">
              <p className="font-heading text-5xl bg-gradient-to-r from-[#503BEE] to-[#8A2BE2] text-transparent bg-clip-text mb-4 transform transition-transform group-hover:scale-110">
                {stat.value}
              </p>
              <p className="font-body text-xl text-gray-300">
                {stat.label}
              </p>
            </div>
          ))}
        </div>
      </div>
    </section>
  );
};

const stats = [
  {
    value: "$5+",
    label: "24h Trading Volume"
  },
  {
    value: "2",
    label: "Active Traders"
  },
  {
    value: "5",
    label: "Listed Memecoins"
  },
  {
    value: "<3s",
    label: "Trading Latency"
  }
];

export default TradingStats; 