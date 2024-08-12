package main

import "errors"

type mongoDB struct{}

type mongoUser struct {
	Name string
	Age  int
}

func (mongo mongoDB) getUsers() []user {
	var users []user

	mongoUsers := mongo.getMongoUsers()

	for _, mongoUsr := range mongoUsers {
		// convert mongoUser to user and append
		newUser := user(mongoUsr)
		users = append(users, newUser)
	}

	return users

}

func (mongo mongoDB) getMongoUsers() []mongoUser {
	return []mongoUser{
		{
			Name: "shrek",
			Age:  40,
		},
		{
			Name: "fiona",
			Age:  35,
		},
	}
}

func (mongo mongoDB) deleteUser(name string) (user, error) {
	if name == "" {
		return user{Name: "", Age: 0}, errors.New("error deleting user")
	}
	return user{Name: name, Age: 0}, nil
}
