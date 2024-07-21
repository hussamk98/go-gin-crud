package handler

import (
	"net/http"
	"strconv"

	"crud-app/internal/app/crud/dto"
	"crud-app/internal/app/crud/models"
	"crud-app/internal/app/crud/service"
	"crud-app/pkg/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var userDTO dto.CreateUserDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		utils.NewErrorHandler().BadRequest(c, err)
		return
	}

	user := models.NewUserFromDTO(userDTO)

	createdUser, err := h.userService.CreateUser(user)
	if err != nil {
		utils.NewErrorHandler().InternalServerError(c, err)
		return
	}
	c.JSON(http.StatusCreated, createdUser)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.userService.GetUserByID(uint(id))
	if err != nil {
		utils.NewErrorHandler().NotFound(c)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetUserByUsername(c *gin.Context) {
	username := c.Param("username")
	user, err := h.userService.GetUserByUsername(username)
	if err != nil {
		utils.NewErrorHandler().NotFound(c)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		utils.NewErrorHandler().InternalServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.NewErrorHandler().BadRequest(c, err)
		return
	}
	updatedUser, err := h.userService.UpdateUser(user)
	if err != nil {
		utils.NewErrorHandler().InternalServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, updatedUser)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.userService.DeleteUser(uint(id))
	if err != nil {
		utils.NewErrorHandler().InternalServerError(c, err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
