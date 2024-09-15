package prices

import (
	"fmt"

	"github.com/douglasmg7/go-study/price-calaculator/conversion"
	"github.com/douglasmg7/go-study/price-calaculator/iomanager"
)

type TaxIncludedPriceJobs struct {
	IOManager        iomanager.IOManager `json:"-"`
	TaxRate          float64             `json:"tax_rate"`
	InputPrices      []float64           `json:"input_prices"`
	TaxIncludedPrice map[string]string   `json:"tax_included_price"`
}

func NewTaxIncludedPriceJobs(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJobs {
	return &TaxIncludedPriceJobs{
		IOManager:   iom,
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}

func (jobs *TaxIncludedPriceJobs) LoadData() error {
	lines, err := jobs.IOManager.ReadLines()
	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloat(lines)
	if err != nil {
		return err
	}
	jobs.InputPrices = prices
	return nil
}

func (job *TaxIncludedPriceJobs) Process() error {
	err := job.LoadData()
	if err != nil {
		return err
	}

	result := make(map[string]string)
	for _, price := range job.InputPrices {
		pricePlusTax := (1 + job.TaxRate) * price
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", pricePlusTax)
	}
	job.TaxIncludedPrice = result
	return job.IOManager.WriteResult(job)
}
