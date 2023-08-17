package handler

import (
	"alghanim/mediacmsAPI/model"
	"alghanim/mediacmsAPI/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Get(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, _ := h.userService.Get(id)

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Create(c echo.Context) error {
	// TODO: Implement
	return nil
}

func (h *UserHandler) Update(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	updateUser, err := h.userService.Update(id, user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, updateUser)
}
