package api

import (
	"github.com/labstack/echo"
	"net/http"
)

/*
Author: kolly.li@klook.com
Date: 2019/9/29
*/

func HelloWorld(c echo.Context) error {
	return c.JSON(http.StatusOK, "hello world")
}
