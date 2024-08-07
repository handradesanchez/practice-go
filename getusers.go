package main

import (
	"fmt"
)

type User struct {
	Name string
	Age int
}

func GetUsers () {
	getUsers()
}

func getUsers() {

	users := []User{
		{Name: "Juan",  Age: 20},
		{Name: "Pedro", Age: 30},
	}

	for i, user := range users {
		fmt.Println("-------------------")
		fmt.Println("Name:  " + user.Name)
		fmt.Println("Age:  ", user.Age)
		fmt.Println("Order:", i + 1)
	}

}