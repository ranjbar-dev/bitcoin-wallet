package feecrawler

import (
	"encoding/json"
	"fmt"

	httpclient "github.com/ranjbar-dev/bitcoin-wallet/httpClient"
)

type BitCoinerCrawler struct {
	BaseURL string
}

func (c BitCoinerCrawler) GetEstimatedFee() (float64, error) {
	client := httpclient.NewHttpclient()

	url := fmt.Sprintf("%s/fees/estimates/latest", c.BaseURL)

	res, err := client.NewRequest().Get(url)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	var v map[string]interface{}

	err = json.Unmarshal(res.Body(), &v)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	fmt.Println(v)

	return 0.0, nil
}
