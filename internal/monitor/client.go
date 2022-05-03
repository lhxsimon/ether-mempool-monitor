package monitor

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	rpcConn *rpc.Client
	ethConn *ethclient.Client
}

func NewClient(rpcURL string) (*Client, error) {
	log.Infof("connecting to ethereum node %s ...", rpcURL)
	rpcConn, err := rpc.Dial(rpcURL)
	if err != nil {
		return nil, err
	}
	ethConn, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}
	return &Client{
		rpcConn: rpcConn,
		ethConn: ethConn,
	}, nil
}

func (c *Client) Close() {
	c.rpcConn.Close()
	c.ethConn.Close()
}
