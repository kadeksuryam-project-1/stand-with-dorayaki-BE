package handler

import (
	"backend/internal/schema"
	"backend/internal/service"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type IDorayakiStoreStockHandler interface {
	GetStocks(c echo.Context) error
	UpdateStock(c echo.Context) error
}

type dorayakiStoreStockHandler struct {
	dorayakiStoreStockService service.IDorayakiStoreStockService
}

func NewDorayakiStoreStockHandler(dorayakiStoreStockService service.IDorayakiStoreStockService) IDorayakiStoreStockHandler {
	return &dorayakiStoreStockHandler{
		dorayakiStoreStockService: dorayakiStoreStockService,
	}
}

// GetStocks get stocks
//
//		@Summary      Get Stocks
//		@Tags         stocks
//	    @Param dorayaki_id query int false "dorayaki_id"
//	    @Param store_id query int false "store_id"
//		@Produce      json
//		@Success      200  {object} schema.UpdateStockResponseDTO
//		@Failure	  400  {object} schema.Error
//	    @Failure      500  {object} schema.Error
//		@Router       /v1/stocks [get]
func (h *dorayakiStoreStockHandler) GetStocks(c echo.Context) error {
	dorayakiIDParam := c.QueryParam("dorayaki_id")
	storeIDParam := c.QueryParam("store_id")

	var dorayakiID, storeID *int

	if dorayakiIDParam != "" {
		id, err := strconv.Atoi(dorayakiIDParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, schema.NewError(err.Error()))
		}
		dorayakiID = &id
	}

	if storeIDParam != "" {
		id, err := strconv.Atoi(storeIDParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, schema.NewError(err.Error()))
		}
		storeID = &id
	}

	stockDTO := new(schema.StockRequestDTO)
	if err := c.Bind(stockDTO); err != nil {
		return c.JSON(http.StatusBadRequest, schema.NewError(err.Error()))
	}
	if err := c.Validate(stockDTO); err != nil {
		return c.JSON(http.StatusBadRequest, schema.NewError(err.Error()))
	}

	stocks, err := h.dorayakiStoreStockService.GetStocks(dorayakiID, storeID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, schema.NewError(err.Error()))
	}

	return c.JSON(http.StatusOK, schema.GetStocksResponseDTO{
		Response: schema.Response{
			Success: true,
		},
		Data: stocks,
	})
}

// UpdateStock update stock
//
//		@Summary      Update Stock
//		@Tags         stocks
//	    @Param id path int true "id"
//	    @Param stock body schema.StockRequestDTO true "stock"
//		@Accept       json
//		@Produce      json
//		@Success      200  {object} schema.UpdateStockResponseDTO
//		@Failure	  400  {object} schema.Error
//		@Failure	  404  {object} schema.Error
//	    @Failure      500  {object} schema.Error
//		@Router       /v1/stocks/{id} [patch]
func (h *dorayakiStoreStockHandler) UpdateStock(c echo.Context) error {
	id, err := h.parseIDParam(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, schema.NewError(err.Error()))
	}

	stockDTO := new(schema.StockRequestDTO)
	if err := c.Bind(stockDTO); err != nil {
		return c.JSON(http.StatusBadRequest, schema.NewError(err.Error()))
	}
	if err := c.Validate(stockDTO); err != nil {
		return c.JSON(http.StatusBadRequest, schema.NewError(err.Error()))
	}

	stockItem, err := h.dorayakiStoreStockService.UpdateStock(stockDTO.Stock, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, schema.NewError(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, schema.NewError(err.Error()))
	}

	return c.JSON(http.StatusOK, schema.UpdateStockResponseDTO{
		Response: schema.Response{
			Success: true,
		},
		Data: stockItem,
	})
}

func (h *dorayakiStoreStockHandler) parseIDParam(c echo.Context) (int, error) {
	idParam := c.Param("id")
	return strconv.Atoi(idParam)
}
