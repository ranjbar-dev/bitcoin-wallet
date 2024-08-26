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
		N     int    `json:"n"`
		Value string `json:"value"`
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

func (e TrezorExplorer) GetBlockByNumber(num int) (models.BlockChainBlock, error) {

	url := fmt.Sprintf("%s/block/%d", e.baseURL, num)

	client := httpclient.NewHttpclient()

	res, err := client.NewRequest().Get(url)

	if err != nil {
		fmt.Println("Error", err)
		return models.BlockChainBlock{}, err
	}

	var v TrezorBlockResponse

	err = json.Unmarshal(res.Body(), &v)

	if err != nil {
		fmt.Println(err)
	}

	return models.BlockChainBlock{
		Hash:              v.Hash,
		Page:              v.Page,
		TotalPages:        v.TotalPages,
		PreviousBlockHash: v.PreviousBlockHash,
		NextBlockHash:     v.NextBlockHash,
		Height:            v.Height,
		Confirmations:     v.Confirmations,
		Size:              v.Size,
		Time:              v.Time,
		Version:           v.Version,
		MerkleRoot:        v.MerkleRoot,
		Nonce:             v.Nonce,
		Bits:              v.Bits,
		Difficulty:        v.Difficulty,
		TxCount:           v.TxCount,
		Txs:               []models.BlockTxs{},
	}, nil
}
