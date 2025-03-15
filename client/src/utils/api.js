const API_BASE_URL = 'http://localhost:8080';

export const swapTokens = async (swapRequest) => {
  const response = await fetch(`${API_BASE_URL}/trade/swap`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(swapRequest)
  });
  
  if (!response.ok) {
    throw new Error('Swap request failed');
  }
  
  return response.json();
};

// Add other API endpoints as needed
