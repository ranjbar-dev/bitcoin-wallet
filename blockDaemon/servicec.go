package blockDaemon

import (
	"github.com/Amirilidan78/bitcoin-wallet/blockDaemon/response"
	"github.com/Amirilidan78/bitcoin-wallet/httpClient"
)

type BlockDaemon interface {
	CurrentBlockNumber() (response.CurrentBlockNumberResponse, error)
	CurrentBlockHash() (response.CurrentBlockHashResponse, error)
	BlockByNumber(number int64) (response.BlockResponse, error)
	BlockByHash(hash string) (response.BlockResponse, error)
	AddressBalance(address string) (response.BalanceResponse, error)
	AddressUTXO(address string) (response.UTXOResponse, error)
	AddressTxs(address string) (response.AddressTransactionResponse, error)
	Tx(hash string) (response.Transaction, error)
	Broadcast(raw string) (response.BroadcastTransactionResponse, error)
}

type blockDaemon struct {
	conf ConfigBlockDaemon
	hc   httpClient.HttpClient
}

func NewBlockDaemonService(conf ConfigBlockDaemon) BlockDaemon {
	return &blockDaemon{conf, httpClient.NewHttpClient()}
}
