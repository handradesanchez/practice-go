package main

import "errors"

type redisDB struct {}

func (redis redisDB) getUsers() ([]user) {

	users := []user{
		{Name: "Johny",  Age: 22},
		{Name: "Peter", Age: 33},
	}

	return users

}

func (redis redisDB) deleteUser(name string) (user,error) {
	if(name == "") {
		return user{Name: name, Age: 0}, errors.New("error deleting user")
	}
	return user{Name: name, Age: 0}, nil
}

func (redis *redisDB) thingOnlyRedisCanDO() error {
	return nil
}