package main

import (
	"net/http"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func (c echo.Context)
