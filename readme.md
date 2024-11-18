# Bitcoin wallet v2 

### Config 

you should config wallet first 

```

bitcoinwallet.SetConfig(bitcoinwallet.Config{
    Timeout: 30, // http client timeout in second 
    Explorer: explorer.TrezorExplorer{ BaseUrl: "https://bc1.trezor.io" }, // exporer to get blockchain info, check ./exporer folder for other options 
    PriceCrawler: pricecrawler.BinanceCrawler{ BaseURL: "https://eapi.binance.com" }, // price crawler to get current price data, check ./pricecrawler folder for other options  
    FeeCrawler: feecrawler.BlockstreamCrawler{ BaseURL: "https://blockstream.info/api" }, // fee crawler to get current fee data, check ./feecrawler folder for other options 
    Chaincfg: chaincfg.MainNetParams, // Bitcoin chain config 
})

```