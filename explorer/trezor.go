package explorer

import (
	"encoding/json"
	"fmt"
	"strconv"

	httpclient "github.com/ranjbar-dev/bitcoin-wallet/httpClient"
)

type TrezorExplorer struct {
	baseURL string
}

type TrezorBalanceResponse struct {
	Balance string `json:"balance"`
}

func (e TrezorExplorer) GetAddressBalance(address string) (int, error) {

	headers := map[string]string{
		"Accept": "application/json",
	}

	url := fmt.Sprintf("%s/address/%s", e.baseURL, address)

	client := httpclient.NewHttpclient()

	res, err := client.NewRequest().SetHeaders(headers).Get(url)

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
