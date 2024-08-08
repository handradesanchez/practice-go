package main

type User struct {
	Name string
	Age int
}

func GetUsers () []User {
	return getUsers()
}

func getUsers() []User {

	users := []User{
		{Name: "Juan",  Age: 20},
		{Name: "Pedro", Age: 30},
	}

	return users

}
