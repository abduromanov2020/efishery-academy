package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Dunia !")
	})
	e.GET("/hostname", func(c echo.Context) error {
		hostname, err := os.Hostname()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		return c.String(http.StatusOK, hostname)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
