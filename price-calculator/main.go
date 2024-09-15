package main

import (

	// "github.com/douglasmg7/go-study/price-calaculator/cmdmanager"
	"fmt"

	"github.com/douglasmg7/go-study/price-calaculator/filemanager"
	"github.com/douglasmg7/go-study/price-calaculator/prices"
)

func main() {
	taxRates := []float64{0, .07, .1, .15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%0.f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJobs(fm, taxRate)
		err:=priceJob.Process()
		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}
	}
}
