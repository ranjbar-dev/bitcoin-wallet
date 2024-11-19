package bitcoinwallet

import (
	"fmt"

	"github.com/ranjbar-dev/bitcoin-wallet/models"
)

// FetchCurrentBlockNumber retrieves the current block number from the blockchain explorer.
// It returns the block number as an integer and an error if any issues occur during the fetch process.
// If an error occurs, the function prints the error and returns -1 along with the error.
func FetchCurrentBlockNumber() (int, error) {

	num, err := config.Explorer.GetCurrentBlockNumber()

	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	return num, nil
}

// FetchCurrentBlockHash retrieves the current block hash from the configured blockchain explorer.
// It returns the block hash as a string and an error if any issues occur during the retrieval process.
func FetchCurrentBlockHash() (string, error) {

	hash, err := config.Explorer.GetCurrentBlockHash()

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return hash, nil
}

// FetchBlockByNumber retrieves a block by its number from the blockchain explorer.
// It takes an integer `num` representing the block number and returns a `models.Block`
// and an error if any occurs during the retrieval process.
//
// Parameters:
//
//	num (int): The block number to fetch.
//
// Returns:
//
//	models.Block: The block corresponding to the given number.
//	error: An error object if an error occurs, otherwise nil.
func FetchBlockByNumber(num int) (models.Block, error) {

	block, err := config.Explorer.GetBlockByNumber(num)

	if err != nil {
		fmt.Println(err)
		return models.Block{}, err
	}

	return block, nil
}
