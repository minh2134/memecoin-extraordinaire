// Wallet service for interacting with wallet API endpoints

const API_URL = 'http://localhost:8080';

export const createWallet = async (name) => {
  try {
    const response = await fetch(`${API_URL}/wallet/create`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ name }),
    });
    
    if (!response.ok) {
      throw new Error(`Error: ${response.status}`);
    }
    
    return await response.json();
  } catch (error) {
    console.error('Failed to create wallet:', error);
    throw error;
  }
};

export const addFunds = async (address, currency, amount) => {
  try {
    // Convert amount to integer (amount * 100) as expected by the API
    const amountInt = Math.round(amount * 100);
    
    const response = await fetch(`${API_URL}/wallet/add-funds`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        address,
        currency,
        amount: amountInt,
      }),
    });
    
    if (!response.ok) {
      throw new Error(`Error: ${response.status}`);
    }
    
    return await response.json();
  } catch (error) {
    console.error('Failed to add funds:', error);
    throw error;
  }
};

export const getBalance = async (address) => {
  try {
    const response = await fetch(`${API_URL}/wallet/balance?address=${address}`);
    
    if (!response.ok) {
      throw new Error(`Error: ${response.status}`);
    }
    
    return await response.json();
  } catch (error) {
    console.error('Failed to get balance:', error);
    throw error;
  }
};

// Add other wallet service functions here 