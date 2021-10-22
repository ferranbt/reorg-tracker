package main

import (
	"flag"
	"fmt"

	"github.com/umbracle/go-web3/blocktracker"
	"github.com/umbracle/go-web3/jsonrpc"
)

func main() {
	var endpoint string
	var maxReorg uint64

	flag.StringVar(&endpoint, "endpoint", "localhost:8545", "")
	flag.Uint64Var(&maxReorg, "max-reorg", 50, "")
	flag.Parse()

	provider, err := jsonrpc.NewClient(endpoint)
	if err != nil {
		panic(err)
	}
	tracker := blocktracker.NewBlockTracker(provider.Eth(), blocktracker.WithBlockMaxBacklog(maxReorg))
	if err := tracker.Init(); err != nil {
		panic(err)
	}
	if err := tracker.Start(); err != nil {
		panic(err)
	}

	sub := tracker.Subscribe()
	for {
		evnt := <-sub
		fmt.Println("-------")
		for _, i := range evnt.Removed {
			fmt.Printf("REMOVED: %d %s\n", i.Number, i.Hash)
		}
		for _, i := range evnt.Added {
			fmt.Printf("ADDED: %d %s\n", i.Number, i.Hash)
		}
	}
}
