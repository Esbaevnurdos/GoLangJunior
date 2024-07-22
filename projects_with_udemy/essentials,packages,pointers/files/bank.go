package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func getBalanceToFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

if err != nil {
	return 1000, errors.New("Failed to find balance file.")
}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
	return 1000, errors.New("Failed to parse stored balance value.")
}
	return balance, nil
}


func main() {
    var accountBalance, err = getBalanceToFile()

    if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------")
		// panic("Can't continue, sorry.")
	}

	fmt.Println("Welcome to Go Bank")
	fmt.Println("Reach us 24/7" ,randomdata.PhoneNumber())

	for {
    questionValue()

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	

	// wantsCheckBalance := choice == 1

switch choice {

case 1:
	fmt.Println("Your balance is", accountBalance)

case 2:
    fmt.Print("Your Deposit: ")
    var depositAmount float64
    fmt.Scan(&depositAmount)
     
		 if depositAmount <= 0 {
    fmt.Println("Invalid amount. Must be greater than 0.")
	continue
	}

	accountBalance += depositAmount // accountBalance = depositAmount + accountBalance
	fmt.Println("Balance updated! New amount: ", accountBalance)
	writeBalanceToFile(accountBalance)
case 3:
	    fmt.Print("Withdrawal amount: ")
        var withdrawalAmount float64
        fmt.Scan(&withdrawalAmount)
        accountBalance -= withdrawalAmount // accountBalance = depositAmount + accountBalance
		fmt.Println("Balance updated! New amount: ", accountBalance)
		writeBalanceToFile(accountBalance)

		     if withdrawalAmount <= 0 {
        fmt.Println("Invalid amount. Must be greater than 0.")
		continue
		}

			     if withdrawalAmount > accountBalance {
        fmt.Println("Invalid amount. You can't withdraw more than you have.")
		continue
		}
	default:
		 fmt.Println("Goodbye!")
		 fmt.Println("Thanks for choosing our bank")
		return
		
}
}
}


