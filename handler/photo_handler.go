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

type PhotoHandler interface {
	PostPhotoHandler(ctx *gin.Context)
	GetPhotosHandler(ctx *gin.Context)
	GetPhotoHandler(ctx *gin.Context)
	PutPhotoHandler(ctx *gin.Context)
	DeletePhotoHandler(ctx *gin.Context)
}

type photoHandler struct {
	photoService service.PhotoService
}

// Deletephoto godoc
// @Summary Delete photo identified by the given id
// @Description Delete the photo corresponding to the input Id
// @Tags photo
// @Accept json
// @Produce json
// @Param id path int true "ID of the photo to be deleted"
// @Security JWT
// @Success 200 {object} helper.SuccessResult{data=entity.Photo,code=int,message=string}
// @Failure 400 {object} helper.BadRequest{code=int,message=string}
// @Success 500 {object} helper.InternalServerError{code=int,message=string}
// @Router /photo/{id} [delete]
// DeletePhotoHandler implements PhotoHandler
func (h *photoHandler) DeletePhotoHandler(ctx *gin.Context) {
	var photo entity.Photo
	requestParam := ctx.Param("id")
	photoId, _ := strconv.Atoi(requestParam)
	photo.ID = uint(photoId)

	h.photoService.Delete(photo)
	helper.SuccessResponse(ctx, http.StatusOK, nil)
}


// Getphoto godoc
// @Summary Get photo details for a given id
// @Description Get details of photo corresponding is the input Id
// @Tags photo
// @Accept json
// @Produce json
// @Param id path int true "ID of the photo"
// @Security JWT
// @Success 200 {object} helper.SuccessResult{data=entity.Photo,code=int,message=string}
// @Failure 400 {object} helper.BadRequest{code=int,message=string}
// @Success 500 {object} helper.InternalServerError{code=int,message=string}
// @Router /photo/{id} [get]
// GetPhotoHandler implements PhotoHandler
func (h *photoHandler) GetPhotoHandler(ctx *gin.Context) {
	requestParam := ctx.Param("id")
	photoId, _ := strconv.Atoi(requestParam)

	photo, err := h.photoService.GetById(uint(photoId))
	if err != nil {
		helper.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.SuccessResponse(ctx, http.StatusOK, photo)

}

// Getphotos godoc
// @Summary Get all photos
// @Description Get all photos
// @Tags photo
// @Accept json
// @Produce json
// @Security JWT
// @Success 200 {object} helper.SuccessResult{data=entity.Photo,code=int,message=string}
// @Failure 400 {object} helper.BadRequest{code=int,message=string}
// @Success 500 {object} helper.InternalServerError{code=int,message=string}
// @Router /photo [get]
// GetPhotosHandler implements PhotoHandler
func (h *photoHandler) GetPhotosHandler(ctx *gin.Context) {
	photos, err := h.photoService.GetAll()
	if err != nil {
		helper.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.SuccessResponse(ctx, http.StatusOK, photos)
}

// CreatePhoto godoc
// @Summary Create photo
// @Description Create new photo
// @Tags photo
// @Accept json
// @Produce json
// @Param photo body entity.PhotoRequest true "create photo"
// @Security JWT
// @Success 201 {object} helper.SuccessResult{data=entity.Photo,code=int,message=string}
// @Failure 400 {object} helper.BadRequest{code=int,message=string}
// @Success 500 {object} helper.InternalServerError{code=int,message=string}
// @Router /photo [post]
// PostPhotoHandler implements PhotoHandler
func (h *photoHandler) PostPhotoHandler(ctx *gin.Context) {
	var payload entity.PhotoRequest

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		helper.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	photo, err := h.photoService.Create(payload, userID)
	if err != nil {
		helper.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.SuccessResponse(ctx, http.StatusCreated, photo)

}

// UpdatePhoto godoc
// @Summary Update photo identified by the given id
// @Description Update the photo corresponding to the input id
// @Tags photo
// @Accept json
// @Produce json
// @Param photo body entity.PhotoRequest true "create photo"
// @Param id path int true "ID of the photo"
// @Security JWT
// @Success 200 {object} helper.SuccessResult{data=entity.Photo,code=int,message=string}
// @Failure 400 {object} helper.BadRequest{code=int,message=string}
// @Success 500 {object} helper.InternalServerError{code=int,message=string}
// @Router /photo/{id} [put]
// PutPhotoHandler implements PhotoHandler
func (h *photoHandler) PutPhotoHandler(ctx *gin.Context) {
	requestParam := ctx.Param("id")
	photoId, _ := strconv.Atoi(requestParam)

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	var payload entity.PhotoRequest
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		helper.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	photo, err := h.photoService.Update(payload, uint(photoId), userID)
	if err != nil {
		helper.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.SuccessResponse(ctx, http.StatusOK, photo)

}

func NewPhotoHandler(photoService service.PhotoService) PhotoHandler {
	return &photoHandler{photoService: photoService}
}
