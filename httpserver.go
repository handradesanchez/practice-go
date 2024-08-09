package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func httpServer () {
	e := echo.New()
	e.GET("/helloworld", helloWorld)
	e.GET("/users", users)
	e.GET("/usersdb", usersdb)
	e.DELETE("/usersdb", deleteUser)
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
	usr,err := r.deleteUsers(user.Name)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error deleting user")
	}
	return c.String(http.StatusOK, fmt.Sprintf("Name: %s\t\t Age: %d", usr.Name, usr.Age))
}
