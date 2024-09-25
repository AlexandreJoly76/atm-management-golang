package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func ReadFile() {
	FileContent, err := os.ReadFile("./data/record.json")
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(FileContent, &Records)
	if err != nil {
		fmt.Println(err)
	}
	FileContent, err = os.ReadFile("./data/users.json")
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(FileContent, &Users)
	if err != nil {
		fmt.Println(err)
	}
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func login(name string, password string) User {
	for _, user := range Users {
		if user.Name == name && user.Password == password {
			fmt.Println("hey", user)
			return user
		} 
	}
	return User{}
}

func Register(NewAccount User, password string) User {
	NewAccount.Password = password
	NewAccount.Id = len(Users) + 1
	SaveUsers(NewAccount)
	return NewAccount
}

func SaveUsers(user User) {
	Users = append(Users, user)
	file, err := json.MarshalIndent(Users, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile("./data/users.json", file, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func SaveRecords(record Record) {
	Records = append(Records, record)
	file, err := json.MarshalIndent(Records, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile("./data/record.json", file, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func CreateAccount(user *User) {
	var newRecord Record
	newRecord.Owner = *user
	newRecord.OwnerId = user.Id
	fmt.Println("Enter your account ID :")
	fmt.Scanln(&newRecord.Id)
	newRecord.Id = len(Records) + 1
	ClearScreen()
	fmt.Println("Enter your AccountNumber:")
	fmt.Scanln(&newRecord.AccountNumber)
	ClearScreen()
	newRecord.CreationDate = time.Now().Format("2006-01-02")
	ClearScreen()
	fmt.Println("Enter your Country:")
	fmt.Scanln(&newRecord.Country)
	ClearScreen()
	fmt.Println("Enter your PhoneNumber:")
	fmt.Scanln(&newRecord.PhoneNumber)
	ClearScreen()
	fmt.Println("Enter your Amount:")
	fmt.Scanln(&newRecord.Amount)
	ClearScreen()
	fmt.Println("Enter your AccountType:")
	fmt.Scanln(&newRecord.AccountType)
	SaveRecords(newRecord)
	ClearScreen()
	fmt.Println("Account created successfully")
	fmt.Println("-----------------------------")
}
