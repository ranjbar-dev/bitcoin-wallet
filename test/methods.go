package test

import (
	bitcoinWallet "github.com/Amirilidan78/bitcoin-wallet"
	"github.com/Amirilidan78/bitcoin-wallet/enums"
)

var node = enums.TEST_NODE
var validPrivateKey = "88414dbb373a211bc157265a267f3de6a4cec210f3a5da12e89630f2c447ad27"
var invalidPrivateKey = "invalid"
var validOwnerAddress = "tb1qppv790u4dz48ctnk3p7ss7fmspckagp3wrfyp0"
var invalidOwnerAddress = "tb15111190u4dz48ctn1273333ss7fmspckag341fyp0"
var validToAddress = "tb1q0r23g66m9rhhak8aahsg53wfp5egt2huuc4tnu"
var invalidToAddress = "tb15111190u4dz48ctn1273333ss7fmspckag341fyp0"
var btcAmount = int64(10000) // 0.00001
var feeAmount = int64(10000) // 0.00001

func wallet() *bitcoinWallet.BitcoinWallet {
	w, _ := bitcoinWallet.CreateBitcoinWallet(node, validPrivateKey)
	return w
}

func crawler() *bitcoinWallet.Crawler {

	return &bitcoinWallet.Crawler{
		Node:      node,
		Addresses: []string{validOwnerAddress},
	}
}
