package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	scanner := bufio.NewScanner(os.Stdin)
	cmdFlags := CmdFlags{}

	users, err := LoadUsers("users.json")
	if err != nil {
		fmt.Println("Error loading users:", err)
		return
	}

	currentUser := ""
	for {
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")

		if scanner.Scan() {
			choice := scanner.Text()
			switch choice {
			case "1":
				currentUser = login(scanner, &users)
				if currentUser != "" {
					fmt.Println("Welcome", currentUser, "!")
					runTodoApp(scanner, &todos, storage, cmdFlags, currentUser)
				}
			case "2":
				register(scanner, &users)
			case "3":
				fmt.Println("Exiting...")
				return
			default:
				fmt.Println("Invalid option. Please try again.")
			}
		} else {
			fmt.Println("Exiting...")
			return
		}
	}
}

func login(scanner *bufio.Scanner, users *Users) string {
	fmt.Print("Enter username: ")
	scanner.Scan()
	username := scanner.Text()

	fmt.Print("Enter password: ")
	scanner.Scan()
	password := scanner.Text()

	if users.Authenticate(username, password) {
		return username
	}

	fmt.Println("Invalid username or password")
	return ""
}

func register(scanner *bufio.Scanner, users *Users) {
	fmt.Print("Enter new username: ")
	scanner.Scan()
	username := scanner.Text()

	fmt.Print("Enter new password: ")
	scanner.Scan()
	password := scanner.Text()

	users.AddUser(username, password)
	err := SaveUsers(*users, "users.json")
	if err != nil {
		fmt.Println("Error saving users:", err)
	} else {
		fmt.Println("User registered successfully")
	}
}

func runTodoApp(scanner *bufio.Scanner, todos *Todos, storage *storage[Todos], cmdFlags CmdFlags, currentUser string) {
	todos.print()
	fmt.Println("Enter commands (list, add, edit, del, toggle, exit):")
	for {
		fmt.Print("> ")
		if scanner.Scan() {
			input := scanner.Text()
			cmdFlags.Parse(input)
			cmdFlags.Execute(todos, currentUser)
			storage.Save(*todos)
			if cmdFlags.Exit {
				fmt.Println("Logging out...")
				break
			}
		} else {
			fmt.Println("Logging out...")
			break
		}
	}
}
