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

	fmt.Println("Enter commands (add, edit, del, toggle, list):")
	for {
		fmt.Print("> ")
		if scanner.Scan() {
			input := scanner.Text()
			cmdFlags.Parse(input)
			cmdFlags.Execute(&todos)
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
