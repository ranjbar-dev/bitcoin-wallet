package crawler

type Crawler interface {
	GetBTCPrice() (float64, error)
}

func NewCoinGeckoCraweler() CoinGeckoCrawer {

	return CoinGeckoCrawer{
		BaseURL: "https://api.coingecko.com/api/v3",
	}
}

func NewBinanceCraweler() BinanceCrawler {

	return BinanceCrawler{
		BaseURL: "https://api.binance.com/api/v3",
	}
}
