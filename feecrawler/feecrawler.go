package feecrawler

type Crawler interface {
	GetEstimatedFee() (float64, error)
}

func NewBitcoinerCrawler() BitCoinerCrawler {

	return BitCoinerCrawler{
		BaseURL: "https://bitcoiner.live/api",
	}
}

func NewBlockstreamCrawler() BlockstreamCrawler {

	return BlockstreamCrawler{}
}

func NewBlockchainCrawler() BlockchainCrawler {

	return BlockchainCrawler{}
}