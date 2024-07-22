package prices

import (
	"fmt"

	"example.com/calculator/conversation"
	"example.com/calculator/iomanager"
)

type TaxIncludedPricesJob struct {
	IOManager iomanager.IOManager    `json."-"`
	TaxRate           float64            `json."tax_rate"`
	InputPrices       []float64          `json."input_prices"`
	taxIncludedPrices map[string]float64 `json."tax_included_prices"`
}

func (job TaxIncludedPricesJob) LoadData() error {
    lines, err :=  job.IOManager.ReadLines()

	if err != nil {
		
		return err
	}

    prices, err:= conversation.StringsToFloat(lines)



		   if err != nil {
		      

		        return err
		   }
		   	job.InputPrices = prices
			return nil
		
	}



func (job TaxIncludedPricesJob) Process() error {
	err := job.LoadData()

    if err != nil {
		return err
	}

	result := make(map[string]float64)

	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}
	job.taxIncludedPrices = result

	return job.IOManager.WriteResult(job)
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPricesJob {
	return &TaxIncludedPricesJob{
		IOManager: iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}