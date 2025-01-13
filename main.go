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
	todos.print()
	currentUser := ""

	fmt.Println("Enter your Name:")
	fmt.Print("> ")

	for {
		if scanner.Scan() {
			currentUser = scanner.Text()
			if currentUser == "" {
				fmt.Println("Please Enter your Name:")
				fmt.Print("> ")
				continue
			}
			fmt.Println("Welcome", currentUser, "!")
			break
		}
	}

	fmt.Println("Enter commands (list, add, edit, del, toggle, exit):")
	for {
		fmt.Print("> ")
		if scanner.Scan() {
			input := scanner.Text()
			cmdFlags.Parse(input)
			cmdFlags.Execute(&todos, currentUser)
			storage.Save(todos)
			if cmdFlags.Exit {
				fmt.Println("Exiting...")
				break
			}
		} else {
			fmt.Println("Exiting...")
			break
		}
	}
}
