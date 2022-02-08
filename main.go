package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dustin/go-humanize"
	coingecko "github.com/superoo7/go-gecko/v3"
)

func main() {
	// Init the CoinGecko client
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	cg := coingecko.NewClient(httpClient)

	// Get global data
	global, err := cg.Global()
	if err != nil {
		log.Fatal(err)
	}

	// The global crypto market cap
	var totalMarketCap float64 = global.TotalMarketCap["usd"]
	var tmc int64 = int64(totalMarketCap)

	fmt.Printf("TMC is $%s\n", humanize.Comma(tmc))
}
