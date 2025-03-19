import React, { useState } from 'react';
import { useWallet } from '../contexts/WalletContext';
import { useNavigate } from 'react-router-dom';

function Login() {
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState('');
  const { assignWallet } = useWallet();
  const navigate = useNavigate();

  const handleLogin = async (e) => {
    e.preventDefault();
    setIsLoading(true);
    setError('');

    try {
      // Assign a random prefunded wallet from Kurtosis
      const walletInfo = await assignWallet();
      console.log('Wallet assigned:', walletInfo);

      // Navigate to the main page or dashboard
      navigate('/');
    } catch (err) {
      console.error("Login error:", err);
      setError(err.message);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div>
      <h2>Login</h2>
      {error && <div className="error">{error}</div>}
      <form onSubmit={handleLogin}>
        {/* Your login form fields */}
        <button type="submit" disabled={isLoading}>
          {isLoading ? 'Connecting to Kurtosis...' : 'Login'}
        </button>
      </form>
    </div>
  );
}

export default Login; 