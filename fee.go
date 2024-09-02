package bitcoinwallet

import "fmt"

// return Low,  Medium, High, fee Satoshi/Byte
func FetchEstimateFee() (float64, float64, float64, error) {

	low, med, high, err := config.FeeCrawler.GetEstimatedFee()

	if err != nil {
		fmt.Println(err)
		return 0, 0, 0, err
	}

	return low, med, high, nil
}
