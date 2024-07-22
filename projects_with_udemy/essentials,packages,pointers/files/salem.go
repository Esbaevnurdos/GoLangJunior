// package main

// import (
// 	"errors"
// 	"fmt"
// 	"os"
// )


// func main() {

// 	// var revenue int
// 	// var expenses int
// 	// var taxRate float64

// 	revenue, err1 := getUserInput("Revenue: ")

// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }

// 	expenses, err2:= getUserInput("Expenses: ")

// 	// 	if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }

// 	taxRate, err3 := getUserInput("Tax Rate: ")

// 	if err1 != nil || err2 != nil || err3 != nil {
// 		fmt.Println(err1)
// 		return
// 	}

//     ebt, profit, ratio := calculateFinancials(revenue,expenses,taxRate)

//     fmt.Printf("%.1f\n", ebt)
// 	fmt.Printf("%.1f\n", profit)
// 	fmt.Printf("%.1f\n", ratio)
// 	storeUserValues(ebt, profit, ratio)
// }


// func storeUserValues (ebt, profit, ratio float64) {
// 	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n",ebt,profit,ratio)
// os.WriteFile("results.txt",[]byte(results), 0644)
// }


// func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
// 	ebt := revenue - expenses
// 	profit := float64(ebt) * (1 - taxRate/100)
// 	ratio := float64(ebt) / profit
// 	return ebt, profit, ratio
// }

// func getUserInput(infoText string) (float64, error) {
//     fmt.Print(infoText)
//     var getInfo float64
//     fmt.Scan(&getInfo)

//     if getInfo <= 0 {
// 		return 0, errors.New("Value must be positive number.")
// 	}

//     return getInfo, nil
// }