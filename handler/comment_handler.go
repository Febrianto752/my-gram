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

type CommentHandler interface {
	PostCommentHandler(ctx *gin.Context)
	GetCommentsHandler(ctx *gin.Context)
	GetCommentHandler(ctx *gin.Context)
	PutCommentHandler(ctx *gin.Context)
	DeleteCommentHandler(ctx *gin.Context)
}

type commentHandler struct {
	commentService service.CommentService
}

// DeleteComment godoc
// @Summary Delete comment identified by the given id
// @Description Delete the comment corresponding to the input Id
// @Tags comment
// @Accept json
// @Produce json
// @Param id path int true "ID of the comment to be deleted"
// @Security JWT
// @Success 200 {object} helper.SuccessResult{data=entity.Comment,code=int,message=string}
// @Failure 400 {object} helper.BadRequest{code=int,message=string}
// @Success 500 {object} helper.InternalServerError{code=int,message=string}
// @Router /comment/{id} [delete]
// DeleteCommentHandler implements CommentHandler
func (h *commentHandler) DeleteCommentHandler(ctx *gin.Context) {
	var comment entity.Comment
	requestParam := ctx.Param("id")
	commentId, _ := strconv.Atoi(requestParam)
	comment.ID = uint(commentId)

	comment, err := h.commentService.Delete(comment)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "something went wrong",
			"message": err,
		})
		return
	}

	helper.SuccessResponse(ctx, http.StatusOK, gin.H{
		"message": "successfully deleted comment",
	})
}

// GetComment godoc
// @Summary Get details for a given id
// @Description Get details of comment corresponding is the input Id
// @Tags comment
// @Accept json
// @Produce json
// @Param id path int true "ID of the comment"
// @Security JWT
// @Success 200 {object} helper.SuccessResult{data=entity.Comment,code=int,message=string}
// @Failure 400 {object} helper.BadRequest{code=int,message=string}
// @Success 500 {object} helper.InternalServerError{code=int,message=string}
// @Router /comment/{id} [get]
// GetCommentHandler implements CommentHandler
func (h *commentHandler) GetCommentHandler(ctx *gin.Context) {
	requestParam := ctx.Param("id")
	commentId, _ := strconv.Atoi(requestParam)

	comment, err := h.commentService.GetById(uint(commentId))
	if err != nil {
		helper.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.SuccessResponse(ctx, http.StatusOK, comment)

}

// GetComment godoc
// @Summary Get all comment
// @Description Get details of all comment
// @Tags comment
// @Accept json
// @Produce json
// @Security JWT
// @Success 200 {object} helper.SuccessResult{data=entity.Comment,code=int,message=string}
// @Failure 400 {object} helper.BadRequest{code=int,message=string}
// @Success 500 {object} helper.InternalServerError{code=int,message=string}
// @Router /comment [get]
// GetCommentsHandler implements CommentHandler
func (h *commentHandler) GetCommentsHandler(ctx *gin.Context) {
	comments, err := h.commentService.GetAll()
	if err != nil {
		helper.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.SuccessResponse(ctx, http.StatusOK, comments)
}

// CreateComment godoc
// @Summary Create comment
// @Description Create new comment
// @Tags comment
// @Accept json
// @Produce json
// @Param comment body entity.CommentRequest true "create comment"
// @Security JWT
// @Success 201 {object} helper.SuccessResult{data=entity.Comment,code=int,message=string}
// @Failure 400 {object} helper.BadRequest{code=int,message=string}
// @Success 500 {object} helper.InternalServerError{code=int,message=string}
// @Router /comment [post]
// PostCommentHandler implements CommentHandler
func (h *commentHandler) PostCommentHandler(ctx *gin.Context) {
	var payload entity.CommentRequest

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		helper.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	comment, err := h.commentService.Create(payload, userID)
	if err != nil {
		helper.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.SuccessResponse(ctx, http.StatusCreated, comment)
}

// UpdateComment godoc
// @Summary Update comment identified by the given id
// @Description Update the comment corresponding to the input id
// @Tags comment
// @Accept json
// @Produce json
// @Param comment body entity.CommentRequest true "create comment"
// @Param id path int true "ID of the comment"
// @Security JWT
// @Success 200 {object} helper.SuccessResult{data=entity.Comment,code=int,message=string}
// @Failure 400 {object} helper.BadRequest{code=int,message=string}
// @Success 500 {object} helper.InternalServerError{code=int,message=string}
// @Router /comment/{id} [put]
// PutCommentHandler implements CommentHandler
func (h *commentHandler) PutCommentHandler(ctx *gin.Context) {
	requestParam := ctx.Param("id")
	commentId, _ := strconv.Atoi(requestParam)

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	var payload entity.CommentRequest
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		helper.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	comment, err := h.commentService.Update(payload, uint(commentId), userID)
	if err != nil {
		helper.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.SuccessResponse(ctx, http.StatusOK, comment)

}

func NewCommentHandler(commentService service.CommentService) CommentHandler {
	return &commentHandler{commentService: commentService}
}
