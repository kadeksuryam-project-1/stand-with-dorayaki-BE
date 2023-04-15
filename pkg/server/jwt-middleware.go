package server

import (
	"backend/config"
	"backend/internal/schema"
	"backend/pkg/util"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UnsetCookie(name string) *http.Cookie {
	unsetCookie := new(http.Cookie)
	unsetCookie.Name = name
	unsetCookie.Value = ""
	unsetCookie.Path = "/"
	unsetCookie.MaxAge = -1

	return unsetCookie
}

func CheckJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := c.Cookie("logged_in")
		if err != nil {
			if err == http.ErrNoCookie {
				return c.JSON(http.StatusUnauthorized, schema.NewError("Unauthorized"))
			}
			return err
		}

		accessTokenCookie, err := c.Cookie("access_token")
		if err != nil {
			if err == http.ErrNoCookie {
				return c.JSON(http.StatusUnauthorized, schema.NewError("Unauthorized"))
			}
			return err
		}

		tokenString := accessTokenCookie.Value
		_, err = util.ValidateToken(tokenString, config.C.AccessTokenPublicKey)

		if err != nil {
			c.SetCookie(UnsetCookie("access_token"))
			c.SetCookie(UnsetCookie("logged_in"))
			return c.JSON(http.StatusUnauthorized, schema.NewError("Token expired or invalid"))
		}

		return next(c)
	}
}
