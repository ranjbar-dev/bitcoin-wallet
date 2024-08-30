package pricecrawler

import (
	"encoding/json"
	"fmt"

	httpclient "github.com/ranjbar-dev/bitcoin-wallet/httpClient"
)

type CoinGeckoCrawer struct {
	BaseURL string
	Token   string
}

type CoinGeckoPriceResponse struct {
	Bitcoin struct {
		Usd float64 `json:"usd"`
	} `json:"bitcoin"`
}

func (c CoinGeckoCrawer) GetBTCPrice() (float64, error) {
	client := httpclient.NewHttpclient()

	url := fmt.Sprintf("%s/simple/price?ids=bitcoin&vs_currencies=usd&precision=5", c.BaseURL)

	res, err := client.NewRequest().Get(url)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	var v CoinGeckoPriceResponse

	err = json.Unmarshal(res.Body(), &v)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return v.Bitcoin.Usd, nil
}
