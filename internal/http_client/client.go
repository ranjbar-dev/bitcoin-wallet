package httpclient

import (
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/ranjbar-dev/bitcoin-wallet/config"
)

type HttpClient struct {
	client *resty.Client
}

var defaultHeaders = map[string]string{
	"User-Agent":   config.GetHttpClientAgent(),
	"Accept":       "application/json",
	"Content-Type": "application/json",
}

func NewHttpClient() *HttpClient {

	client := resty.New()

	client = client.SetTimeout(time.Second * time.Duration(config.GetHttpClientTimeout()))

	client.SetDebug(config.GetHttpClientDebug())

	return &HttpClient{
		client: client,
	}
}
