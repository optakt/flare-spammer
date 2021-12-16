package main

import (
	"fmt"
	spammer2 "github.com/optakt/flare-spammer/spammer"
	"os"
)

func main() {
	spammer := spammer2.NewSpammer()
	err := spammer.CreateRandomTransactions(10)

	if err != nil {
		fmt.Printf("couldn't run spammer: %s\n", err)
		os.Exit(1)
	}
}
