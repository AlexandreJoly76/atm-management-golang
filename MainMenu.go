package main

import (
	"fmt"
	"os"
)

func MainMenu(user *User) {
	fmt.Println("[1] Create Account")
	fmt.Println("[2] Deposit")
	fmt.Println("[3] Withdraw")
	fmt.Println("[4] Transfer")
	fmt.Println("[5] Delete Account")
	fmt.Println("[6] Check Account")
	fmt.Println("[7] Check All Accounts")
	fmt.Println("[8] Exit")
	var Choice int
	fmt.Scanln(&Choice)
	switch Choice {
	case 1: // Create Account
		CreateAccount(user)
	case 2: // Deposit
		//Deposit()
	case 3: // Withdraw
		//Withdraw()
	case 4: // Transfer
		//Transfer()
	case 5: // Delete Account
		//DeleteAccount()
	case 6: // Check Account
		//CheckAccount()
	case 7: // Check All Accounts
		//CheckAllAccounts()
	case 8: // Exit
		os.Exit(0)
	default:
		fmt.Println("Invalid choice")
	}
}
