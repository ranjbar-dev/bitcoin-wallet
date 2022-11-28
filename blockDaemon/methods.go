package blockDaemon

import (
	"github.com/Amirilidan78/bitcoin-wallet/blockDaemon/response"
	"strconv"
)

func (bd *blockDaemon) CurrentBlockNumber() (response.CurrentBlockNumberResponse, error) {

	var res response.CurrentBlockNumberResponse
	path := "/sync/block_number"

	err := bd.get(path, &res)

	return res, err
}

func (bd *blockDaemon) CurrentBlockHash() (response.CurrentBlockHashResponse, error) {

	var res response.CurrentBlockHashResponse
	path := "/sync/block_id"

	err := bd.get(path, &res)

	return res, err
}

func (bd *blockDaemon) BlockByNumber(number int64) (response.BlockResponse, error) {

	var res response.BlockResponse
	path := "/block/" + strconv.FormatInt(number, 10)

	err := bd.get(path, &res)

	return res, err
}

func (bd *blockDaemon) BlockByHash(hash string) (response.BlockResponse, error) {

	var res response.BlockResponse
	path := "/block/" + hash

	err := bd.get(path, &res)

	return res, err
}

func (bd *blockDaemon) AddressBalance(address string) (response.BalanceResponse, error) {

	var res response.BalanceResponse
	path := "/account/" + address

	err := bd.get(path, &res)

	return res, err
}

func (bd *blockDaemon) AddressUTXO(address string) (response.UTXOResponse, error) {

	var res response.UTXOResponse
	path := "/account/" + address + "/utxo"

	err := bd.get(path, &res)

	return res, err
}

func (bd *blockDaemon) AddressTxs(address string) (response.AddressTransactionResponse, error) {

	var res response.AddressTransactionResponse
	path := "/account/" + address + "/txs"

	err := bd.get(path, &res)

	return res, err
}

func (bd *blockDaemon) Tx(hash string) (response.Transaction, error) {

	var res response.Transaction
	path := "/tx/" + hash

	err := bd.get(path, &res)

	return res, err
}

func (bd *blockDaemon) Broadcast(raw string) (response.BroadcastTransactionResponse, error) {

	var res response.BroadcastTransactionResponse
	path := "/tx/send"

	params := make(map[string]string)
	params["tx"] = raw

	err := bd.post(path, params, &res)

	return res, err
}
