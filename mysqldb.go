package main

import "errors"

type mySQLDB struct{}

func (mySQL mySQLDB) getUsers() []user {

	users := []user{
		{Name: "Juanito", Age: 83},
		{Name: "Pedrito", Age: 1},
	}

	return users

}

func (mySQL mySQLDB) deleteUser(name string) (user,error) {
	if name == "" {
		return user{Name: "", Age: 0}, errors.New("error deleting user")
	}
	return user{Name: name, Age: 0}, nil
}
