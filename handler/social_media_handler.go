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

// DeleteSocialMedia godoc
// @Summary Delete social media identified by the given id
// @Description Delete the social media corresponding to the input Id
// @Tags social media
// @Accept json
// @Produce json
// @Param id path int true "ID of the social media to be deleted"
// @Security JWT
// @Success 200 {object} helper.SuccessResult{data=entity.SocialMedia,code=int,message=string}
// @Failure 400 {object} helper.BadRequest{code=int,message=string}
// @Success 500 {object} helper.InternalServerError{code=int,message=string}
// @Router /socialmedia/{id} [delete]
// DeleteSocialMediaHandler implements SocialMediaHandler
func (h *socialMediaHandler) DeleteSocialMediaHandler(ctx *gin.Context) {
	var socialMedia entity.SocialMedia
	requestParam := ctx.Param("id")
	socialMediaId, _ := strconv.Atoi(requestParam)
	socialMedia.ID = uint(socialMediaId)

	h.socialMediaService.Delete(socialMedia)
	helper.SuccessResponse(ctx, http.StatusOK, nil)
}

// GetSocialMedia godoc
// @Summary Get details for a given id
// @Description Get details of social media corresponding is the input Id
// @Tags social media
// @Accept json
// @Produce json
// @Param id path int true "ID of the social media"
// @Security JWT
// @Success 200 {object} helper.SuccessResult{data=entity.SocialMedia,code=int,message=string}
// @Failure 400 {object} helper.BadRequest{code=int,message=string}
// @Success 500 {object} helper.InternalServerError{code=int,message=string}
// @Router /socialmedia/{id} [get]
// GetSocialMediaHandler implements SocialMediaHandler
func (h *socialMediaHandler) GetSocialMediaHandler(ctx *gin.Context) {
	requestParam := ctx.Param("id")
	socialMediaId, _ := strconv.Atoi(requestParam)

	socialMedia, err := h.socialMediaService.GetById(uint(socialMediaId))
	if err != nil {
		helper.FailResponse(ctx, http.StatusBadRequest, err.Error())
	}

	helper.SuccessResponse(ctx, http.StatusOK, socialMedia)

}

// GetSocialMedias godoc
// @Summary Get all social medias
// @Description Get all social medias
// @Tags social media
// @Accept json
// @Produce json
// @Security JWT
// @Success 200 {object} helper.SuccessResult{data=entity.SocialMedia,code=int,message=string}
// @Failure 400 {object} helper.BadRequest{code=int,message=string}
// @Success 500 {object} helper.InternalServerError{code=int,message=string}
// @Router /socialmedia [get]
// GetSocialMediasHandler implements SocialMediaHandler
func (h *socialMediaHandler) GetSocialMediasHandler(ctx *gin.Context) {
	socialMedias, err := h.socialMediaService.GetAll()
	if err != nil {
		helper.FailResponse(ctx, http.StatusBadRequest, err.Error())
	}

	helper.SuccessResponse(ctx, http.StatusOK, socialMedias)
}

// CreateSocialMedia godoc
// @Summary Create social media
// @Description Create new social media
// @Tags social media
// @Accept json
// @Produce json
// @Param socialMedia body entity.SocialMediaRequest true "create social media"
// @Security JWT
// @Success 201 {object} helper.SuccessResult{data=entity.SocialMedia,code=int,message=string}
// @Failure 400 {object} helper.BadRequest{code=int,message=string}
// @Success 500 {object} helper.InternalServerError{code=int,message=string}
// @Router /socialmedia [post]
// PostSocialMediaHandler implements SocialMediaHandler
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

// UpdateSocialMedia godoc
// @Summary Update social media identified by the given id
// @Description Update the social media corresponding to the input id
// @Tags social media
// @Accept json
// @Produce json
// @Param id path int true "ID of the social media"
// @Param socialMedia body entity.SocialMediaRequest true "update social media"
// @Security JWT
// @Success 200 {object} helper.SuccessResult{data=entity.SocialMedia,code=int,message=string}
// @Failure 400 {object} helper.BadRequest{code=int,message=string}
// @Success 500 {object} helper.InternalServerError{code=int,message=string}
// @Router /socialmedia/{id} [put]
// PutSocialMediaHandler implements SocialMediaHandler
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
