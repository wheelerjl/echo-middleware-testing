package main

import (
	"net/http"
	"strconv"

	"github.com/wheelerjl/personal-cnr/temp/middleware-test/internal/data"

	"github.com/labstack/echo/v4"
)

func main() {
	echo := echo.New()
	echo.GET("/fruits-db/:id", getFruit, data.ResponseMiddleware(data.DBFruit{}))
	echo.GET("/fruits-web/:id", getFruit, data.ResponseMiddleware(data.WebFruit{}))
	echo.Logger.Fatal(echo.Start(":1323"))
}

func getFruit(c echo.Context) error {
	storage := data.GetStorage()
	id, _ := strconv.Atoi(c.Param("id"))

	if id == 0 {
		// Trigger internal error in middleware
		return nil
	}

	var fruit data.DBFruit
	for _, value := range storage {
		if value.ID == id {
			fruit = value
			break
		}
	}

	if fruit.ID == 0 {
		c.Set(data.EchoCtxRespKey, data.Response{
			Code: http.StatusNotFound,
			Data: data.InternalError{
				Code:    "cnrerror-0001",
				Message: "Not found",
			}})
		return nil
	}

	c.Set(data.EchoCtxRespKey, data.Response{
		Code: http.StatusOK,
		Data: fruit})
	return nil
}
