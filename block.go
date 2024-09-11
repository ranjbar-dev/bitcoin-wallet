package bitcoinwallet

import (
	"fmt"

	"github.com/ranjbar-dev/bitcoin-wallet/models"
)

func FetchCurrentBlockNumber() (int, error) {

	num, err := config.Explorer.GetCurrentBlockNumber()

	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	return num, nil
}

func FetchCurrentBlockHash() (string, error) {

	hash, err := config.Explorer.GetCurrentBlockHash()

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return hash, nil
}

func FetchBlockByNumber(num int) (models.Block, error) {

	block, err := config.Explorer.GetBlockByNumber(num)

	if err != nil {
		fmt.Println(err)
		return models.Block{}, err
	}

	return block, nil
}
