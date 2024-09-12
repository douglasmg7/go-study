package main

import (
	"github.com/douglasmg7/go-study/price-calaculator/prices"
)

func main() {
	taxRates := []float64{0, .07, .1, .15}

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJobs(taxRate)
		priceJob.Process()
	}
}
