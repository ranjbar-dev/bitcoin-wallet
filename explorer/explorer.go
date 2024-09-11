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

func NewTrezorExplorer() TrezorExplorer {

	return TrezorExplorer{
		baseURL: "https://btc1.trezor.io/api/v2",
	}
}

func NewBlockdaemonExplorer(network string, apiKey string) BlockdaemonExplorer {

	return BlockdaemonExplorer{
		Network: network,
		baseURL: "https://svc.blockdaemon.com/universal/v1/bitcoin",
		ApiKey:  apiKey,
	}
}
