package auth

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (c *AuthController) Login(ctx echo.Context) error {
	type AuthInfo struct {
		UserID   string
		Password string
	}
	var authInfo AuthInfo
	if err := ctx.Bind(&authInfo); err != nil {
		return err
	}

	// FIXME: DBに保存されているのユーザーID＆パスワードと比較する
	if authInfo.UserID != "aaaa" && authInfo.Password != "bbbb" {
		return fmt.Errorf("login error")
	}

	sess, err := session.Get("session", ctx)
	if err != nil {
		return err
	}
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["userID"] = authInfo.UserID
	sess.Save(ctx.Request(), ctx.Response())
	return ctx.NoContent(http.StatusOK)
}
