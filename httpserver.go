package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func httpServer() {
	e := echo.New()
	e.GET("/helloworld", helloWorld)
	e.GET("/users", users)
	e.GET("/usersdb", usersdb)
	e.DELETE("/usersdb", deleteUser)
	e.GET("/users1", GetAllUsers)
	e.DELETE("/users1", deleteUsers)
	e.Logger.Fatal(e.Start(":8080"))
}

func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func users(c echo.Context) error {
	userNames := formatUsers(getUsers())

	return c.String(http.StatusOK, strings.Join(userNames, "\n"))
}

func formatUsers(users []user) []string {
	names := make([]string, len(users))
	for i, user := range users {
		names[i] = fmt.Sprintf("Name: %s\t Age: %d", user.Name, user.Age)
	}
	return names
}

func usersdb(c echo.Context) error {
	r := redisDB{}
	users := formatUsers(r.getUsers())

	return c.String(http.StatusOK, strings.Join(users, "\n"))
}

func deleteUser(c echo.Context) error {
	r := redisDB{}
	user := user{}
	err := c.Bind(&user)
	if err != nil {
		return err
	}
	usr, err := r.deleteUser(user.Name)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error deleting user")
	}
	return c.String(http.StatusOK, fmt.Sprintf("Name: %s\t\t Age: %d", usr.Name, usr.Age))
}

func GetAllUsers(c echo.Context) error {
	var redis redisDB
	var mySQL mySQLDB
	var mongo mongoDB
	dbs := []db{mongo, mySQL, redis}
	users := getAllUsers(dbs)
	return c.JSON(http.StatusOK, users)
}

type db interface {
	getUsers() []user
	deleteUser(name string) (user, error)
}

func getAllUsers(dbs []db) []user {
	var us []user

	for _, db := range dbs {
		users := db.getUsers()
		us = append(users, users...)
	}

	return us
}

func deleteUsers(c echo.Context) error{
	var redis redisDB
	var mySQL mySQLDB
	var mongo mongoDB
	dbs := []db{mongo, mySQL, redis}
	var users []user

	var user user
	err := c.Bind(&user)
	if err != nil {
		return err
	}

	for _, db := range dbs {
		user, err := db.deleteUser(user.Name)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error deleting user")
		}	
		users = append(users, user)
	}

	return c.JSON(http.StatusOK, users)
}
