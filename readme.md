# Bitcoin wallet v2 

Bitcoin wallt version 2 for golang programming language 

check out avaiable methods in pacakge 

install using `go get -u github.com/ranjbar-dev/bitcoin-wallet@2.0.6`

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


## Block 

In the Bitcoin blockchain, a block is a collection of transactions that have been confirmed and recorded on the blockchain.


### Current block number 

fetch current block number using explorer configured in package 

```

blockNumber, _ := bitcoinwallet.FetchCurrentBlockNumber()
fmt.Println("blockNumber:", blockNumber)

```

### Current block hash 

fetch current block hash using explorer configured in package 

```

blockHash, _ := bitcoinwallet.FetchCurrentBlockHash()
fmt.Println("blockHash:", blockHash)

```

### Get block by number 

fetch block data by block number using explorer configured in package 

```

blockData, _ := bitcoinwallet.FetchBlockByNumber(blockNumber)
fmt.Println("blockData:", blockData)

```


## Transaction 

### create / sign / broadcast transaction 

```

var inputs models.TransactionInput

// add transaction inputs
inputs[0] = models.NewTransactionInput(privateKey, utxoValue, utxoIndex, utxoTxId)

var outputs models.TransactionOutput

// add transaction outputs
outputs[0] = models.NewTransactionOutput(toAddressBytes, valueInSatoshi)

// create a new transaction
transaction := bitcoinwallet.NewTransaction(inputs, outputs)

// get transaction size in bytes
size := transaction.Size()
fmt.Println("size of transaction in bytes: ", size)

// get transaction inputs
fmt.Println("transaction inputs: ", transaction.Inputs())

// get transaction outputs
fmt.Println("transaction outputs: ", transaction.Outputs())

// calculate transaction fee based on inputs and outputs
fmt.Println("transaction fee: ", transaction.Fee())

// sign transaction
err := transaction.SignAndSerialize()
if err != nil {

    fmt.Println("error on sign and serialize transaction: ", err)
    return
}

// broadcast transaction into blockchain
txID, err := transaction.Broadcast()
if err != nil {

    fmt.Println("error on broadcast transaction: ", err)
    return
}

fmt.Println("transaction broadcasted successfully with txID: ", txID)

```

## TODOS 

- add tests 

- replace int and int64 with big.Int for balance and satoshi transfer usages 
