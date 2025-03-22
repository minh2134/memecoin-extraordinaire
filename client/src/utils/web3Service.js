//@author Dinh Le Hoang Anh - 105508318
//@author Pham Vu Minh - 105110564
import Web3 from 'web3';
import TradingContractABI from '../contracts/Trading.json'; // ABI for the smart contract

class Web3Service {
  constructor() {
    this.web3 = null;
    this.contract = null;
    this.account = null;
    this.isInitialized = false;
    this.contractAddress = "0x050A1644A9A5364e9c6bc42d7A7B10Dd1beF07d0"; // Using the specified address
  }

  async initialize() {
    try {
      // Connect to Sepolia instead of Kurtosis
      this.web3 = new Web3('https://sepolia.infura.io/v3/5039db0cac5f44b3bd15771581f51116');
      
      // Check connection by getting network ID
      const networkId = await this.web3.eth.net.getId();
      console.log(`Connected to Sepolia testnet. Network ID: ${networkId}`);
      
      this.isInitialized = true;
      
      // Load the contract with the specified address
      this.loadContract(this.contractAddress);
      
      return true;
    } catch (error) {
      console.error('Failed to initialize Web3:', error);
      this.isInitialized = false;
      throw error;
    }
  }

  loadContract(address) {
    if (!this.isInitialized) {
      throw new Error('Web3 not initialized');
    }
    
    try {
      this.contractAddress = address;
      this.contract = new this.web3.eth.Contract(
        TradingContractABI.abi,
        address
      );
      console.log('Contract loaded at:', address);
      return this.contract;
    } catch (error) {
      console.error('Failed to load contract:', error);
      throw error;
    }
  }

  async assignWalletToUser() {
    if (!this.isInitialized) {
      console.log('Web3 not initialized, initializing now...');
      await this.initialize();
    }
    
    try {
      console.log('Requesting a prefunded wallet via backend...');
      
      const response = await fetch('http://localhost:8080/auth', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'Accept': 'application/json'
        }
      });
      
      if (!response.ok) {
        const errorText = await response.text();
        console.error('Server error response:', response.status, errorText);
        throw new Error(`Failed to get wallet from backend: ${response.status} ${response.statusText}`);
      }
      
      const data = await response.json();
      console.log('Received wallet from backend:', data);
      
      if (!data.address) {
        throw new Error('Wallet data is missing address property');
      }
      
      // Save the account for future transactions
      this.account = data.address;
      
      // Set as default account for web3 transactions
      this.web3.eth.defaultAccount = this.account;
      
      // Store in localStorage for persistence across refreshes
      localStorage.setItem('userWalletAddress', this.account);
      
      return {
        address: this.account,
        balance: data.balance || '0'
      };
    } catch (error) {
      console.error('Failed to assign wallet:', error);
      throw error;
    }
  }

  async getEthBalance(address = null) {
    if (!this.isInitialized) await this.initialize();
    
    try {
      const targetAddress = address || this.account;
      if (!targetAddress) throw new Error('No wallet address provided');
      
      const balance = await this.web3.eth.getBalance(targetAddress);
      return this.web3.utils.fromWei(balance, 'ether');
    } catch (error) {
      console.error('Failed to get ETH balance:', error);
      throw error;
    }
  }

  async getTokenBalance(tokenSymbol, address = null) {
    if (!this.isInitialized) await this.initialize();
    if (!this.contract) throw new Error('Contract not initialized');
    
    try {
      const targetAddress = address || this.account;
      if (!targetAddress) throw new Error('No wallet address provided');
      
      const balance = await this.contract.methods.getBalance(targetAddress, tokenSymbol).call();
      return balance;
    } catch (error) {
      console.error(`Failed to get ${tokenSymbol} balance:`, error);
      throw error;
    }
  }

  async executeSwap(fromToken, toToken, amount, targetAddress = null) {
    if (!this.isInitialized) await this.initialize();
    if (!this.contract) throw new Error('Contract not initialized');
    if (!this.account) throw new Error('No wallet connected');
    
    try {
      const recipient = targetAddress || this.account;
      
      // Execute the swap through the contract
      const transaction = await this.contract.methods
        .executeTrade(recipient, fromToken, toToken, amount)
        .send({ from: this.account });
      
      return transaction;
    } catch (error) {
      console.error('Swap execution failed:', error);
      throw error;
    }
  }

  async createLimitOrder(fromToken, toToken, amount, rate) {
    if (!this.isInitialized) await this.initialize();
    if (!this.contract) throw new Error('Contract not initialized');
    if (!this.account) throw new Error('No wallet connected');
    
    try {
      // Create a limit order through the contract
      const transaction = await this.contract.methods
        .createLimitOrder(fromToken, toToken, amount, rate)
        .send({ from: this.account });
      
      return transaction;
    } catch (error) {
      console.error('Limit order creation failed:', error);
      throw error;
    }
  }
  
  // Add any other blockchain interaction methods your application needs
}

// Create and export a singleton instance
const web3Service = new Web3Service();
export default web3Service; 