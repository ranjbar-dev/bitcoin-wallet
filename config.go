package bitcoinwallet

import (
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/ranjbar-dev/bitcoin-wallet/explorer"
)

var config *Config

// TODO : add configs here, we can update these config using SetConfig function
type Config struct {
	Explorer explorer.Explorer // trezor, blockdaemon, etc
	Chaincfg *chaincfg.Params  // mainnet, testnet
}

func init() {

	config = &Config{}
}

func SetConfig(c *Config) {

	config = c
}
