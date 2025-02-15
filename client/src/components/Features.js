import React from 'react';

const Features = () => {
  return (
    <section className="py-20">
      <h2 className="font-heading text-5xl text-white text-center mb-16 bg-gradient-to-r from-[#503BEE] via-[#8A2BE2] to-[#361480] text-transparent bg-clip-text">
        Why Choose Memecoin Extraordinaire?
      </h2>
      <div className="grid grid-cols-1 md:grid-cols-3 gap-12">
        {features.map((feature, index) => (
          <div key={index} className="bg-gradient-to-br from-[#0B041A] to-[#361480] p-8 rounded-3xl border border-mystery-accent/20 hover:border-mystery-accent/40 transition-colors">
            <div className="text-5xl mb-6">{feature.icon}</div>
            <h3 className="font-heading text-2xl text-white mb-4 bg-gradient-to-r from-[#503BEE] to-[#8A2BE2] text-transparent bg-clip-text">
              {feature.title}
            </h3>
            <p className="font-body text-lg text-gray-300 leading-relaxed">
              {feature.description}
            </p>
          </div>
        ))}
      </div>
    </section>
  );
};

const features = [
  {
    icon: "🔒",
    title: "Secure Trading",
    description: "Enterprise-grade security with multi-signature wallets and advanced encryption. Your memes are safe with us!"
  },
  {
    icon: "⚡",
    title: "Lightning Fast",
    description: "Execute trades instantly with our high-performance matching engine. Because memes wait for no one!"
  },
  {
    icon: "🌟",
    title: "Community First",
    description: "Join a vibrant community of meme coin traders and enthusiasts. Together, we go to the moon! 🚀"
  }
];

export default Features; 