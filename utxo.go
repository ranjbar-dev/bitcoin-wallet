package bitcoinwallet

import (
	"github.com/ranjbar-dev/bitcoin-wallet/models"
)

func FetchAddressUTXOs(address string) ([]models.UTXO, error) {

	v, err := config.Explorer.GetAddressUTXOs(address, config.Timeout)

	if err != nil {

		return nil, err
	}

	return v, nil
}
