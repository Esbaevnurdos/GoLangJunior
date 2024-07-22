package main

import (
	"fmt"

	"example.com/calculator/cmdmanager"
	"example.com/calculator/prices"
)



func main() {
taxRates := []float64{0, 0.7, 10.7}


	for _, taxRate := range taxRates {
		// fm := filemanager.New("price.txt", fmt.Sprintf("result_%.0f.json", taxRate* 100))
		cmdm := cmdmanager.New()
priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
 err := priceJob.Process()

 if err != nil {
	fmt.Println("Could not proccess job")
	fmt.Println(err)
 }

	}
}