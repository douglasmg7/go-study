package prices

import (
	"fmt"
	"github.com/douglasmg7/go-study/price-calaculator/conversion"
	"github.com/douglasmg7/go-study/price-calaculator/filemanager"
)

type TaxIncludedPriceJobs struct {
	TaxRate          float64
	InputPrices      []float64
	TaxIncludedPrice map[string]float64
}

func NewTaxIncludedPriceJobs(taxRate float64) *TaxIncludedPriceJobs {
	return &TaxIncludedPriceJobs{
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}

func (jobs *TaxIncludedPriceJobs) LoadData() {
	lines, err := filemanager.ReadLines("prices.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloat(lines)
	if err != nil {
		fmt.Println(err)
		return
	}
	jobs.InputPrices = prices
}

func (jobs *TaxIncludedPriceJobs) Process() {
	jobs.LoadData()
	result := make(map[string]string)
	for _, price := range jobs.InputPrices {
		pricePlusTax := (1 + jobs.TaxRate) * price
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", pricePlusTax)
	}
	fmt.Println(result)
}
