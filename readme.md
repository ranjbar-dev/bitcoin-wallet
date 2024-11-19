# Bitcoin wallet v2 

Bitcoin wallt version 2 for golang programming language 

check out avaiable methods in pacakge 

## Configuration 

#### Set config 

```

// configure package 
bitcoinwallet.SetConfig(&bitcoinwallet.Config{
    Timeout:      30,                                 // http client timeout in second
    Explorer:     explorer.NewTrezorExplorer(),       // exporer to get blockchain info, check ./exporer folder for other options
    PriceCrawler: pricecrawler.NewBinanceCraweler(),  // price crawler to get current price data, check ./pricecrawler folder for other options
    FeeCrawler:   feecrawler.NewBlockstreamCrawler(), // fee crawler to get current fee data, check ./feecrawler folder for other options
    Chaincfg:     &chaincfg.MainNetParams,            // Bitcoin chain config
})

```

## Private key  

#### Generate new private key 

```

privateKey, _ := bitcoinwallet.GeneratePrivateKey()
fmt.Println("*ecdsa.PrivateKey:", privateKey)

```

#### Private key to hex string

```

hex, _ := bitcoinwallet.PrivateKeyToHex(privateKey)
fmt.Println("private key hex:", hex)

```

#### Private key to bytes 

```

bytes, _ := bitcoinwallet.PrivateKeyToBytes(privateKey)
fmt.Println("private key bytes:", bytes)

```

#### Hex string to private key 

```

privateKey, _ := bitcoinwallet.BytesToPrivateKey(bytes)
fmt.Println("*ecdsa.PrivateKey:", bytes)

```

#### Bytes to private key 

```

privateKey, _ := bitcoinwallet.HexToPrivateKey(hex)
fmt.Println("*ecdsa.PrivateKey:", bytes)

```

## Address   

#### Private key to address 

```

address, _ := bitcoinwallet.PrivateKeyToAddress(privateKey)
fmt.Println("Address:", address)

```

#### Address balance  

retrieve address balacne from explorer configured in package

```

balance, _ := bitcoinwallet.FetchAddressBalance(address)
fmt.Println("Balance in satoshi:", balance)

```


#### Address UTXO's

retrieve address utxo from explorer configured in package 

```

utxos, _ := bitcoinwallet.FetchAddressUTXOs(address)
fmt.Println("UTXO records:", utxos)

```


## Utils 

#### Estimate fee 

retrieve blockchain high, mid and low fee in sat/vbyte using fee crawler configured in package 

```

low, mid, high, err := bitcoinwallet.FetchEstimateFee()

```

#### Current BTC price 

retrieve BTC price in USD using price crawler configured in package 

```

price, err := bitcoinwallet.FetchPrice()

```

// TODO : complete Transaction and Block doc 
// TODO : add tests 
// TODO : replace int and int64 with big.Int for balance and satoshi transfer usages 
