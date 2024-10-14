package ogmiosgo

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type CallBackFunc func(*Client, []byte)

type RPCParams struct {
	JsonRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
}

type Client struct {
	conn         *websocket.Conn
	doneCh       chan struct{}
	errCh        chan error
	CallBackFunc CallBackFunc
}

func NewClient(url string, header http.Header) (*Client, error) {
	conn, _, err := websocket.DefaultDialer.Dial(url, header)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn:   conn,
		doneCh: make(chan struct{}),
		errCh:  make(chan error, 20),
	}, nil
}

func (c *Client) Start() {
	go c.recvLoop()
}

func (c *Client) Done() chan<- struct{} {
	return c.doneCh
}

func (c *Client) recvLoop() {
	defer func() {
		close(c.doneCh)
		close(c.errCh)
	}()

	for {

		select {
		case <-c.doneCh:
			return
		default:
		}

		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			c.errCh <- err
			continue
		}

		if c.CallBackFunc != nil {
			c.CallBackFunc(c, msg)
		}
	}
}
