package main

import (
    "net/http"
    "fmt"

    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
)

func main() {
    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.POST("/run", HandleRunPost)

    e.Logger.Fatal(e.Start(":1234"))
}

func HandleRunPost(c echo.Context) error {
    code := c.FormValue("code")
    fmt.Println(code)
    return c.JSON(http.StatusOK, map[string]interface{}{"status": "0"})
}
