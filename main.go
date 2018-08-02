package main

import (
	"fmt"

	candlev1 "github.com/olivere/candle"
	candlev2 "github.com/olivere/candle/v2"
)

func main() {
	fmt.Printf("Candle 1 version %s\n", candlev1.Version())
	fmt.Printf("Candle 2 version %s\n", candlev2.Version())
}
