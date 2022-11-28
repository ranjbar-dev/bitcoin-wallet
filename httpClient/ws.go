package httpClient

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type WsConnection interface {
	WriteMessage(messageType int, data []byte) error
	ReadMessage() (messageType int, p []byte, err error)
	Close() error
}

type WsClient interface {
	Dial(urlStr string, requestHeader http.Header) (WsConnection, error)
}

type wsClient struct {
	dialer *websocket.Dialer
}

func (wsClient *wsClient) Dial(urlStr string, requestHeader http.Header) (WsConnection, error) {
	c, _, err := wsClient.dialer.Dial(urlStr, requestHeader)
	return c, err
}

func NewWsClient() WsClient {
	dialer := websocket.DefaultDialer
	return &wsClient{dialer}
}
