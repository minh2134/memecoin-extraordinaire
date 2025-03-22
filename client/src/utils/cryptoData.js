//@author Dinh Le Hoang Anh - 105508318
//@author Pham Vu Minh - 105110564
// Centralized cryptocurrency data for reuse across the application
export const cryptoList = [
  { 
    id: 'bitcoin', 
    symbol: 'BTC', 
    name: 'Bitcoin', 
    nicknames: ['btc', 'bitcoin', 'digital gold', 'crypto king']
  },
  { 
    id: 'dai', 
    symbol: 'DAI', 
    name: 'Dai', 
    nicknames: ['dai', 'stablecoin', 'maker']
  },
  { 
    id: 'havven', 
    symbol: 'SNX', 
    name: 'Synthetix', 
    nicknames: ['snx', 'synthetix', 'synthetic assets']
  }
];

// Helper function to get all coin symbols
export const getAllCoinSymbols = () => cryptoList.map(coin => coin.symbol);

// Helper function to get all coin IDs for API calls
export const getAllCoinIds = () => cryptoList.map(coin => coin.id);

// Helper function to find a coin by symbol
export const getCoinBySymbol = (symbol) => cryptoList.find(coin => coin.symbol === symbol);

// Helper function to find a coin by ID
export const getCoinById = (id) => cryptoList.find(coin => coin.id === id);

// Helper function to get simplified token list (for dropdowns, etc.)
export const getSimpleTokenList = () => cryptoList.map(coin => ({
  id: coin.id,
  name: coin.name,
  symbol: coin.symbol,
  image: coin.id === 'bitcoin' 
    ? 'https://assets.coingecko.com/coins/images/1/large/bitcoin.png'
    : coin.id === 'dai' 
    ? 'https://assets.coingecko.com/coins/images/9956/large/4943.png'
    : 'https://assets.coingecko.com/coins/images/3406/large/SNX.png'
})); 