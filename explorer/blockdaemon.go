package explorer

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	httpclient "github.com/ranjbar-dev/bitcoin-wallet/httpClient"
	"github.com/ranjbar-dev/bitcoin-wallet/models"
)

type BlockdaemonExplorer struct {
	Network string
	ApiKey  string
	baseURL string
}

type BlockdaemonBalanceResponse struct {
	ConfirmedBalance string `json:"confirmed_balance"`
	ConfirmedBlock   int    `json:"confirmed_block"`
}

type BlockdaemonUTXOsResponse struct {
	Total int    `json:"total"`
	UTXOs []UTXO `json:"data"`
	Meta  struct {
		Paging struct {
			NextPageToken string `json:"next_page_token"`
		} `json:"paging"`
	} `json:"meta"`
}

type UTXO struct {
	Status string `json:"status"`
	Value  int    `json:"value"`
	Mined  Mined  `json:"mined"`
}

type Mined struct {
	Index         int    `json:"index"`
	TxID          string `json:"tx_id"`
	Date          int    `json:"date"`
	BlockID       string `json:"block_id"`
	BlockNumber   int    `json:"block_number"`
	Confirmations int    `json:"confirmations"`
	Meta          Meta   `json:"meta"`
}

type Meta struct {
	Addresses  []string `json:"addresses"`
	Index      int      `json:"index"`
	Script     string   `json:"script"`
	ScriptType string   `json:"script_type"`
}

type BlockdaemonBlockResponse struct {
	Number   int    `json:"number"`
	ID       string `json:"id"`
	ParentID string `json:"parent_id"`
	Date     int    `json:"date"`
	NumTxs   int    `json:"num_txs"`
	Txs      []Txs  `json:"txs"`
}

type Txs struct {
	ID        string `json:"id"`
	BlockID   string `json:"block_id"`
	Date      int    `json:"date"`
	Stasus    string `json:"status"`
	NumEvents int    `json:"num_events"`
	Meta      struct {
		VSize int `json:"vsize"`
	} `json:"meta"`
	Events        []Events `json:"events"`
	Confirmations int      `json:"confirmations"`
}

type Events struct {
	ID          string `json:"id"`
	TxID        string `json:"transaction_id"`
	Type        string `json:"type"`
	Destination string `json:"destination"`
	Meta        struct {
		Addresses []string `json:"addresses"`
		Index     int      `json:"index"`
	} `json:"meta"`
	Amount  int `json:"amount"`
	Date    int `json:"date"`
	Decimal int `json:"decimal"`
}

func (e BlockdaemonExplorer) GetAddressBalance(address string) (int, error) {

	headers := map[string]string{
		"X-API-Key": e.ApiKey,
		"Accept":    "application/json",
	}

	url := fmt.Sprintf("%s/%s/account/%s", e.baseURL, e.Network, address)

	client := httpclient.NewHttpclient()

	res, err := client.NewRequest().SetHeaders(headers).Get(url)

	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	var v []BlockdaemonBalanceResponse

	err = json.Unmarshal(res.Body(), &v)

	if err != nil {
		fmt.Println(err)
	}

	balance, err := strconv.ParseInt(v[0].ConfirmedBalance, 10, 64)

	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	return int(balance), nil
}

func (e BlockdaemonExplorer) GetCurrentBlockNumber() (int, error) {

	headers := map[string]string{
		"X-API-Key": e.ApiKey,
		"Accept":    "application/json",
	}

	url := fmt.Sprintf("%s/%s/sync/block_number", e.baseURL, e.Network)

	client := httpclient.NewHttpclient()

	res, err := client.NewRequest().SetHeaders(headers).Get(url)

	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	val, err := strconv.ParseInt(res.String(), 10, 64)

	if err != nil {
		fmt.Println(err)
		return -1, err
	}
	return int(val), nil
}

func (e BlockdaemonExplorer) GetCurrentBlockHash() (string, error) {

	headers := map[string]string{
		"X-API-Key": e.ApiKey,
		"Accept":    "application/json",
	}

	url := fmt.Sprintf("%s/%s/sync/block_id", e.baseURL, e.Network)

	client := httpclient.NewHttpclient()

	res, err := client.NewRequest().SetHeaders(headers).Get(url)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	fmt.Println(res.String())

	return res.String(), nil
}

func (e BlockdaemonExplorer) GetBlockByNumber(num int) (models.Block, error) {

	headers := map[string]string{
		"X-API-Key": e.ApiKey,
		"Accept":    "application/json",
	}

	url := fmt.Sprintf("%s/%s/block/%d", e.baseURL, e.Network, num)

	client := httpclient.NewHttpclient()

	res, err := client.NewRequest().SetHeaders(headers).Get(url)

	if err != nil {
		fmt.Println(err)
		return models.Block{}, err
	}

	var v BlockdaemonBlockResponse

	err = json.Unmarshal(res.Body(), &v)

	if err != nil {
		fmt.Println("ERRR")
		fmt.Println(err)
		return models.Block{}, err
	}

	var txs []models.Transaction

	for _, tx := range v.Txs {
		var inputs []models.Input
		var outputs []models.Output

		totalIputs := 0
		totalOutputs := 0

		for _, event := range tx.Events {
			if event.Type == "utxo_input" {
				address := ""
				if len(event.Meta.Addresses) > 0 {
					address = event.Meta.Addresses[0]
				}
				inputs = append(inputs, models.Input{
					TxID:    event.TxID,
					Index:   event.Meta.Index,
					Address: address,
					Value:   event.Amount,
				})
				totalIputs += event.Amount
			}

			if event.Type == "utxo_output" {
				outputs = append(outputs, models.Output{
					Address: event.Destination,
					Value:   event.Amount,
					Index:   event.Meta.Index,
				})
				totalOutputs += event.Amount
			}
		}
		txs = append(txs, models.Transaction{
			TxID:              tx.ID,
			Confirmations:     tx.Confirmations,
			BlockHash:         tx.BlockID,
			Timestamp:         tx.Date,
			VBytes:            tx.Meta.VSize,
			Inputs:            inputs,
			Outputs:           outputs,
			TotalInputsValue:  totalIputs,
			TotalOutputsValue: totalOutputs,
			Fee:               totalIputs - totalOutputs,
		})

	}

	model := models.Block{
		Hash:              v.ID,
		PreviousBlockHash: v.ParentID,
		Height:            v.Number,
		TxCount:           v.NumTxs,
		Txs:               txs,
	}

	return model, nil
}

// timeout in seconds between requests. set to zero for default timeout of 30 seconds
func (e BlockdaemonExplorer) GetAddressUTXOs(address string, timeOut int) ([]models.UTXO, error) {

	var all BlockdaemonUTXOsResponse

	nextPageToken := ""

	for {
		path := "account/" + address + "/utxo?page_size=100&spent=false"
		if nextPageToken != "" {
			path += "&page_token=" + nextPageToken
		}

		headers := map[string]string{
			"X-API-Key": e.ApiKey,
			"Accept":    "application/json",
		}

		url := fmt.Sprintf("%s/%s/%s", e.baseURL, e.Network, path)

		client := httpclient.NewHttpclient()

		res, err := client.NewRequest().SetHeaders(headers).Get(url)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		var v BlockdaemonUTXOsResponse

		err = json.Unmarshal(res.Body(), &v)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		all.UTXOs = append(all.UTXOs, v.UTXOs...)

		if v.Meta.Paging.NextPageToken == "" {
			fmt.Println("Finished Loading, loaded ", len(all.UTXOs), " UTXOs")
			break
		}

		nextPageToken = v.Meta.Paging.NextPageToken

		fmt.Println("Loaded ", len(all.UTXOs), " UTXOs")
		fmt.Println("Loading... ")

		time.Sleep(time.Duration(timeOut) * time.Second)
	}

	var utxos []models.UTXO
	for _, datum := range all.UTXOs {
		utxos = append(utxos, models.UTXO{
			Amount:        int64(datum.Value),
			TxID:          datum.Mined.TxID,
			Index:         uint32(datum.Mined.Index),
			Confirmations: datum.Mined.Confirmations,
		})
	}

	return utxos, nil
}

func (e BlockdaemonExplorer) GetTransactionByTxID(txID string) (models.Transaction, error) {
	headers := map[string]string{
		"X-API-Key": e.ApiKey,
		"Accept":    "application/json",
	}

	url := fmt.Sprintf("%s/%s/tx/%s", e.baseURL, e.Network, txID)

	client := httpclient.NewHttpclient()

	res, err := client.NewRequest().SetHeaders(headers).Get(url)

	if err != nil {
		fmt.Println(err)
		return models.Transaction{}, err
	}

	var v Txs

	err = json.Unmarshal(res.Body(), &v)

	if err != nil {
		fmt.Println(err)
		return models.Transaction{}, err
	}

	var inputs []models.Input
	var outputs []models.Output

	totalIputs := 0
	totalOutputs := 0

	for _, event := range v.Events {
		if event.Type == "utxo_input" {
			address := ""
			if len(event.Meta.Addresses) > 0 {
				address = event.Meta.Addresses[0]
			}
			inputs = append(inputs, models.Input{
				TxID:    event.TxID,
				Index:   event.Meta.Index,
				Address: address,
				Value:   event.Amount,
			})
			totalIputs += event.Amount
		}

		if event.Type == "utxo_output" {
			outputs = append(outputs, models.Output{
				Address: event.Destination,
				Value:   event.Amount,
				Index:   event.Meta.Index,
			})
			totalOutputs += event.Amount
		}
	}

	return models.Transaction{
		TxID:              v.ID,
		Confirmations:     v.Confirmations,
		BlockHash:         v.BlockID,
		Timestamp:         v.Date,
		VBytes:            v.Meta.VSize,
		Inputs:            inputs,
		Outputs:           outputs,
		TotalInputsValue:  totalIputs,
		TotalOutputsValue: totalOutputs,
		Fee:               totalIputs - totalOutputs,
	}, nil
}
