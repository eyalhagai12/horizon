package handlers

import (
	"horizon/server"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HandlerFunc[Req any, Res any] func(echo.Context, server.Env, Req) (Res, error)

func FromFunc[Req any, Res any](handler HandlerFunc[Req, Res], env server.Env, successCode int) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request Req
		if err := c.Bind(&request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		res, err := handler(c, env, request)
		if err != nil {
			return err
		}

		return c.JSON(successCode, res)
	}
}
