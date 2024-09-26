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
	ClearScreen()
	switch Choice {
	case 1: // Create Account
		CreateAccount(user)
	case 2: // Deposit
		fmt.Println("Enter your account number:")
		var accountNumber int
		fmt.Scanln(&accountNumber)
		ClearScreen()
		for i, record := range Records {
			if record.AccountNumber == accountNumber && record.OwnerId == user.Id {
				var amount float64
				fmt.Println("Enter the amount you want to deposit:")
				fmt.Scanln(&amount)
				Records[i].Deposit(amount)
				UpdateRecords()
				ClearScreen()
				fmt.Printf("New solde: %2.f €\n", Records[i].Amount)
				fmt.Println("-------------------------")
				Exit(user)
			} else {
				fmt.Println("You don't have an account with this number")
				fmt.Println("-------------------------")
				MainMenu(user)
			}
		}
	case 3: // Withdraw
		for i := 0; i < len(Records); i++ {
			if Records[i].OwnerId == user.Id {
				var Withdraw float64
				fmt.Println("Enter what do you want withdraw:")
				fmt.Scanln(&Withdraw)
				ClearScreen()
				if Records[i].Amount < Withdraw {
					fmt.Println("You don't have enough money")
					fmt.Println("-------------------------")
					MainMenu(user)
				}
				Records[i].Withdraw(Withdraw)
				UpdateRecords()
				fmt.Println("-------------------------")
				fmt.Printf("New solde: %2.f €\n", Records[i].Amount)
				Exit(user)
			} else {
				ClearScreen()
				fmt.Println("You don't have an account with this number")
				fmt.Println("-------------------------")
				MainMenu(user)
			}
		}
	case 4: // Transfer Owner
		fmt.Println("Enter your account number:")
		var accountNumber int
		fmt.Scanln(&accountNumber)
		ClearScreen()
		fmt.Printf("Hello %s, you are on the account N°%d\n", user.Name, accountNumber)
		fmt.Println("-------------------------")
		for i := 0; i < len(Records); i++ {
			if Records[i].OwnerId == user.Id && Records[i].AccountNumber == accountNumber {
				var NewOwner User
				fmt.Println("Enter the new owner ID:")
				fmt.Scanln(&NewOwner.Id)
				ClearScreen()
				for _, user := range Users {
					if NewOwner.Id == user.Id {
						Records[i].Transfer(NewOwner)
						UpdateRecords()
						fmt.Printf("New owner of the account N°%d is %s\n", Records[i].AccountNumber, user.Name)
						fmt.Println("-------------------------")
						Exit(&user)
					} else {
						fmt.Println("You don't have an account with this number")
						fmt.Println("-------------------------")
						MainMenu(&user)
					}
				}

			}
		}
	case 5: // Delete Account
		fmt.Println("Enter your account number:")
		var accountNumber int
		fmt.Scanln(&accountNumber)
		ClearScreen()
		for i := 0; i < len(Records); i++ {
			if Records[i].OwnerId == user.Id {
				Records[i].Delete()
				fmt.Println("Account deleted successfully")
				fmt.Println("-------------------------")
				Exit(user)
			} else {
				fmt.Println("You don't have an account with this number")
				fmt.Println("-------------------------")
				MainMenu(user)
			}
		}
	case 6: // Check Account
		fmt.Println("Enter your account number:")
		var accountNumber int
		fmt.Scanln(&accountNumber)
		ClearScreen()
		for i := 0; i < len(Records); i++ {
			if Records[i].OwnerId == user.Id && Records[i].AccountNumber == accountNumber {
				fmt.Println("-------------------------")
				fmt.Println("Owner:", Records[i].Owner.Name)
				fmt.Println("Account Number:", Records[i].AccountNumber)
				fmt.Println("Creation Date:", Records[i].CreationDate)
				fmt.Println("Country:", Records[i].Country)
				fmt.Println("PhoneNumber:", Records[i].PhoneNumber)
				fmt.Println("Amount:", Records[i].Amount)
				fmt.Println("Account Type:", Records[i].AccountType)
				fmt.Println("-------------------------")
				Exit(user)
			} else {
				fmt.Println("You don't have an account with this number")
				fmt.Println("-------------------------")
				MainMenu(user)
			}
		}
	case 7: // Check All Accounts
		CheckAllAccount(user)
	case 8: // Exit
		os.Exit(0)
	default:
		fmt.Println("Invalid choice")
	}
}
