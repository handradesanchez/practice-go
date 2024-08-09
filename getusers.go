package main

type user struct {
	Name string `query:"username"`
	Age int     `query:"age"`
}

func getUsers () []user {
	return getusers()
}

func getusers() []user {

	users := []user{
		{Name: "Juan",  Age: 20},
		{Name: "Pedro", Age: 30},
	}

	return users

}
