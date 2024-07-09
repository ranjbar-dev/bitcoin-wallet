package httpclient

import (
	"crypto/tls"
	"time"

	"github.com/go-resty/resty/v2"
)

type Httpclient struct {
	c *resty.Client
}

func (h *Httpclient) Client() *resty.Client {

	return h.c
}

func NewHttpclient() *Httpclient {

	return &Httpclient{
		c: resty.New().SetTimeout(time.Second * 5).SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}),
	}
}
