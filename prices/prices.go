package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/filemanager"
)

type TaxIndludedPriceJob struct {
	IOManager	filemanager.FileManager
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

func (job *TaxIndludedPriceJob) loadData() {
	lines, err := job.IOManager.ReadLines()

	if err!= nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloats(lines)

	if err!= nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = prices
}

func (job *TaxIndludedPriceJob) Process() {
	job.loadData()

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result

	job.IOManager.WriteResult(job)
}

func NewTaxIndludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIndludedPriceJob {
	return &TaxIndludedPriceJob{
		IOManager: fm,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}