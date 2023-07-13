package auth

import (
	"fmt"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("session", c)
		if err != nil {
			return err
		}
		userID := sess.Values["userID"]
		if userID == nil {
			return fmt.Errorf("auth error")
		}
		c.Set("userID", userID)
		return next(c)
	}
}
