# coinfo
Open cryptocurrency price info API. https://coinfo.wanglei.io
## Usage
request
```
get https://coinfo.wanglei.io/?coin=bitcoin
```
response
```
{
    "Symbol": "BTC",
    "Price": "7421.63",
    "Source": "CoinMarketCap",
    "UpdateAt": "2018-06-05T15:24:36+08:00",
    "PercentChange1h": -0.42,
    "PercentChange24h": -2.76
}
```