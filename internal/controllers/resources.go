package controllers

import (
	"net/http"

	"github.com/cyverse/QMS/internal/model"
	"github.com/labstack/echo/v4"
)

func (s Server) GetAllResources(ctx echo.Context) error {
	data := []model.ResourceType{}
	err := s.GORMDB.Debug().Find(&data).Error
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.ErrorResponse(err.Error(), http.StatusInternalServerError))
	}
	return ctx.JSON(http.StatusOK, model.SuccessResponse(data, http.StatusOK))
}
