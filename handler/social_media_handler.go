package handler

import (
	"net/http"
	"strconv"

	"github.com/Febrianto752/my-gram/entity"
	"github.com/Febrianto752/my-gram/helper"
	"github.com/Febrianto752/my-gram/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type SocialMediaHandler interface {
	PostSocialMediaHandler(ctx *gin.Context)
	GetSocialMediasHandler(ctx *gin.Context)
	GetSocialMediaHandler(ctx *gin.Context)
	PutSocialMediaHandler(ctx *gin.Context)
	DeleteSocialMediaHandler(ctx *gin.Context)
}

type socialMediaHandler struct {
	socialMediaService service.SocialMediaService
}

func (h *socialMediaHandler) DeleteSocialMediaHandler(ctx *gin.Context) {
	var socialMedia entity.SocialMedia
	requestParam := ctx.Param("id")
	socialMediaId, _ := strconv.Atoi(requestParam)
	socialMedia.ID = uint(socialMediaId)

	h.socialMediaService.Delete(socialMedia)
	helper.SuccessResponse(ctx, http.StatusOK, nil)
}

func (h *socialMediaHandler) GetSocialMediaHandler(ctx *gin.Context) {
	requestParam := ctx.Param("id")
	socialMediaId, _ := strconv.Atoi(requestParam)

	socialMedia, err := h.socialMediaService.GetById(uint(socialMediaId))
	if err != nil {
		helper.FailResponse(ctx, http.StatusBadRequest, err.Error())
	}

	helper.SuccessResponse(ctx, http.StatusOK, socialMedia)

}

func (h *socialMediaHandler) GetSocialMediasHandler(ctx *gin.Context) {
	socialMedias, err := h.socialMediaService.GetAll()
	if err != nil {
		helper.FailResponse(ctx, http.StatusBadRequest, err.Error())
	}

	helper.SuccessResponse(ctx, http.StatusOK, socialMedias)
}

func (h *socialMediaHandler) PostSocialMediaHandler(ctx *gin.Context) {
	var payload entity.SocialMediaRequest

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		helper.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	socialMedia, err := h.socialMediaService.Create(payload, userID)
	if err != nil {
		helper.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.SuccessResponse(ctx, http.StatusCreated, socialMedia)

}

func (h *socialMediaHandler) PutSocialMediaHandler(ctx *gin.Context) {
	requestParam := ctx.Param("id")
	socialMediaId, _ := strconv.Atoi(requestParam)

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	var payload entity.SocialMediaRequest
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		helper.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	socialMedia, err := h.socialMediaService.Update(payload, uint(socialMediaId), userID)
	if err != nil {
		helper.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.SuccessResponse(ctx, http.StatusOK, socialMedia)

}

func NewSocialMediaHandler(socialMediaService service.SocialMediaService) SocialMediaHandler {
	return &socialMediaHandler{socialMediaService: socialMediaService}
}
