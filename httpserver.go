package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func HttpServer () {
	e := echo.New()
	e.GET("/helloworld", helloWorld)
	e.GET("/users", users)
	e.GET("/usersdb", usersdb)
	e.Logger.Fatal(e.Start(":8080"))
}

func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func users(c echo.Context) error {
	userNames := getAllUsers(GetUsers())

	return c.String(http.StatusOK, strings.Join(userNames, "\n"))
}

func getAllUsers(users []User) []string {
	names := make([]string, len(users))
	for i, user := range users {
		names[i] = fmt.Sprintf("Name: %s\t Age: %d", user.Name, user.Age)
	}
	return names
}

func usersdb(c echo.Context) error {
	r := RedisDB{}
	users := getAllUsers(r.GetUsers())

	return c.String(http.StatusOK, strings.Join(users, "\n"))
}
