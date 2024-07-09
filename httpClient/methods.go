package httpclient

import "github.com/go-resty/resty/v2"

func (h *Httpclient) Get(url string, params map[string]string) (*resty.Response, error) {

	return h.c.R().SetQueryParams(params).Get(url)
}

func (h *Httpclient) Post(url string, body map[string]interface{}) (*resty.Response, error) {

	return h.c.R().SetBody(body).Post(url)
}
