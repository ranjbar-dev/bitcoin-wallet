package feecrawler

import (
	"encoding/json"
	"fmt"

	httpclient "github.com/ranjbar-dev/bitcoin-wallet/httpClient"
)

type BlockchainCrawler struct {
	BaseURL string
}

type BlockChainResponse struct {
	Limits struct {
		Min float64 `json:"min"`
	}
	Regular  float64 `json:"regular"`
	Priority float64 `json:"priority"`
}

func (c BlockchainCrawler) SetBaseURL(url string) {
	c.BaseURL = url
}

func (c BlockchainCrawler) GetEstimatedFee() (float64, float64, float64, error) {
	client := httpclient.NewHttpclient()

	url := fmt.Sprintf("%s/mempool/fees", c.BaseURL)

	res, err := client.NewRequest().Get(url)

	if err != nil {
		fmt.Println(err)
		return 0, 0, 0, err
	}

	// var v map[string]interface{}
	var v BlockChainResponse

	err = json.Unmarshal(res.Body(), &v)

	if err != nil {
		fmt.Println(err)
		return 0, 0, 0, err
	}

	fmt.Println(v)

	return v.Limits.Min, v.Regular, v.Priority, nil
}
