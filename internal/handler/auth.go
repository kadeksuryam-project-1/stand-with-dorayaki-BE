package handler

import (
	"backend/config"
	"backend/internal/schema"
	"backend/internal/service"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type IAuthHandler interface {
	GoogleOAuth(c echo.Context) error
}

type authHandler struct {
	authService service.IAuthService
}

func NewAuthHandler(authService service.IAuthService) IAuthHandler {
	return &authHandler{
		authService: authService,
	}
}

func (h *authHandler) GoogleOAuth(c echo.Context) error {
	code := c.QueryParam("code")
	state := c.QueryParam("state")

	pathURL := "/"

	if state != "" {
		pathURL = state
	}

	if code == "" {
		return c.JSON(http.StatusUnauthorized, schema.NewError("Authorization code not provided!"))
	}

	accessToken, err := h.authService.GoogleOAuth(code, pathURL)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, schema.NewError(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, schema.NewError(err.Error()))
	}

	c.SetCookie(&http.Cookie{
		Name:     "access_token",
		Value:    *accessToken,
		MaxAge:   config.C.AccessTokenMaxAge * 60,
		Path:     "/",
		Domain:   config.C.CookieDomain,
		HttpOnly: true,
	})

	c.SetCookie(&http.Cookie{
		Name:     "logged_in",
		Value:    "true",
		MaxAge:   config.C.AccessTokenMaxAge * 60,
		Path:     "/",
		Domain:   config.C.CookieDomain,
		HttpOnly: false,
	})

	return c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s%s", config.C.ClientOrigin, pathURL))
}
