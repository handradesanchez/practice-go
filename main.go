package main

import "fmt"

func main() {

	users := GetUsers()
	for _, user := range users {
		fmt.Printf("Name: %s\t Age: %d\n", user.Name, user.Age)
	}

	HttpServer()

}
