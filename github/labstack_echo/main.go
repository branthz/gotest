package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var e *echo.Echo

func main() {
	// Echo instance
	e = echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/json", jsont)
	e.POST("/p", post)
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

type (
	user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

func post(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	e.Logger.Info(u)
	return c.JSON(200, `{code:0,msg:ok}`)
}

func jsont(c echo.Context) error {
	return c.JSON(200, "{\"code\":0,\"msg\":\"success\"}")
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
