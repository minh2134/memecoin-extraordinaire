// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Trading {
    // Events for logging transactions
    event TradeExecuted(
        uint256 id,
        address sourceAddress,
        address targetAddress,
        string sourceCurr,
        string targetCurr,
        uint256 sourceAmount,
        uint256 targetAmount
    );

    // Struct to represent a currency
    struct Currency {
        string name;
        uint256 totalSupply;
        bool exists;
    }

    // Mapping from currency name to Currency struct
    mapping(string => Currency) public currencies;
    
    // Mapping from address to balances for each currency
    mapping(address => mapping(string => uint256)) public balances;
    
    // Transaction counter for IDs
    uint256 private transactionCount = 0;

    // Constructor to initialize with some currencies
    constructor() {
        // Add some initial currencies
        addCurrency("MEME", 1000000);
        addCurrency("ETH", 1000000);
        addCurrency("BTC", 1000000);
        addCurrency("USDT", 1000000);
    }

    // Function to add a new currency
    function addCurrency(string memory name, uint256 initialSupply) public {
        require(!currencies[name].exists, "Currency already exists");
        currencies[name] = Currency(name, initialSupply, true);
    }

    // Function to mint tokens to an address (for testing)
    function mintTokens(address to, string memory currencyName, uint256 amount) public {
        require(currencies[currencyName].exists, "Currency does not exist");
        balances[to][currencyName] += amount;
    }

    // Function to execute a trade
    function executeTrade(
        address targetAddress,
        string memory sourceCurr,
        string memory targetCurr,
        uint256 sourceAmount,
        uint256 targetAmount
    ) public returns (uint256) {
        // Validate currencies exist
        require(currencies[sourceCurr].exists, "Source currency does not exist");
        require(currencies[targetCurr].exists, "Target currency does not exist");
        
        // Check if sender has enough balance
        require(balances[msg.sender][sourceCurr] >= sourceAmount, "Insufficient balance");
        
        // Check if target has enough balance
        require(balances[targetAddress][targetCurr] >= targetAmount, "Target has insufficient balance");

        // Update balances
        balances[msg.sender][sourceCurr] -= sourceAmount;
        balances[msg.sender][targetCurr] += targetAmount;
        balances[targetAddress][sourceCurr] += sourceAmount;
        balances[targetAddress][targetCurr] -= targetAmount;

        // Generate transaction ID
        uint256 tradeId = transactionCount++;

        // Emit trade event
        emit TradeExecuted(
            tradeId,
            msg.sender,
            targetAddress,
            sourceCurr,
            targetCurr,
            sourceAmount,
            targetAmount
        );

        return tradeId;
    }

    // Function to get balance of an address for a specific currency
    function getBalance(address account, string memory currencyName) public view returns (uint256) {
        require(currencies[currencyName].exists, "Currency does not exist");
        return balances[account][currencyName];
    }
}