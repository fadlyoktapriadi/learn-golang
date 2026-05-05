package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type M map[string]interface{}

func main() {
	e := echo.New()

	e.GET("/", func(ctx echo.Context) error {
		data := "Hello from /index"
		return ctx.String(http.StatusOK, data)
	})

	e.Start(":9000")
}
