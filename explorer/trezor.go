package explorer

import (
	"encoding/json"
	"fmt"
	"strconv"

	httpclient "github.com/ranjbar-dev/bitcoin-wallet/httpClient"
	"github.com/ranjbar-dev/bitcoin-wallet/models"
)

type TrezorExplorer struct {
	BaseURL string
}

type TrezorBalanceResponse struct {
	Balance string `json:"balance"`
}

type TrezorBlockResponse struct {
	Hash              string     `json:"hash"`
	Page              int        `json:"page"`
	TotalPages        int        `json:"totalPages"`
	PreviousBlockHash string     `json:"previousBlockHash"`
	NextBlockHash     string     `json:"nextBlockHash"`
	Height            int        `json:"height"`
	Confirmations     int        `json:"confirmations"`
	Size              int        `json:"size"`
	Time              int        `json:"time"`
	Version           int        `json:"version"`
	MerkleRoot        string     `json:"merkleRoot"`
	Nonce             string     `json:"nonce"`
	Bits              string     `json:"bits"`
	Difficulty        string     `json:"difficulty"`
	TxCount           int        `json:"txCount"`
	Txs               []BlockTxs `json:"txs"`
}

type BlockTxs struct {
	TxID string `json:"txid"`
	Vin  []struct {
		N         int      `json:"n"`
		Addresses []string `json:"addresses"`
		TxID      string   `json:"txid"`
		Value     string   `json:"value"`
	}
	Vout []struct {
		N         int      `json:"n"`
		Value     string   `json:"value"`
		Addresses []string `json:"addresses"`
		IsAddress bool     `json:"isAddress"`
	}
	BlockHash     string `json:"blockHash"`
	BlockHeight   int    `json:"blockHeight"`
	Confirmations int    `json:"confirmations"`
	BlockTime     int    `json:"blockTime"`
	Value         string `json:"value"`
	ValueIn       string `json:"valueIn"`
	Fees          string `json:"fees"`
}

type TrezorBlockHashResponse struct {
	BlockHash string `json:"blockHash"`
}

type TrezorCurrentBlockResponse struct {
	BlockBook BlockBook `json:"blockbook"`
}

type BlockBook struct {
	BestHeight int `json:"bestHeight"`
}

type TrezorUTXOsResponse struct {
	TxID          string `json:"txid"`
	Value         string `json:"value"`
	Confirmations int    `json:"confirmations"`
	Vout          int    `json:"vout"`
	Height        int    `json:"height"`
}

type TrezorBroadcastTransactionResponse struct {
	Result string `json:"result"`
}

func (e *TrezorExplorer) SetBaseURL(url string) {

	e.BaseURL = url
}

func (e *TrezorExplorer) GetAddressBalance(address string) (int, error) {

	url := fmt.Sprintf("%s/address/%s", e.BaseURL, address)

	client := httpclient.NewHttpclient()

	res, err := client.NewRequest().Get(url)
	if err != nil {

		return -1, err
	}

	var v TrezorBalanceResponse
	err = json.Unmarshal(res.Body(), &v)
	if err != nil {

		return -1, err
	}

	balance, err := strconv.ParseInt(v.Balance, 10, 64)
	if err != nil {

		return -1, err
	}

	return int(balance), nil
}

func (e *TrezorExplorer) GetCurrentBlockNumber() (int, error) {

	client := httpclient.NewHttpclient()

	res, err := client.NewRequest().Get(e.BaseURL)

	if err != nil {

		return -1, err
	}

	var v TrezorCurrentBlockResponse

	err = json.Unmarshal(res.Body(), &v)

	if err != nil {

		return -1, err
	}

	return v.BlockBook.BestHeight, nil
}

func (e *TrezorExplorer) GetCurrentBlockHash() (string, error) {

	url := fmt.Sprintf("%s/block-index/", e.BaseURL)

	client := httpclient.NewHttpclient()

	res, err := client.NewRequest().Get(url)

	if err != nil {

		return "", err
	}

	var v TrezorBlockHashResponse
	err = json.Unmarshal(res.Body(), &v)
	if err != nil {

		return "", err
	}

	return v.BlockHash, nil
}

func (e *TrezorExplorer) GetBlockByNumber(num int) (models.Block, error) {

	url := fmt.Sprintf("%s/block/%d", e.BaseURL, num)

	client := httpclient.NewHttpclient()

	res, err := client.NewRequest().Get(url)
	if err != nil {

		return models.Block{}, err
	}

	var v TrezorBlockResponse
	err = json.Unmarshal(res.Body(), &v)
	if err != nil {

		return models.Block{}, err
	}

	var txs []models.Transaction
	for _, tx := range v.Txs {

		var inputs []models.Input
		for _, vin := range tx.Vin {

			val, err := strconv.ParseInt(vin.Value, 10, 64)
			if err != nil {

				return models.Block{}, err
			}

			address := ""
			if len(vin.Addresses) > 0 {

				address = vin.Addresses[0]
			}

			inputs = append(inputs, models.Input{
				Address: address,
				Value:   int(val),
				Index:   vin.N,
				TxID:    tx.TxID,
			})
		}

		var outputs []models.Output
		for _, vout := range tx.Vout {

			val, err := strconv.ParseInt(vout.Value, 10, 64)
			if err != nil {

				return models.Block{}, err
			}

			address := ""
			if len(vout.Addresses) > 0 {

				address = vout.Addresses[0]
			}

			outputs = append(outputs, models.Output{
				Address: address,
				Value:   int(val),
				Index:   vout.N,
			})
		}

		txs = append(txs, models.Transaction{
			TxID:          tx.TxID,
			BlockHash:     tx.BlockHash,
			Confirmations: tx.Confirmations,
			Inputs:        inputs,
			Outputs:       outputs,
		})
	}

	return models.Block{
		Hash:              v.Hash,
		PreviousBlockHash: v.PreviousBlockHash,
		NextBlockHash:     v.NextBlockHash,
		Height:            v.Height,
		Confirmations:     v.Confirmations,
		TxCount:           v.TxCount,
		Txs:               txs,
	}, nil
}

func (e *TrezorExplorer) GetAddressUTXOs(address string, timeOut int) ([]models.UTXO, error) {

	if timeOut == 0 {
		timeOut = 30
	}

	url := fmt.Sprintf("%s/utxo/%s?confirmed=true", e.BaseURL, address)

	client := httpclient.NewHttpclient()

	res, err := client.NewRequest().Get(url)

	if err != nil {

		return nil, err
	}

	var v []TrezorUTXOsResponse
	err = json.Unmarshal(res.Body(), &v)
	if err != nil {

		return nil, err
	}

	utxos := []models.UTXO{}
	for _, v := range v {

		val, err := strconv.ParseInt(v.Value, 10, 64)

		if err != nil {

			return nil, err
		}

		utxos = append(utxos, models.UTXO{
			Amount:        val,
			TxID:          v.TxID,
			Index:         uint32(v.Vout),
			Confirmations: int(v.Confirmations),
		})
	}

	return utxos, nil
}

func (e *TrezorExplorer) GetTransactionByTxID(txID string) (models.Transaction, error) {

	url := fmt.Sprintf("%s/tx/%s", e.BaseURL, txID)

	client := httpclient.NewHttpclient()

	res, err := client.NewRequest().Get(url)

	if err != nil {

		return models.Transaction{}, err
	}

	var v BlockTxs
	err = json.Unmarshal(res.Body(), &v)
	if err != nil {

		return models.Transaction{}, err
	}

	var inputs []models.Input
	for _, vin := range v.Vin {
		val, err := strconv.ParseInt(vin.Value, 10, 64)
		if err != nil {
			fmt.Println(err)
			return models.Transaction{}, err
		}
		inputs = append(inputs, models.Input{
			Address: vin.Addresses[0],
			Value:   int(val),
			Index:   vin.N,
			TxID:    v.TxID,
		})
	}

	var outputs []models.Output
	for _, vout := range v.Vout {

		val, err := strconv.ParseInt(vout.Value, 10, 64)
		if err != nil {

			return models.Transaction{}, err
		}

		outputs = append(outputs, models.Output{
			Address: vout.Addresses[0],
			Value:   int(val),
			Index:   vout.N,
		})
	}

	return models.Transaction{
		TxID:          v.TxID,
		BlockHash:     v.BlockHash,
		Confirmations: v.Confirmations,
		Inputs:        inputs,
		Outputs:       outputs,
	}, nil
}

func (e *TrezorExplorer) BroadcastTransaction(hex string) (string, error) {

	url := fmt.Sprintf("%s/tx/%s", e.BaseURL, hex)

	client := httpclient.NewHttpclient()

	res, err := client.NewRequest().Get(url)

	if err != nil {

		return "", err
	}

	var v TrezorBroadcastTransactionResponse
	err = json.Unmarshal(res.Body(), &v)
	if err != nil {

		return "", err
	}

	return v.Result, nil
}
