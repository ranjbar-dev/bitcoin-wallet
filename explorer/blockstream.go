package explorer

import (
	"encoding/json"
	"fmt"
	"strconv"

	httpclient "github.com/ranjbar-dev/bitcoin-wallet/httpClient"
	"github.com/ranjbar-dev/bitcoin-wallet/models"
)

type BlockstreamExplorer struct {
	Network string
	ApiKey  string
	baseURL string
}

type BlockstreamUTXO struct {
	TxID   string `json:"txid"`
	Vout   int    `json:"vout"`
	Status struct {
		Confirmed   bool   `json:"confirmed"`
		BlockHeight int    `json:"block_height"`
		BlockHash   string `json:"block_hash"`
		BlockTime   int64  `json:"block_time"`
	} `json:"status"`
	Value int64 `json:"value"`
}

func (e *BlockstreamExplorer) SetBaseURL(url string) {

	e.baseURL = url
}

func (e *BlockstreamExplorer) GetAddressBalance(address string) (int, error) {

	headers := map[string]string{
		"Authorization": "Bearer " + e.ApiKey,
		"Accept":        "application/json",
	}

	url := fmt.Sprintf("%s/address/%s/utxo", e.baseURL, address)

	client := httpclient.NewHttpclient()

	res, err := client.NewRequest().SetHeaders(headers).Get(url)
	if err != nil {

		return -1, err
	}

	var records []BlockstreamUTXO
	err = json.Unmarshal(res.Body(), &records)
	if err != nil {

		return -1, err
	}

	var balance int

	for _, record := range records {

		balance += int(record.Value)
	}

	return balance, nil
}

func (e *BlockstreamExplorer) GetCurrentBlockNumber() (int, error) {

	headers := map[string]string{
		"Authorization": "Bearer " + e.ApiKey,
		"Accept":        "application/json",
	}

	url := fmt.Sprintf("%s/blocks/tip/height", e.baseURL)

	client := httpclient.NewHttpclient()

	res, err := client.NewRequest().SetHeaders(headers).Get(url)
	if err != nil {

		return -1, err
	}

	height, err := strconv.Atoi(string(res.Body()))
	if err != nil {

		return -1, err
	}

	return height, nil
}

func (e *BlockstreamExplorer) GetCurrentBlockHash() (string, error) {

	headers := map[string]string{
		"Authorization": "Bearer " + e.ApiKey,
		"Accept":        "application/json",
	}

	url := fmt.Sprintf("%s/blocks/tip/hash", e.baseURL)

	client := httpclient.NewHttpclient()

	res, err := client.NewRequest().SetHeaders(headers).Get(url)
	if err != nil {

		return "", err
	}

	return string(res.Body()), nil
}

func (e *BlockstreamExplorer) GetBlockByNumber(num int, withSleep bool) (models.Block, error) {

	return models.Block{}, nil
}

func (e *BlockstreamExplorer) GetAddressUTXOs(address string, timeOut int) ([]models.UTXO, error) {

	return nil, nil
}

func (e *BlockstreamExplorer) GetTransactionByTxID(txID string) (models.Transaction, error) {

	return models.Transaction{}, nil
}

func (e *BlockstreamExplorer) BroadcastTransaction(hex string) (string, error) {

	return "", nil
}

func (e *BlockstreamExplorer) getBlockHashByNumber(num int) (string, error) {

	headers := map[string]string{
		"Authorization": "Bearer " + e.ApiKey,
		"Accept":        "application/json",
	}

	url := fmt.Sprintf("%s/block-height/%d", e.baseURL, num)

	client := httpclient.NewHttpclient()

	res, err := client.NewRequest().SetHeaders(headers).Get(url)
	if err != nil {

		return "", err
	}

	return string(res.Body()), nil
}

// CLIENT_ID: bfb1e1fb-49e1-4317-bed8-602d93f16968
// CLIENT_SECRET: 2EftkBIMgwkhobNGgQSTCziMVZVl37bx
