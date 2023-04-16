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

// GetCommentsHandler implements CommentHandler
func (h *commentHandler) GetCommentsHandler(ctx *gin.Context) {
	comments, err := h.commentService.GetAll()
	if err != nil {
		helper.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.SuccessResponse(ctx, http.StatusOK, comments)
}

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
