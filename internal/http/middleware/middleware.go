package middleware

import "github.com/labstack/echo/v4"

type responseError struct {
	Message string `json:"message"`
}

func CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return next(c)
	}
}
