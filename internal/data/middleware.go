package data

import (
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

func ResponseMiddleware(out any) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			if err := next(ctx); err != nil {
				ctx.Error(err)
			}

			x := ctx.Get(EchoCtxRespKey)
			var response Response
			if x == nil {
				return ctx.JSON(http.StatusInternalServerError, InternalError{Code: "middleware-0001", Message: "No response set"})
			}
			response = x.(Response)

			//response := ctx.Get(EchoCtxRespKey).(Response)
			if response.Code > 399 {
				return ctx.JSON(response.Code, response.Data)
			}

			if err := copier.CopyWithOption(&out, response.Data, copier.Option{}); err != nil {
				return ctx.JSON(http.StatusInternalServerError, "unable to copy")
			}

			return ctx.JSON(http.StatusOK, out)
		}
	}
}
