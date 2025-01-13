package main

import (
	"encoding/json"
	"os"
)

type User struct {
	Username string
	Password string
}

type Users []User

func (u *Users) AddUser(username, password string) {
	*u = append(*u, User{Username: username, Password: password})
}

func (u *Users) Authenticate(username, password string) bool {
	for _, user := range *u {
		if user.Username == username && user.Password == password {
			return true
		}
	}
	return false
}

func SaveUsers(users Users, filename string) error {
	data, err := json.MarshalIndent(users, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func LoadUsers(filename string) (Users, error) {
	var users Users
	if !FileExists(filename) {
		data, err := json.MarshalIndent(users, "", "    ")
		if err != nil {
			return users, err
		}

		os.WriteFile(filename, data, 0644)

	}
	data, err := os.ReadFile(filename)
	if err != nil {
		return users, err
	}
	err = json.Unmarshal(data, &users)
	return users, err
}
