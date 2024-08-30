package bitcoinwallet

// return High, Medium, Low fee Satoshi/Byte
func FetchEstimateFee() (float64, float64, float64, error) {

	config.FeeCrawler.GetEstimatedFee()
	// TODO : implement

	// get estimate fee using this api: https://blockstream.info/api/fee-estimates

	// https://api.blockchain.info/mempool/fees

	// https://bitcoiner.live/api/fees/estimates/latest

	return 0.0, 0.0, 0.0, nil
}
