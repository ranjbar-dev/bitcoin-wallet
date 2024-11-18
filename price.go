package bitcoinwallet

import "fmt"

// FetchBTCPrice retrieves the current price of Bitcoin (BTC) using the configured price crawler.
// It returns the price as a float64 and an error if there was an issue fetching the price.
func FetchBTCPrice() (float64, error) {

	price, err := config.PriceCrawler.GetBTCPrice()

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return price, nil
}
