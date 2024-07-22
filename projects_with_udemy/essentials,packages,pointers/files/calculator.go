// package main

// import (
// 	"fmt"
// 	"math"
// )

// 	const inflationRate = 2.5


// func main() {
// 	var expectedReturnRate float64
// 	var investAmount float64
// 	var years float64

// 	fmt.Print("Investment Amount: ")
// 	fmt.Scan(&investAmount)

// 	fmt.Print("Years: ")
// 	fmt.Scan(&years)

// 	fmt.Print("Expected Return Rate: ")
// 	fmt.Scan(&expectedReturnRate)

// 	futureValue, futurRealValue := calculateFutureValues(investAmount, expectedReturnRate, years)
// 	// var futureValue = float64(investAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
// 	// var futurRealValue = futureValue / math.Pow(1+inflationRate/100, years)

// 	fmt.Println(futureValue)
// 	fmt.Println(futurRealValue)
// }

// func calculateFutureValues(investAmount, expectedReturnRate, years float64) (fv float64,rfv float64) {
// 	fv = investAmount * math.Pow(1+expectedReturnRate/100, years)
// 	rfv = fv / math.Pow(1+inflationRate/100, years)
// 	return fv, rfv
// 	// return
// }

// CALCULATOR
// func main() {
// 	var revenue int
// 	var expenses int
// 	var taxRate float64

// 	// fmt.Print("Revenue: ")
// 	outputText("Revenue: ")
// 	fmt.Scan(&revenue)

// 	// fmt.Print("Expenses: ")
// 	outputText("Expenses: ")
// 	fmt.Scan(&expenses)

// 	// fmt.Print("Tax Rate: ")
// 	outputText("Tax Rate: ")
// 	fmt.Scan(&taxRate)

// 	ebt := revenue - expenses
// 	profit := float64(ebt) * (1 - taxRate/100)
// 	ratio := float64(ebt) / profit

//     fmt.Println("ebt: ", ebt)
// 	fmt.Println(profit)
// 	fmt.Println(ratio)

//     formattedFV := fmt.Sprintf("Future Value: %.1f\n", ebt)
//     formattedRFV := fmt.Sprintf("Future Value (something): %.1f\n", profit)

// 	fmt.Printf(formattedFV, formattedRFV)
// }

// func outputText (text string) {
// 	fmt.Print(text)
// }
