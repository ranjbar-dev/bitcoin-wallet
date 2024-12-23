package httpclient

import (
	"crypto/tls"
	"time"

	"github.com/go-resty/resty/v2"
)

type Httpclient struct {
	c *resty.Client
}

func (h *Httpclient) NewRequest() *resty.Request {

	return h.c.R()
}

func NewHttpclient() *Httpclient {

	return &Httpclient{
		c: resty.New().SetRetryCount(3).SetRetryWaitTime(time.Millisecond * 500).SetTimeout(time.Second * 10).SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}),
	}
}
