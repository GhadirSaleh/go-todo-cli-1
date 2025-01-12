package main

import (
	"fmt"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
	Exit   bool
}

func (cf *CmdFlags) Parse(cmd string) {
	if strings.HasPrefix(cmd, "add ") {
		cf.Add = strings.TrimSpace(strings.TrimPrefix(cmd, "add "))
		cf.Exit, cf.List, cf.Edit, cf.Del, cf.Toggle = false, false, "", -1, -1
	} else if strings.HasPrefix(cmd, "edit ") {
		cf.Edit = strings.TrimSpace(strings.TrimPrefix(cmd, "edit "))
		cf.Exit, cf.List, cf.Add, cf.Del, cf.Toggle = false, false, "", -1, -1
	} else if strings.HasPrefix(cmd, "del ") {
		index, _ := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(cmd, "del ")))
		cf.Del = index
		cf.Exit, cf.List, cf.Add, cf.Edit, cf.Toggle = false, false, "", "", -1
	} else if strings.HasPrefix(cmd, "toggle ") {
		index, _ := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(cmd, "toggle ")))
		cf.Toggle = index
		cf.Exit, cf.List, cf.Add, cf.Edit, cf.Del = false, false, "", "", -1
	} else if strings.HasPrefix(cmd, "list") {
		cf.List = true
		cf.Exit, cf.Add, cf.Edit, cf.Del, cf.Toggle = false, "", "", -1, -1
	} else if strings.HasPrefix(cmd, "exit") {
		cf.Exit = true
		cf.List, cf.Add, cf.Edit, cf.Del, cf.Toggle = false, "", "", -1, -1
	} else {
		cf.Exit, cf.List, cf.Add, cf.Edit, cf.Del, cf.Toggle = false, false, "", "", -1, -1
	}
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.Exit:
		return
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
		fmt.Println("Done!")
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error,Invalid format for edit.Please use id:new_title")
			break
		}
		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Error: invalid index for edit")
			break
		}
		todos.edit(index, parts[1])
		fmt.Println("Done!")

	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)
		fmt.Println("Done!")

	case cf.Del != -1:
		todos.delete(cf.Del)
		fmt.Println("Done!")

	default:
		fmt.Println("Invalid command")
	}
}
