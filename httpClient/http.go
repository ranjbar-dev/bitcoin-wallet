package httpClient

import (
	"net/http"
)

const UserAgent = "crypto-wallet"

type HttpClient interface {
	SimpleGet(url string, res interface{}) error
	SimplePost(url string, body interface{}, res interface{}) error
	HttpGet(url string, headers map[string]string) ([]byte, http.Header, int, error)
	HttpPost(url string, body interface{}, headers map[string]string) ([]byte, http.Header, int, error)
}

type httpClient struct {
}

func NewHttpClient() HttpClient {
	return &httpClient{}
}
