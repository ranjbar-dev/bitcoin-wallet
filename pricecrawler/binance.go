package pricecrawler

import (
	"encoding/json"
	"fmt"
	"strconv"

	httpclient "github.com/ranjbar-dev/bitcoin-wallet/httpClient"
)

type BinanceCrawler struct {
	BaseURL string
}

type BinancePriceResponse struct {
	Price string `json:"price"`
}

func (c BinanceCrawler) GetBTCPrice() (float64, error) {

	client := httpclient.NewHttpclient()

	url := fmt.Sprintf("%s/ticker/price?symbol=BTCUSDT", c.BaseURL)

	res, err := client.NewRequest().Get(url)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	var v BinancePriceResponse

	err = json.Unmarshal(res.Body(), &v)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	price, err := strconv.ParseFloat(v.Price, 64)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return price, nil
}
