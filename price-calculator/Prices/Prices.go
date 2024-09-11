package Prices

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
