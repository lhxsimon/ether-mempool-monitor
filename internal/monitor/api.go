package monitor

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	log "github.com/sirupsen/logrus"
)

func (c *Client) Run() {
	ctx := context.Background()
	txChan := make(chan common.Hash)

	// init signer
	chainID, err := c.ethConn.NetworkID(ctx)
	if err != nil {
		log.Fatal(err)
	}
	signer := types.NewLondonSigner(chainID)

	// init subscriber
	subscriber := gethclient.New(c.rpcConn)
	_, err = subscriber.SubscribePendingTransactions(ctx, txChan)
	if err != nil {
		log.Fatal(err)
	}

	// monitor transactions
	log.Infof("start monitor transactions ...")
	for tx := range txChan {
		txData, _, err := c.ethConn.TransactionByHash(ctx, tx)
		if err != nil {
			continue
		}
		message, err := txData.AsMessage(signer, nil)
		if err != nil {
			log.Errorf("get message error: %v", err)
			continue
		}
		log.Infof("tx: %s, to: %s", tx.String(), message.To())
	}
}
