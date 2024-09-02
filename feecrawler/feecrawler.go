package feecrawler

type Crawler interface {
	GetEstimatedFee() (float64, float64, float64, error)
}

func NewBitcoinerCrawler() BitCoinerCrawler {

	return BitCoinerCrawler{
		BaseURL: "https://bitcoiner.live/api",
	}
}

func NewBlockstreamCrawler() BlockstreamCrawler {

	return BlockstreamCrawler{
		BaseURL: "https://blockstream.info/api",
	}
}

func NewBlockchainCrawler() BlockchainCrawler {

	return BlockchainCrawler{
		BaseURL: "https://api.blockchain.info",
	}
}
