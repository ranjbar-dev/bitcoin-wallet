package explorer

import (
	"encoding/json"
	"fmt"
	"strconv"

	httpclient "github.com/ranjbar-dev/bitcoin-wallet/httpClient"
	"github.com/ranjbar-dev/bitcoin-wallet/models"
)

type TrezorExplorer struct {
	baseURL string
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

func (e TrezorExplorer) GetAddressBalance(address string) (int, error) {

	url := fmt.Sprintf("%s/address/%s", e.baseURL, address)

	client := httpclient.NewHttpclient()

	res, err := client.NewRequest().Get(url)

	if err != nil {
		fmt.Println("Error", err)
		return -1, err
	}

	var v TrezorBalanceResponse

	err = json.Unmarshal(res.Body(), &v)

	if err != nil {
		fmt.Println(err)
	}

	balance, err := strconv.ParseInt(v.Balance, 10, 64)

	if err != nil {
		fmt.Println(err)
		return -1, err
	}
	return int(balance), nil
}

func (e TrezorExplorer) GetCurrentBlockNumber() (int, error) {

	client := httpclient.NewHttpclient()

	res, err := client.NewRequest().Get(e.baseURL)

	if err != nil {
		fmt.Println("Error", err)
		return -1, err
	}

	var v TrezorCurrentBlockResponse

	err = json.Unmarshal(res.Body(), &v)

	if err != nil {
		fmt.Println(err)
	}

	return v.BlockBook.BestHeight, nil
}

func (e TrezorExplorer) GetCurrentBlockHash() (string, error) {

	url := fmt.Sprintf("%s/block-index/", e.baseURL)

	client := httpclient.NewHttpclient()

	res, err := client.NewRequest().Get(url)

	if err != nil {
		fmt.Println("Error", err)
		return "", err
	}

	var v TrezorBlockHashResponse

	err = json.Unmarshal(res.Body(), &v)

	if err != nil {
		fmt.Println(err)
	}

	return v.BlockHash, nil
}

func (e TrezorExplorer) GetBlockByNumber(num int) (models.Block, error) {

	url := fmt.Sprintf("%s/block/%d", e.baseURL, num)

	client := httpclient.NewHttpclient()

	res, err := client.NewRequest().Get(url)

	if err != nil {
		fmt.Println("Error", err)
		return models.Block{}, err
	}

	var v TrezorBlockResponse

	err = json.Unmarshal(res.Body(), &v)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(v.Txs[2])

	var txs []models.Transaction

	for _, tx := range v.Txs {
		var inputs []models.Input
		for _, vin := range tx.Vin {
			val, err := strconv.ParseInt(vin.Value, 10, 64)
			if err != nil {
				fmt.Println(err)
				return models.Block{}, err
			}
			inputs = append(inputs, models.Input{
				Address: vin.Addresses[0],
				Value:   int(val),
				Index:   vin.N,
				TxID:    tx.TxID,
			})
		}

		var outputs []models.Output

		for _, vout := range tx.Vout {
			val, err := strconv.ParseInt(vout.Value, 10, 64)

			if err != nil {
				fmt.Println(err)
				return models.Block{}, err
			}

			outputs = append(outputs, models.Output{
				Address: vout.Addresses[0],
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

func (e TrezorExplorer) GetAddressUTXOs(address string, timeOut int) ([]models.UTXO, error) {

	if timeOut == 0 {
		timeOut = 30
	}

	url := fmt.Sprintf("%s/utxo/%s?confirmed=true", e.baseURL, address)

	client := httpclient.NewHttpclient()

	res, err := client.NewRequest().Get(url)

	if err != nil {
		fmt.Println("Error", err)
		return nil, err
	}

	var v []TrezorUTXOsResponse

	err = json.Unmarshal(res.Body(), &v)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	utxos := []models.UTXO{}

	for _, v := range v {
		val, err := strconv.ParseInt(v.Value, 10, 64)

		if err != nil {
			fmt.Println(err)
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
