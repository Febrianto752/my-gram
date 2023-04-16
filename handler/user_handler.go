package handler

import (
	"net/http"

	"github.com/Febrianto752/my-gram/entity"
	"github.com/Febrianto752/my-gram/helper"
	"github.com/Febrianto752/my-gram/service"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	PostUserRegisterHandler(ctx *gin.Context)
	PostUserLoginHandler(ctx *gin.Context)
}

type userHandler struct {
	userService service.UserService
}

// PostUserLoginHandler implements UserHandler
func (h *userHandler) PostUserLoginHandler(ctx *gin.Context) {
	var payload entity.UserLogin

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {

		helper.FailResponse(ctx, http.StatusBadRequest, err.Error())
	}

	loggedInUser, err := h.userService.Login(payload)
	if err != nil {
		helper.FailResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	token := helper.GenerateToken(loggedInUser.ID, loggedInUser.Email)
	helper.SuccessResponse(ctx, http.StatusOK, gin.H{
		"access_token": token,
	})
}

// PostUserRegisterHandler implements UserHandler
func (h *userHandler) PostUserRegisterHandler(ctx *gin.Context) {
	var payload entity.UserRequest

	err := ctx.ShouldBindJSON(&payload)

	if err != nil {
		helper.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	newUser, err := h.userService.Register(payload)
	if err != nil {

		helper.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.SuccessResponse(ctx, http.StatusCreated, gin.H{
		"id":       newUser.ID,
		"email":    newUser.Email,
		"username": newUser.Username,
	})
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{userService: userService}
}
