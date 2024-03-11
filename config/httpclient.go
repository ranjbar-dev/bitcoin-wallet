package config

func GetHttpClientAgent() string {
	return readConfigString("BTC_HTTP_CLIENT_AGENT", "BitcoinWallet")
}

func GetHttpClientDebug() bool {
	return readConfigInt("BTC_HTTP_CLIENT_DEBUG", 0) == 1
}

func GetHttpClientTimeout() int {
	return readConfigInt("BTC_HTTP_CLIENT_TIMEOUT", 60)
}
