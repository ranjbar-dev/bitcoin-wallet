package bitcoinwallet

import (
	"fmt"

	"github.com/ranjbar-dev/bitcoin-wallet/models"
)

func FetchAddressUTXOs(address string, timeOut int) ([]models.UTXO, error) {

	if timeOut == 0 {
		timeOut = 30
	}
	v, err := config.Explorer.GetAddressUTXOs(address, timeOut)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return v, nil
}
