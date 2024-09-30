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
	for i, record := range Records {
		Records[i].Owner = GetOwnerById(record.OwnerId)
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

func UpdateRecords() {
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

func (record *Record) Deposit(amount float64) {
	record.Amount += amount
}

func (record *Record) Withdraw(amount float64) {
	record.Amount -= amount
}

func (record *Record) Transfer(NewOwner User) {
	record.Owner = NewOwner
	record.OwnerId = NewOwner.Id
}

func (record *Record) Delete() {
	for i := 0; i < len(Records); i++ {
		if Records[i].Id == record.Id {
			Records = append(Records[:i], Records[i+1:]...)
			UpdateRecords()
		}
	}
}

func GetOwnerById(id int) User {
	for _, user := range Users {
		if user.Id == id {
			return user
		}
	}
	return User{}
}

func Exit(user *User) {
	var Choice int
	fmt.Println("[1] Exit")
	fmt.Println("[2] Return to the main menu")
	fmt.Scanln(&Choice)
	switch Choice {
	case 1:
		os.Exit(0)
	case 2:
		ClearScreen()
		fmt.Println("Welcome:", user.Name)
		fmt.Println("-----------------------------")
		MainMenu(user)
	}
}

func CheckAllAccount(user *User) {
	for _, record := range Records {
		if record.OwnerId == user.Id {
			fmt.Println("AccountNumber:", record.AccountNumber)
			fmt.Println("CreationDate:", record.CreationDate)
			fmt.Println("Country:", record.Country)
			fmt.Println("PhoneNumber:", record.PhoneNumber)
			fmt.Println("Amount:", record.Amount)
			fmt.Println("AccountType:", record.AccountType)
			fmt.Println("-------------------------")
		}
	}
	Exit(user)
}

func CheckAllAccountNumberOfUser(user *User) {
	for _, record := range Records {
		if record.OwnerId == user.Id {
			fmt.Println("AccountNumber:", record.AccountNumber)
			fmt.Println("-------------------------")
		}
	}
}