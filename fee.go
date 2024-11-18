package bitcoinwallet

import "fmt"

// FetchEstimateFee retrieves the estimated transaction fees for low, medium, and high priority transactions.
// It returns three float64 values representing the estimated fees for low, medium, and high priority transactions,
// and an error if the fee estimation fails.
//
// Returns:
//   - low (float64): The estimated fee for low priority transactions.
//   - med (float64): The estimated fee for medium priority transactions.
//   - high (float64): The estimated fee for high priority transactions.
//   - err (error): An error if the fee estimation fails, otherwise nil.
func FetchEstimateFee() (float64, float64, float64, error) {

	low, med, high, err := config.FeeCrawler.GetEstimatedFee()

	if err != nil {
		fmt.Println(err)
		return 0, 0, 0, err
	}

	return low, med, high, nil
}
