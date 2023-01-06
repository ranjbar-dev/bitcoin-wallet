package main

import (
	"fmt"
	bitcoinWallet "github.com/ranjbar-dev/bitcoin-wallet"
	"github.com/ranjbar-dev/bitcoin-wallet/enums"
)

func main() {

	c := bitcoinWallet.Crawler{
		Node: enums.MAIN_NODE,
		Addresses: []string{
			"bc1q7lcpsenzj7l9dkm9arayp4fmpr7ggwt9a8ztf8",
			"bc1q4s973affke7lr8vmqfnvye4m0gepjt40csey4r",
		},
	}

	fmt.Println(c.ScanBlocks(30))
}
