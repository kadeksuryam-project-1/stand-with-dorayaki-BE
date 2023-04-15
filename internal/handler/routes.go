package handler

import (
	"backend/internal/repository"
	"backend/internal/service"
	"net/http"

	_ "backend/docs"

	"backend/pkg/server"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"
)

func RoutesInit(s *echo.Echo, db *gorm.DB) {
	dorayakiRepository := repository.NewDorayakiRepository(db)
	dorayakiService := service.NewDorayakiService(dorayakiRepository)
	dorayakiHandler := NewDorayakiHandler(dorayakiService)
	storeRepository := repository.NewStoreRepository(db)
	storeService := service.NewStoreService(storeRepository)
	storeHandler := NewStoreHandler(storeService)
	dorayakiStoreStockRepository := repository.NewDorayakiStoreStockRepository(db)
	dorayakiStoreStockService := service.NewDorayakiStoreStockService(dorayakiStoreStockRepository)
	dorayakiStoreStockHandler := NewDorayakiStoreStockHandler(dorayakiStoreStockService)
	userRepository := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepository)
	authHandler := NewAuthHandler(authService)

	api := s.Group("/api")
	api.GET("/healthz", HealthCheck)
	api.GET("/swagger/*", echoSwagger.WrapHandler)
	api.GET("/sessions/oauth/google", authHandler.GoogleOAuth)

	v1 := api.Group("/v1")
	v1.Use(server.CheckJWT)
	v1.GET("/is-authenticated", AuthenticatedCheck)

	dorayakis := v1.Group("/dorayakis")
	dorayakis.POST("", dorayakiHandler.CreateDorayaki)
	dorayakis.GET("", dorayakiHandler.GetDorayakis)
	dorayakis.GET("/:id", dorayakiHandler.GetDorayaki)
	dorayakis.PUT("/:id", dorayakiHandler.UpdateDorayaki)
	dorayakis.DELETE("/:id", dorayakiHandler.DeleteDorayaki)

	stores := v1.Group("/stores")
	stores.POST("", storeHandler.CreateStore)
	stores.GET("", storeHandler.GetStores)
	stores.GET("/:id", storeHandler.GetStore)
	stores.PUT("/:id", storeHandler.UpdateStore)
	stores.DELETE("/:id", storeHandler.DeleteStore)

	dorayakiStoreStocks := v1.Group("/stocks")
	dorayakiStoreStocks.GET("", dorayakiStoreStockHandler.GetStocks)
	dorayakiStoreStocks.PATCH("/:id", dorayakiStoreStockHandler.UpdateStock)
}

func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Server is up and running",
	})
}

func AuthenticatedCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Authenticated",
	})
}
