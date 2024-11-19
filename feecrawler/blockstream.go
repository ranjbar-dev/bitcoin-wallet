package feecrawler

import (
	"encoding/json"
	"fmt"

	httpclient "github.com/ranjbar-dev/bitcoin-wallet/httpClient"
)

type BlockstreamCrawler struct {
	BaseURL string
}

type BlockStreamResponse struct {
	Low    float64 `json:"144"`
	Medium float64 `json:"25"`
	High   float64 `json:"1"`
}

func (c BlockstreamCrawler) SetBaseURL(url string) {
	c.BaseURL = url
}

func (c BlockstreamCrawler) GetEstimatedFee() (float64, float64, float64, error) {
	client := httpclient.NewHttpclient()

	url := fmt.Sprintf("%s/fee-estimates", c.BaseURL)

	res, err := client.NewRequest().Get(url)

	if err != nil {
		fmt.Println(err)
		return 0, 0, 0, err
	}

	// var v map[string]interface{}
	var v BlockStreamResponse

	err = json.Unmarshal(res.Body(), &v)

	if err != nil {
		fmt.Println(err)
		return 0, 0, 0, err
	}

	return v.Low, v.Medium, v.High, nil
}
