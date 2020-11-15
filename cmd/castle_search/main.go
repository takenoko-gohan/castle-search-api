package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/takenoko-gohan/castle-search-api/internal/search"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/search", search.CastleSearch)

	e.Logger.Fatal(e.Start(":8080"))
}
