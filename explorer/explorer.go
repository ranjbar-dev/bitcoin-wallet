package explorer

import "github.com/ranjbar-dev/bitcoin-wallet/models"

type Explorer interface {
	GetAddressBalance(string) (int, error)
	GetCurrentBlockNumber() (int, error)
	GetCurrentBlockHash() (string, error)
	GetBlockByNumber(int) (models.Block, error)
	GetAddressUTXOs(string, int) ([]models.UTXO, error)
	GetTransactionByTxID(string) (models.Transaction, error)
	BroadcastTransaction(string) (string, error)
}

func NewTrezorExplorer(baseUrl string) TrezorExplorer {

	return TrezorExplorer{
		baseURL: baseUrl,
	}
}

func NewBlockdaemonExplorer(network string, apiKey string, baseUrl string) BlockdaemonExplorer {

	return BlockdaemonExplorer{
		Network: network,
		baseURL: baseUrl,
		ApiKey:  apiKey,
	}
}
