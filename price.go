package bitcoinwallet

import "fmt"

func FetchBTCPrice() (float64, error) {

	price, err := config.PriceCrawler.GetBTCPrice()

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return price, nil
}
