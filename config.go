package bitcoinwallet

import (
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/ranjbar-dev/bitcoin-wallet/explorer"
	"github.com/ranjbar-dev/bitcoin-wallet/feecrawler"
	"github.com/ranjbar-dev/bitcoin-wallet/pricecrawler"
)

var config *Config

// TODO : add configs here, we can update these config using SetConfig function
type Config struct {
	Timeout      int                  // timeout for http requests
	Explorer     explorer.Explorer    // trezor, blockdaemon, etc
	PriceCrawler pricecrawler.Crawler // binance, coingecko, etc
	FeeCrawler   feecrawler.Crawler   // binance, coingecko, etc
	Chaincfg     *chaincfg.Params     // mainnet, testnet
}

func init() {

	config = &Config{}
}

// SetConfig sets the global configuration for the application.
// It takes a pointer to a Config struct as an argument and assigns it to the global config variable.
//
// Parameters:
//   - c: A pointer to a Config struct that holds the configuration settings.
func SetConfig(c *Config) {

	config = c
}
