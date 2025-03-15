import Web3 from 'web3';
import TradingContractABI from '../contracts/Trading.json'; // This will be generated after compiling the contract

class Web3Service {
  constructor() {
    this.web3 = null;
    this.contract = null;
    this.account = null;
    this.isInitialized = false;
  }

  async initialize() {
    try {
      // Connect to Ganache
      this.web3 = new Web3('http://localhost:7545');
      this.isInitialized = true;
    } catch (error) {
      console.error('Failed to initialize Web3:', error);
      throw error;
    }
  }

  async createMockWallet() {
    if (!this.isInitialized) await this.initialize();
    
    try {
      // Generate a new account
      const newAccount = this.web3.eth.accounts.create();
      
      // Add the private key to the wallet
      this.web3.eth.accounts.wallet.add(newAccount);
      
      // Set as the default account
      this.account = newAccount.address;
      
      // Fund the account with mock ETH (in a real environment, this would require a faucet)
      // For Ganache, we can transfer from one of the pre-funded accounts
      const accounts = await this.web3.eth.getAccounts();
      if (accounts.length > 0) {
        await this.web3.eth.sendTransaction({
          from: accounts[0],
          to: newAccount.address,
          value: this.web3.utils.toWei('1', 'ether')
        });
      }
      
      // Mint some tokens for testing
      await this.mintTokens(newAccount.address, "MEME", 1000);
      await this.mintTokens(newAccount.address, "ETH", 10);
      await this.mintTokens(newAccount.address, "BTC", 5);
      await this.mintTokens(newAccount.address, "USDT", 2000);
      
      return {
        address: newAccount.address,
        privateKey: newAccount.privateKey
      };
    } catch (error) {
      console.error('Failed to create wallet:', error);
      throw error;
    }
  }

  async connectWallet() {
    if (!this.isInitialized) await this.initialize();
    
    try {
      // For a mock wallet, we can either:
      // 1. Use an existing account from Ganache
      const accounts = await this.web3.eth.getAccounts();
      if (accounts.length > 0) {
        this.account = accounts[0];
        return this.account;
      }
      
      // 2. Or create a new account
      return (await this.createMockWallet()).address;
    } catch (error) {
      console.error('Failed to connect wallet:', error);
      throw error;
    }
  }

  async getBalance(address, currency) {
    if (!this.isInitialized) await this.initialize();
    if (!address) address = this.account;
    
    try {
      const balance = await this.contract.methods.getBalance(address, currency).call();
      return balance;
    } catch (error) {
      console.error(`Failed to get ${currency} balance:`, error);
      return 0;
    }
  }

  async mintTokens(address, currency, amount) {
    if (!this.isInitialized) await this.initialize();
    if (!address) address = this.account;
    
    try {
      const accounts = await this.web3.eth.getAccounts();
      await this.contract.methods.mintTokens(address, currency, amount).send({ from: accounts[0] });
      return true;
    } catch (error) {
      console.error(`Failed to mint ${currency} tokens:`, error);
      return false;
    }
  }

  async executeTrade(targetAddress, sourceCurr, targetCurr, sourceAmount, targetAmount) {
    if (!this.isInitialized) await this.initialize();
    if (!this.account) throw new Error('No wallet connected');
    
    try {
      const result = await this.contract.methods.executeTrade(
        targetAddress,
        sourceCurr,
        targetCurr,
        sourceAmount,
        targetAmount
      ).send({ from: this.account });
      
      // Extract the trade ID from the event
      const tradeId = result.events.TradeExecuted.returnValues.id;
      
      // Return transaction details in the format specified
      return {
        id: tradeId,
        sourceAddress: this.account,
        targetAddress: targetAddress,
        sourceCurr: sourceCurr,
        targetCurr: targetCurr,
        sourceAmount: sourceAmount,
        targetAmount: targetAmount
      };
    } catch (error) {
      console.error('Trade execution failed:', error);
      throw error;
    }
  }

  async assignWalletToUser() {
    if (!this.isInitialized) await this.initialize();
    
    try {
      // Get all accounts from Ganache
      const accounts = await this.web3.eth.getAccounts();
      if (accounts.length === 0) {
        throw new Error('No Ganache accounts available');
      }
      
      // Select a random account
      const randomIndex = Math.floor(Math.random() * accounts.length);
      this.account = accounts[randomIndex];
      
      // Get the balance
      const balance = await this.web3.eth.getBalance(this.account);
      const ethBalance = this.web3.utils.fromWei(balance, 'ether');
      
      return {
        address: this.account,
        balance: ethBalance
      };
    } catch (error) {
      console.error('Failed to assign wallet:', error);
      throw error;
    }
  }
}

const web3Service = new Web3Service();
export default web3Service; 