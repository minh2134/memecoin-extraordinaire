//@author Dinh Le Hoang Anh - 105508318
//@author Pham Vu Minh - 105110564
import { Decimal } from 'decimal.js';

const API_BASE_URL = 'http://localhost:8080';

export const executeSwap = async (swapRequest) => {
  try {
    const response = await fetch(`${API_BASE_URL}/trade/swap`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(swapRequest)
    });

    if (!response.ok) {
      throw new Error('No matching orders found');
    }

    const result = await response.json();
    return {
      tradedAddress: result.tradedAddress,
      tradedAmount: new Decimal(result.tradedAmount),
      receivedAmount: new Decimal(result.receivedAmount),
      fromCurr: result.fromCurr,
      toCurr: result.toCurr
    };
  } catch (error) {
    console.error('Swap execution failed:', error);
    throw error;
  }
};

export const getExchangeRate = async (fromToken, toToken) => {
  try {
    const response = await fetch(
      `${API_BASE_URL}/trade/rate?from=${fromToken}&to=${toToken}`
    );
    
    if (!response.ok) {
      throw new Error('Failed to fetch exchange rate');
    }
    
    const data = await response.json();
    return new Decimal(data.rate);
  } catch (error) {
    console.error('Failed to get exchange rate:', error);
    throw error;
  }
};

export const calculateSwapAmount = (amount, rate) => {
  if (!amount || !rate) return '0';
  return new Decimal(amount).times(rate).toFixed(6);
}; 