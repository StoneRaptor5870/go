package main

import "fmt"

func main() {
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	rate := getUserInput("Rate: ")

	ebt := revenue - expenses
	profit := ebt * (1 - rate/100)
	ratio := ebt / profit

	fmt.Println(ebt, profit, ratio)
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}
