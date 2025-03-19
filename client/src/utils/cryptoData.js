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
    id: 'dogecoin', 
    symbol: 'DOGE', 
    name: 'Dogecoin', 
    nicknames: ['doge', 'dogecoin', 'much wow']
  },
  { 
    id: 'shiba-inu', 
    symbol: 'SHIB', 
    name: 'Shiba Inu', 
    nicknames: ['shib', 'shiba', 'doge killer']
  },
  { 
    id: 'bonk', 
    symbol: 'BONK', 
    name: 'Bonk', 
    nicknames: ['bonk', 'solana dog']
  },
  { 
    id: 'pepe', 
    symbol: 'PEPE', 
    name: 'Pepe', 
    nicknames: ['pepe', 'frog coin']
  },
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
  name: coin.name,
  symbol: coin.symbol
})); 