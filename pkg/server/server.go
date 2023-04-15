package server

import (
	"backend/config"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

type IServer interface {
	App() *echo.Echo
	Start(port string)
}

type server struct {
	app *echo.Echo
}

func (s *server) App() *echo.Echo {
	return s.app
}

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func (s *server) Start(port string) {
	s.app.HideBanner = true
	log.Info().Str("port", port).Msg("starting server")
	s.app.Logger.Fatal(s.app.Start(fmt.Sprintf("0.0.0.0:%s", port)))
}

func NewServer() IServer {
	app := echo.New()
	app.Validator = &CustomValidator{Validator: validator.New()}

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     config.C.AllowOrigins,
		AllowCredentials: true,
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	return &server{
		app: app,
	}
}
