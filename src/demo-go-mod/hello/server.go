package main

import (
	"demo-go-mod/hello/api"
	"github.com/labstack/echo"
)

/*
Author: kolly.li@klook.com
Date: 2019/9/29
*/

func main() {
	e := echo.New()
	e.GET("/", api.HelloWorld)
	e.Logger.Fatal(e.Start(":1323"))
}
