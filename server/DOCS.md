# API Documentation

This file explains and documents API endpoints for the go backend.

## /trade/swap

This API expects a {POST} request from the client.

### POST 
The POST request should contain a JSON in its body with the following schema:
`SourceAmount` and `TargetAmount` is the float amount times 100 (since the amount is a float with 2 decimal place), type `Int` to prevent float inaccurate arithmetic.
```
{
    "SourceCurr":   "BTC",       /* string, one of the supported currencies */
    "TargetCurr":   "PEPE",      /* string, one of the supported currencies */
    "SourceAmount": 1499,        /* int, amount willing to trade * 100 */
    "TargetAmount": 1599,        /* int, amount willing to receive * 100 */
    "SourceAddress":"0xDEADBEEF",/* string, source address of the client */
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
    "TradedAmount":     1499,           /* Int, traded amount * 100 */
    "ReceivedAmount":   2599,           /* Int, received amount * 200 */
}
```

