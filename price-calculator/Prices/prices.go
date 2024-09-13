package prices

import (
	"fmt"

	"github.com/douglasmg7/go-study/price-calaculator/conversion"
	"github.com/douglasmg7/go-study/price-calaculator/filemanager"
)

type TaxIncludedPriceJobs struct {
	IOManager        filemanager.FileManager
	TaxRate          float64
	InputPrices      []float64
	TaxIncludedPrice map[string]string
}

func NewTaxIncludedPriceJobs(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJobs {
	return &TaxIncludedPriceJobs{
		IOManager: fm,
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}

func (jobs *TaxIncludedPriceJobs) LoadData() {
	lines, err := jobs.IOManager.ReadLines()
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

func (job *TaxIncludedPriceJobs) Process() {
	job.LoadData()
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		pricePlusTax := (1 + job.TaxRate) * price
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", pricePlusTax)
	}
	job.TaxIncludedPrice = result
	job.IOManager.WriteResult(job)
}
