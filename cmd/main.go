package main

import (
	"log"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/optakt/flare-spammer/spammer"
)

func main() {
	spammer, err := spammer.New()
	if err != nil {
		log.Fatal(err)
	}
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	ticker := time.NewTicker(time.Duration(rand.Int()%2000) * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			err := spammer.CreateRandomTransactions(1)
			if err != nil {
				log.Fatal(err)
			}
		case <-sig:
			os.Exit(0)
		}
	}
}
