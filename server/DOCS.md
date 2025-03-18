# API Documentation

This file explains and documents API endpoints for the go backend.

## /auth

This API expects a {GET} request from the client
This API tells the backend to assume a random pre-funded wallet as the transactor, then return that wallet address + balance to the sender

Returns: status code `200`
```
{
    "address":      "0x5918b2e647464d4743601a865753e64C8059Dc4F",
    "balance":      10000,      /* float, in ether, inaccuracy is to be expected */
    "weiBalance":   10000..0000 /* Int, in wei, accurate balance */
}
```
## /account/balance

This API expects a {GET} request from the client
This API returns the currently assumed wallet address + balance to the sender

Returns: status code 
```
200     Successful
400     Bad Request, need to run /auth before checking balance!
```

Assume the status code is 200, the response JSON has the following schema:
```
{
    "address":      "0x5918b2e647464d4743601a865753e64C8059Dc4F",
    "balance":      10000,      /* float, in ether, inaccuracy is to be expected */
    "weiBalance":   10000..0000 /* Int, in wei, accurate balance */
}
```

## /trade/swap

This API expects a {POST} request from the client.
The POST request should contain a JSON in its body with the following schema:
```
{
    "SourceCurr":   "BTC",          /* string, one of the supported currencies */
    "TargetCurr":   "PEPE",         /* string, one of the supported currencies */
    "SourceAmount": "14.99",        /* string or float, amount willing to trade */
    "Rate":         "0.588",        /* string or float, the trading rate, sourcecurr/targetcurr */
    "Slippage":     "0.05"          /* string or float, the slippage, expressed in numbers, not percentage unit */
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

## /trade/limit

This API expects a {POST} request from the client.
The POST request should contain a JSON in its body with the following schema:
```

{
    "SourceCurr":   "BTC",          /* string, one of the supported currencies */
    "TargetCurr":   "PEPE",         /* string, one of the supported currencies */
    "SourceAmount": "14.99",        /* string or float, amount willing to trade */
    "Rate":         "0.588",        /* string or float, the trading rate, sourcecurr/targetcurr */
    "SourceAddress":"0xDEADBEEF",   /* string, source address of the client */
}
```

Return codes:
```
200: Successful
400: Bad request, check request body
404: Resources not found
```

If return code is `200` the response body should be a JSON with this schema:
```
{
    "IsMatched":    true    /* if the limit order is matched */
    "SwapDetails: {         /* if matched, contains swap details */
        "TradedAddress":    "0xDEADBEEF",       /* string, the traded wallet */
        "TradedAmount":     "14.99",            /* string, traded amount */
        "ReceivedAmount":   "25.99",            /* string, received amount */
        "FromCurr":         "BTC",              /* string, source currency */
        "ToCurr":           "PEPE"              /* string, target currency */
    }
}
                    
```
if `IsMatched` is `false`, the limit request instead will be inserted to the database

for `*Amount`, `Rate`, `Slippage` keys the string should (and expected to) be a valid real number

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
