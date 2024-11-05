package main

import "fmt"

func main() {
	var accountBalance = 0.0
	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your balance is", accountBalance)
		} else if choice == 2 {
			fmt.Println("Your deposits: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount.")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
		} else if choice == 3 {
			fmt.Println("Withdrawal amount: ")
			var WithdrawalAmount float64
			fmt.Scan(&WithdrawalAmount)

			if WithdrawalAmount <= 0 {
				fmt.Println("Invalid amount.")
				continue
			}

			if WithdrawalAmount > accountBalance {
				fmt.Println("Invalid amount.")
				continue
			}

			accountBalance -= WithdrawalAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
		} else {
			fmt.Println("Goodbye!")
			break
		}
	}

	fmt.Println("Thanks for choosing our bank.")
}
