package main

type RedisDB struct {}

func (redis *RedisDB) GetUsers() ([]User) {

	users := []User{
		{Name: "Johny",  Age: 22},
		{Name: "Peter", Age: 33},
	}

	return users

}

func (redis *RedisDB) DeleteUsers(name string) (User,error) {
	return User{Name: "", Age: 0}, nil
}
