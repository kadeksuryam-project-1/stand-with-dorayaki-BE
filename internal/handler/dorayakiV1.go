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

type IDorayakiHandler interface {
	CreateDorayaki(c echo.Context) error
	GetDorayakis(c echo.Context) error
	GetDorayaki(c echo.Context) error
	UpdateDorayaki(c echo.Context) error
	DeleteDorayaki(c echo.Context) error
}

type dorayakiHandler struct {
	dorayakiService service.IDorayakiService
}

func NewDorayakiHandler(dorayakiService service.IDorayakiService) IDorayakiHandler {
	return &dorayakiHandler{
		dorayakiService: dorayakiService,
	}
}

func (h *dorayakiHandler) parseIDParam(c echo.Context) (int, error) {
	idParam := c.Param("id")
	return strconv.Atoi(idParam)
}

// CreateDorayaki create Dorayaki
//
//		@Summary      Create Dorayaki
//		@Tags         dorayakis
//	    @Param flavor formData string true "flavor"
//	    @Param description formData string true "description"
//	    @Param image formData file false "image"
//		@Accept       mpfd
//		@Produce      json
//		@Success      200  {object} schema.CreateDorayakiResponseDTO
//		@Failure 	  400  {object} schema.Error
//		@Failure	  500  {object} schema.Error
//		@Router       /v1/dorayakis [post]
func (h *dorayakiHandler) CreateDorayaki(c echo.Context) error {
	var form schema.DorayakiForm

	form.Flavor = c.FormValue("flavor")
	form.Description = c.FormValue("description")
	form.Image, _ = c.FormFile("image")

	if err := c.Validate(&form); err != nil {
		return c.JSON(http.StatusBadRequest, schema.NewError(err.Error()))
	}
	newDorayaki, err := h.dorayakiService.CreateDorayaki(form)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, schema.NewError(err.Error()))
	}

	return c.JSON(http.StatusOK, schema.CreateDorayakiResponseDTO{
		Response: schema.Response{
			Success: true,
		},
		Data: newDorayaki,
	})
}

// GetDorayakis get Dorayakis
//
//		@Summary      Get Dorayakis
//		@Tags         dorayakis
//		@Produce      json
//		@Success      200  {object} schema.GetDorayakisResponseDTO
//	    @Failure      500  {object} schema.Error
//		@Router       /v1/dorayakis [get]
func (h *dorayakiHandler) GetDorayakis(c echo.Context) error {
	dorayakis, err := h.dorayakiService.GetDorayakis()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, schema.NewError(err.Error()))
	}

	return c.JSON(http.StatusOK, schema.GetDorayakisResponseDTO{
		Response: schema.Response{
			Success: true,
		},
		Data: dorayakis,
	})
}

// GetDorayaki get Dorayaki
//
//		@Summary      Get Dorayaki
//		@Tags         dorayakis
//	    @Param id path int true "id"
//		@Produce      json
//		@Success      200  {object} schema.GetDorayakiResponseDTO
//		@Failure	  400  {object} schema.Error
//		@Failure	  404  {object} schema.Error
//	    @Failure      500  {object} schema.Error
//		@Router       /v1/dorayakis/{id} [get]
func (h *dorayakiHandler) GetDorayaki(c echo.Context) error {
	p := c.Param("id")
	var id int
	id, err := strconv.Atoi(p)

	if err != nil {
		return c.JSON(http.StatusBadRequest, schema.NewError(err.Error()))
	}
	dorayaki, err := h.dorayakiService.GetDorayaki(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, schema.NewError(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, schema.NewError(err.Error()))
	}

	return c.JSON(http.StatusOK, schema.GetDorayakiResponseDTO{
		Response: schema.Response{
			Success: true,
		},
		Data: dorayaki,
	})
}

// UpdateDorayaki update Dorayaki
//
//		@Summary      Update Dorayaki
//		@Tags         dorayakis
//	    @Param id path int true "id"
//	    @Param flavor formData string true "flavor"
//	    @Param description formData string true "description"
//	    @Param image formData file false "image"
//		@Accept       mpfd
//		@Produce      json
//		@Success      200  {object} schema.UpdateDorayakiResponseDTO
//		@Failure	  400  {object} schema.Error
//		@Failure	  404  {object} schema.Error
//	    @Failure      500  {object} schema.Error
//		@Router       /v1/dorayakis/{id} [put]
func (h *dorayakiHandler) UpdateDorayaki(c echo.Context) error {
	id, err := h.parseIDParam(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, schema.NewError(err.Error()))
	}

	var form schema.DorayakiForm
	form.Flavor = c.FormValue("flavor")
	form.Description = c.FormValue("description")
	form.Image, _ = c.FormFile("image")

	if err := c.Validate(&form); err != nil {
		return c.JSON(http.StatusBadRequest, schema.NewError(err.Error()))
	}

	updatedDorayaki, err := h.dorayakiService.UpdateDorayaki(form, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, schema.NewError(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, schema.NewError(err.Error()))
	}

	return c.JSON(http.StatusOK, schema.UpdateDorayakiResponseDTO{
		Response: schema.Response{
			Success: true,
		},
		Data: updatedDorayaki,
	})
}

// DeleteDorayaki delete Dorayaki
//
//		@Summary      Delete Dorayaki
//		@Tags         dorayakis
//	    @Param id path int true "id"
//		@Produce      json
//		@Success      200  {object} schema.DeleteDorayakiResponseDTO
//		@Failure	  400  {object} schema.Error
//		@Failure	  404  {object} schema.Error
//	    @Failure      500  {object} schema.Error
//		@Router       /v1/dorayakis/{id} [delete]
func (h *dorayakiHandler) DeleteDorayaki(c echo.Context) error {
	id, err := h.parseIDParam(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, schema.NewError(err.Error()))
	}

	err = h.dorayakiService.DeleteDorayaki(id)
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
