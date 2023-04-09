package handler

import (
	"backend/internal/repository"
	"backend/internal/service"
	"net/http"

	_ "backend/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"
)

func RoutesInit(server *echo.Echo, db *gorm.DB) {
	dorayakiRepository := repository.NewDorayakiRepository(db)
	dorayakiService := service.NewDorayakiService(dorayakiRepository)
	dorayakiHandler := NewDorayakiHandler(dorayakiService)
	storeRepository := repository.NewStoreRepository(db)
	storeService := service.NewStoreService(storeRepository)
	storeHandler := NewStoreHandler(storeService)
	dorayakiStoreStockRepository := repository.NewDorayakiStoreStockRepository(db)
	dorayakiStoreStockService := service.NewDorayakiStoreStockService(dorayakiStoreStockRepository)
	dorayakiStoreStockHandler := NewDorayakiStoreStockHandler(dorayakiStoreStockService)

	api := server.Group("/api")
	api.GET("/healthz", HealthCheck)
	api.GET("/swagger/*", echoSwagger.WrapHandler)

	v1 := api.Group("/v1")

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
	dorayakiStoreStocks.PATCH("", dorayakiStoreStockHandler.UpdateStock)
}

func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Server is up and running",
	})
}
