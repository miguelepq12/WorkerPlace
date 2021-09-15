package web

import (
	"AccessValidation/app/application"
	"AccessValidation/app/domain/entity"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AccessHandler struct {
	accessUseCase *application.AccessUseCase
}
func NewAccessHandler(useCase *application.AccessUseCase) *AccessHandler{
	return &AccessHandler{accessUseCase: useCase}
}


func (receiver *AccessHandler) ValidateAccess(c echo.Context) error {
	u := entity.NewAccessInformation()
	if err := c.Bind(u); err != nil {
		return err
	}

	if err:= receiver.accessUseCase.ValidateAccess(u); err!= nil {
		return c.JSON(http.StatusForbidden, err.Error())
	}

	return c.JSON(http.StatusOK, make(map[string]int))
}
