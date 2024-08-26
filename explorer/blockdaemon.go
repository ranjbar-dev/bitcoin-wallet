package explorer

import (
	"encoding/json"
	"fmt"
	"strconv"

	httpclient "github.com/ranjbar-dev/bitcoin-wallet/httpClient"
)

const baseURL string = "https://svc.blockdaemon.com/universal/v1"

type BlockdaemonExplorer struct {
	Protocol string
	Network  string
	ApiKey   string
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

	url := fmt.Sprintf("%s/%s/%s/account/%s", baseURL, e.Protocol, e.Network, address)

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
