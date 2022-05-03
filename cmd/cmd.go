package main

import (
	"eth/internal/monitor"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main() {
	prepare()
	err := runCmd()
	if err != nil {
		log.Fatal(err)
	}
}

func runCmd() error {
	// 设置命令行
	var rootCmd = &cobra.Command{
		Use:     "ether-monitor",
		Short:   "Ethereum Monitor",
		Example: "ether-monitor ETHER_NODE_URL",
		Args:    cobra.ExactArgs(1),
		Run:     cmdRunFunc,
	}
	return rootCmd.Execute()
}

func cmdRunFunc(cmd *cobra.Command, args []string) {
	rpcURL := args[0]
	if rpcURL == "" {
		log.Fatal("rpc url empty")
	}
	monitorClient, err := monitor.NewClient(rpcURL)
	if err != nil {
		log.Fatalf("connect to rpc %s error: %s", rpcURL, err)
	}
	defer monitorClient.Close()
	monitorClient.Run()
}
