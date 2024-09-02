package feecrawler

import (
	"encoding/json"
	"fmt"

	httpclient "github.com/ranjbar-dev/bitcoin-wallet/httpClient"
)

type BitCoinerCrawler struct {
	BaseURL string
}

type BitCoinerResponse struct {
	Estimates struct {
		Low struct {
			SatPerVByte float64 `json:"sat_per_vbyte"`
		} `json:"30"`
		Medium struct {
			SatPerVByte float64 `json:"sat_per_vbyte"`
		} `json:"120"`
		High struct {
			SatPerVByte float64 `json:"sat_per_vbyte"`
		} `json:"360"`
	} `json:"estimates"`
}

func (c BitCoinerCrawler) SetBaseURL(url string) {
	c.BaseURL = url
}

func (c BitCoinerCrawler) GetEstimatedFee() (float64, float64, float64, error) {
	client := httpclient.NewHttpclient()

	url := fmt.Sprintf("%s/fees/estimates/latest", c.BaseURL)

	res, err := client.NewRequest().Get(url)

	if err != nil {
		fmt.Println(err)
		return 0, 0, 0, err
	}

	// var v map[string]interface{}
	var v BitCoinerResponse

	err = json.Unmarshal(res.Body(), &v)

	if err != nil {
		fmt.Println(err)
		return 0, 0, 0, err
	}

	return v.Estimates.Low.SatPerVByte, v.Estimates.Medium.SatPerVByte, v.Estimates.High.SatPerVByte, nil
}
