package main

import (
	"fmt"
	"os"

	"github.com/piquette/finance-go/quote"
	"github.com/piquette/finance-go/equity"
)

func run() error {
	// Quote
	q, err := quote.Get("AAPL")
	if err != nil {
		return err
	}
	// https://piquette.io/projects/finance-go/#basics
	fmt.Printf("Symbol: %s\nPrice: %v\n\n", q.Symbol, q.RegularMarketPrice)

	// Equity
	symbols := []string{"IBM", "GOOG", "MSFT"}
	iter := equity.List(symbols)

	if iter.Err() != nil {
		return iter.Err()
	} 
	for iter.Next() {
		e := iter.Equity()
		fmt.Printf("Symbol: %s\nForward P/E: %v\nEPS Fwd: %v\nDividend: %v\nMarket Cap: %v\n\n", e.Quote.Symbol, e.ForwardPE, e.EpsForward, e.TrailingAnnualDividendRate, e.MarketCap/1000000)
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("error running the examples: %s\n", err)
		os.Exit(1)
	}
}