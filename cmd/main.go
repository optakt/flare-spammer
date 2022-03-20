package main

import (
	"fmt"
	spammer2 "github.com/optakt/flare-spammer/spammer"
	"os"
	"time"
)

func main() {
	spammer := spammer2.NewSpammer()
	var err error
	for {
		err = spammer.CreateRandomTransactions(10)
		time.Sleep(60 * time.Second)
	}

	if err != nil {
		fmt.Printf("couldn't run spammer: %s\n", err)
		os.Exit(1)
	}
}
