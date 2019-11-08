package main

import (
	"github.com/labstack/echo"
	"github.com/prongbang/goup/pkg/file"
)

func main() {
	e := echo.New()

	e.Static("/public", "public")
	file.ProvideRoute().Initial(e)

	e.Logger.Fatal(e.Start(":4000"))
}
