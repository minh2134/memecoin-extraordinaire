# API Documentation

This file explains and documents API endpoints for the go backend.

## /trade/swap

This API expects a {POST} request from the client.

### POST 
The POST request should contain a JSON in its body with the following schema:
`SourceAmount` and `TargetAmount` is the float amount times 100 (since the amount is a float with 2 decimal place), type `Int` to prevent float inaccurate arithmetic.
```
{
    "SourceCurr":   "BTC",          /* string, one of the supported currencies */
    "TargetCurr":   "PEPE",         /* string, one of the supported currencies */
    "SourceAmount": "14.99",        /* string or float, amount willing to trade */
    "TargetAmount": "15.99",        /* string or float, amount willing to receive */
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
    "Address":          "0xDEADBEEF",   /* string, the traded wallet */
    "TradedAmount":     "14.99",           /* string, traded amount */
    "ReceivedAmount":   "25.99",           /* string, received amount */
}
```
for `*Amount` keys the string should (and expected) be a valid real number
