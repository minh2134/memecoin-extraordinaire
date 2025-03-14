# API Documentation

This file explains and documents API endpoints for the go backend.

## /trade/swap

This API expects a {POST} request from the client.

### POST 
The POST request should contain a JSON in its body with the following schema:
```
{
    "SourceCurr":   "BTC",          /* string, one of the supported currencies */
    "TargetCurr":   "PEPE",         /* string, one of the supported currencies */
    "SourceAmount": "14.99",        /* string or float, amount willing to trade */
    "Rate": "0.588",                /* string or float, the trading rate, sourcecurr/targetcurr */
    "SourceAddress":"0xDEADBEEF",   /* string, source address of the client */
}
```

Returns:
Return codes:
```
200: Successful
400: Bad request. Check request body
404: Resources not found
```

If the status is `200` the body response JSON has this schema:
```
{
    "TradedAddress":    "0xDEADBEEF",       /* string, the traded wallet */
    "TradedAmount":     "14.99",            /* string, traded amount */
    "ReceivedAmount":   "25.99",            /* string, received amount */
    "FromCurr":         "BTC",              /* string, source currency */
    "ToCurr":           "PEPE"              /* string, target currency */
}
```
for `*Amount` keys the string should (and expected) be a valid real number

Smart Contract Implementation with Mock Wallet using Ganache
1.	Compile and Deploy the Trading Contract:
•	Install Truffle: npm install -g truffle
•	Initiate the contract: truffle init
•	Compile the contract: truffle compile
•	Make sure Ganache is running on port 7545
•	Deploy the contract: truffle migrate
2.	Install Required Dependencies:
•	npm install --save web3
3.	Create the contract JSON file:
•	After compilation, copy the Trading.json file from the build/contracts directory to client/src/contracts/
4.	Set Up the Database Schema:
•	The transactions will follow your specified format
•	The smart contract events will emit the data in this format which can be stored in your database
This implementation will allow you to:
1.	Create mock wallets without connecting to external wallets
2.	Trade between addresses within Ganache
3.	Store transactions in the specified format
4.	Test the entire flow without requiring real cryptocurrencies
