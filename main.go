package main

import (
	"fmt"

	"github.com/markcheno/go-quote"
	"github.com/markcheno/go-talib"
)

func main() {
	aapl, _ := quote.NewQuoteFromYahoo("AAPL", "2021-01-01", "2021-09-10", quote.Daily, true)
	fmt.Print(aapl.CSV())
	rsi2 := talib.Rsi(aapl.Close, 2)
	fmt.Println(rsi2)
}
