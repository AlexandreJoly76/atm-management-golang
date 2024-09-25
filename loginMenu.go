package main

import (
	"fmt"
	"os"
)

func LoginMenu() {
	ClearScreen()
	var Choice int
	fmt.Println("[1] Login")
	fmt.Println("[2] Register")
	fmt.Println("[3] Exit")
	fmt.Scanln(&Choice)
	switch Choice {
	case 1:
		user := User{}
		for user.Name == "" {
			ClearScreen()
			fmt.Println("Enter your name:")
			var name, password string
			fmt.Scanln(&name)
			fmt.Println("Enter your password:")
			fmt.Scanln(&password)
			user = login(name, password)
			ClearScreen()
			fmt.Printf("Welcome: %s,\n", user.Name)
			fmt.Println("-----------------------------")
		}
		MainMenu(&user)
	case 2:
		var NewAccount User
		fmt.Println("Enter your name:")
		fmt.Scanln(&NewAccount.Name)
		fmt.Println("Enter your password:")
		var password string
		fmt.Scanln(&password)
		user := Register(NewAccount, password)
		fmt.Println("Account created successfully")
		ClearScreen()
		fmt.Printf("Welcome: %s,", NewAccount.Name)
		fmt.Println("-----------------------------")
		MainMenu(&user)
	case 3:
		os.Exit(0)
		ClearScreen()
	default:
		fmt.Println("Invalid choice")
	}

}
