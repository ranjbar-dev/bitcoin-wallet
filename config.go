package bitcoinwallet

var config *Config

// TODO : add configs here, we can update these config using SetConfig function
type Config struct {
	Explorer string // trezor, blockdaemon, etc
	ApiKey   string // blockdaemon api key
}

func init() {

	config = &Config{}
}

func SetConfig(c *Config) {

	config = c
}
