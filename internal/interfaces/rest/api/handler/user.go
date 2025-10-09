package handler

import (
	model "auth-system/internal/core/user"
	"auth-system/internal/usecase/user"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUseCase user.UserServiceImpl
}

func NewUserHandler(userUseCase user.UserServiceImpl) UserHandler {
	return UserHandler{
		UserUseCase: userUseCase,
	}
}

func (h *UserHandler) RegisterHandler(c *gin.Context) {
	var userData *model.User
	err := c.ShouldBindJSON(&userData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "something went wrong",
			"status": http.StatusInternalServerError,
			"data":   nil,
		})
		return
	}

	err = h.UserUseCase.RegisterUser(context.Background(), userData)
	if err != nil {
		handlerError(c, err)
		return
	}
	token, err := model.GenerateJWT(userData.Email, 5)
	h.UserUseCase.SendMail(context.Background(),userData,token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "something went wrong",
			"status": http.StatusInternalServerError,
			"data":   nil,
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "user registered successfully",
		"status":  http.StatusOK,
		"data":    nil,
	})
	return

}

func (h *UserHandler) LoginHandler(c *gin.Context) {
	var user *model.LoginRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "invalid request body",
			"status": http.StatusBadRequest,
			"data":   nil,
		})
		return
	}

	data, token, err := h.UserUseCase.LoginUser(context.Background(), user.Email, user.Password)
	if err != nil {
		handlerError(c, err)
		return
	}

	c.SetCookie(
		"authToken",
		token,
		1800,
		"/",
		"",
		true,
		true,
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "user logged in successfully",
		"status":  http.StatusOK,
		"data":    data,
	})
	return
}
