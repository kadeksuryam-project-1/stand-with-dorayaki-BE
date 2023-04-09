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

type IStoreHandler interface {
	CreateStore(c echo.Context) error
	GetStores(c echo.Context) error
	GetStore(c echo.Context) error
	UpdateStore(c echo.Context) error
	DeleteStore(c echo.Context) error
}

type storeHandler struct {
	storeService service.IStoreService
}

func NewStoreHandler(storeService service.IStoreService) IStoreHandler {
	return &storeHandler{
		storeService: storeService,
	}
}

func (h *storeHandler) parseIDParam(c echo.Context) (int, error) {
	idParam := c.Param("id")
	return strconv.Atoi(idParam)
}

// CreateStore create store
//
//		@Summary      Create store
//		@Tags         stores
//	    @Param name formData string true "name"
//	    @Param street formData string true "street"
//	    @Param subdistrict formData string true "subdistrict"
//	    @Param district formData string true "district"
//	    @Param province formData string true "province"
//	    @Param image formData file false "image"
//		@Accept       mpfd
//		@Produce      json
//		@Success      200  {object} schema.CreateStoreResponseDTO
//		@Failure 	  400  {object} schema.Error
//		@Failure	  500  {object} schema.Error
//		@Router       /v1/stores [post]
func (h *storeHandler) CreateStore(c echo.Context) error {
	var form schema.StoreForm

	form.Name = c.FormValue("name")
	form.Street = c.FormValue("street")
	form.Subdistrict = c.FormValue("subdistrict")
	form.District = c.FormValue("district")
	form.Province = c.FormValue("province")
	form.Image, _ = c.FormFile("image")

	if err := c.Validate(&form); err != nil {
		return c.JSON(http.StatusBadRequest, schema.NewError(err.Error()))
	}
	newStore, err := h.storeService.CreateStore(form)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, schema.NewError(err.Error()))
	}

	return c.JSON(http.StatusOK, schema.CreateStoreResponseDTO{
		Response: schema.Response{
			Success: true,
		},
		Data: newStore,
	})
}

// GetStores get stores
//
//		@Summary      Get stores
//		@Tags         stores
//		@Produce      json
//		@Success      200  {object} schema.GetStoresResponseDTO
//	    @Failure      500  {object} schema.Error
//		@Router       /v1/stores [get]
func (h *storeHandler) GetStores(c echo.Context) error {
	stores, err := h.storeService.GetStores()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, schema.NewError(err.Error()))
	}

	return c.JSON(http.StatusOK, schema.GetStoresResponseDTO{
		Response: schema.Response{
			Success: true,
		},
		Data: stores,
	})
}

// GetStore get store
//
//		@Summary      Get store
//		@Tags         stores
//	    @Param id path int true "id"
//		@Produce      json
//		@Success      200  {object} schema.GetStoreResponseDTO
//		@Failure	  400  {object} schema.Error
//		@Failure	  404  {object} schema.Error
//	    @Failure      500  {object} schema.Error
//		@Router       /v1/stores/{id} [get]
func (h *storeHandler) GetStore(c echo.Context) error {
	p := c.Param("id")
	var id int
	id, err := strconv.Atoi(p)

	if err != nil {
		return c.JSON(http.StatusBadRequest, schema.NewError(err.Error()))
	}
	store, err := h.storeService.GetStore(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, schema.NewError(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, schema.NewError(err.Error()))
	}

	return c.JSON(http.StatusOK, schema.GetStoreResponseDTO{
		Response: schema.Response{
			Success: true,
		},
		Data: store,
	})
}

// UpdateStore update store
//
//		@Summary      Update store
//		@Tags         stores
//	    @Param id path int true "id"
//	    @Param name formData string true "name"
//	    @Param street formData string true "street"
//	    @Param subdistrict formData string true "subdistrict"
//	    @Param district formData string true "district"
//	    @Param province formData string true "province"
//	    @Param image formData file false "image"
//		@Accept       mpfd
//		@Produce      json
//		@Success      200  {object} schema.UpdateStoreResponseDTO
//		@Failure	  400  {object} schema.Error
//		@Failure	  404  {object} schema.Error
//	    @Failure      500  {object} schema.Error
//		@Router       /v1/stores/{id} [put]
func (h *storeHandler) UpdateStore(c echo.Context) error {
	id, err := h.parseIDParam(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, schema.NewError(err.Error()))
	}

	var form schema.StoreForm
	form.Name = c.FormValue("name")
	form.Street = c.FormValue("street")
	form.Subdistrict = c.FormValue("subdistrict")
	form.District = c.FormValue("district")
	form.Province = c.FormValue("province")
	form.Image, _ = c.FormFile("image")

	if err := c.Validate(&form); err != nil {
		return c.JSON(http.StatusBadRequest, schema.NewError(err.Error()))
	}

	updatedStore, err := h.storeService.UpdateStore(form, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, schema.NewError(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, schema.NewError(err.Error()))
	}

	return c.JSON(http.StatusOK, schema.UpdateStoreResponseDTO{
		Response: schema.Response{
			Success: true,
		},
		Data: updatedStore,
	})
}

// DeleteStore delete store
//
//		@Summary      Delete store
//		@Tags         stores
//	    @Param id path int true "id"
//		@Produce      json
//		@Success      200  {object} schema.DeleteStoreResponseDTO
//		@Failure	  400  {object} schema.Error
//		@Failure	  404  {object} schema.Error
//	    @Failure      500  {object} schema.Error
//		@Router       /v1/stores/{id} [delete]
func (h *storeHandler) DeleteStore(c echo.Context) error {
	id, err := h.parseIDParam(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, schema.NewError(err.Error()))
	}

	err = h.storeService.DeleteStore(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, schema.NewError(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, schema.NewError(err.Error()))
	}

	return c.JSON(http.StatusOK, schema.DeleteDorayakiResponseDTO{
		Response: schema.Response{
			Success: true,
		},
	})
}
