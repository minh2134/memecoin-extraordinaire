import Web3 from 'web3';
import { Decimal } from 'decimal.js';

const API_BASE_URL = 'http://localhost:8080';
const CURRENCIES = ['BTC', 'DOGE', 'SHIB', 'PEPE', 'BONK'];
const NUM_ORDERS = 50; // Number of random orders to generate

// Random number between min and max
const random = (min, max) => Math.random() * (max - min) + min;

// Random currency pair (not the same currency)
const randomPair = () => {
  const from = CURRENCIES[Math.floor(Math.random() * CURRENCIES.length)];
  let to;
  do {
    to = CURRENCIES[Math.floor(Math.random() * CURRENCIES.length)];
  } while (to === from);
  return [from, to];
};

export const seedLimitOrders = async () => {
  const web3 = new Web3('http://localhost:7545');
  
  try {
    // Get all accounts from Ganache
    const accounts = await web3.eth.getAccounts();
    
    for (let i = 0; i < NUM_ORDERS; i++) {
      const [fromCurrency, toCurrency] = randomPair();
      const sourceAmount = new Decimal(random(100, 10000)).toFixed(2);
      const rate = new Decimal(random(0.1, 10)).toFixed(6);
      
      // Random account from Ganache
      const sourceAddress = accounts[Math.floor(Math.random() * accounts.length)];

      const limitOrder = {
        sourceAddress,
        sourceCurr: fromCurrency,
        targetCurr: toCurrency,
        sourceAmount: new Decimal(sourceAmount),
        rate: new Decimal(rate)
      };

      // Place limit order
      await fetch(`${API_BASE_URL}/trade/limit`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(limitOrder)
      });

      console.log(`Created order: ${fromCurrency}->${toCurrency}, Amount: ${sourceAmount}, Rate: ${rate}`);
    }

    console.log(`Successfully created ${NUM_ORDERS} random limit orders`);
  } catch (error) {
    console.error('Failed to seed limit orders:', error);
  }
}; 