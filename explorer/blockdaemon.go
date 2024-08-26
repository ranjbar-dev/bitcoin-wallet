package explorer

import (
	"encoding/json"
	"fmt"
	"strconv"

	httpclient "github.com/ranjbar-dev/bitcoin-wallet/httpClient"
	"github.com/ranjbar-dev/bitcoin-wallet/models"
)

type BlockdaemonExplorer struct {
	Protocol string
	Network  string
	ApiKey   string
	baseURL  string
}

type BlockdaemonBalanceResponse struct {
	ConfirmedBalance string `json:"confirmed_balance"`
	ConfirmedBlock   int    `json:"confirmed_block"`
}

func (e BlockdaemonExplorer) GetAddressBalance(address string) (int, error) {

	headers := map[string]string{
		"X-API-Key": e.ApiKey,
		"Accept":    "application/json",
	}

	url := fmt.Sprintf("%s/%s/%s/account/%s", e.baseURL, e.Protocol, e.Network, address)

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

	url := fmt.Sprintf("%s/%s/%s/sync/block_number", e.baseURL, e.Protocol, e.Network)

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

	url := fmt.Sprintf("%s/%s/%s/sync/block_id", e.baseURL, e.Protocol, e.Network)

	client := httpclient.NewHttpclient()

	res, err := client.NewRequest().SetHeaders(headers).Get(url)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	fmt.Println(res.String())

	return res.String(), nil
}

func (e BlockdaemonExplorer) GetBlockByNumber(int) (models.BlockChainBlock, error) {

	return models.BlockChainBlock{}, nil
}
