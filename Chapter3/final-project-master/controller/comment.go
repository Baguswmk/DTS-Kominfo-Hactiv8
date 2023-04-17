package controller

import (
	"DTS-Kominfo-Hactiv8/Chapter3/final-project-master/helper"
	"DTS-Kominfo-Hactiv8/Chapter3/final-project-master/dto"
	"DTS-Kominfo-Hactiv8/Chapter3/final-project-master/entity"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type CommentController struct {
	db *gorm.DB
}

func NewCommentController(db *gorm.DB) *CommentController {
	return &CommentController{
		db: db,
	}
}

func (controller *CommentController) CreateComment(c *gin.Context) {
	userId, _ := c.Get("id")
	commentRequest := dto.CommentRequest{}

	err := c.ShouldBindJSON(&commentRequest)
	if err != nil {
		helper.BadRequestResponse(c, err.Error())
		return
	}

	comment := entity.Comment{
		Message: commentRequest.Message,
		PhotoId: commentRequest.PhotoId,
		UserId:  uint(userId.(float64)),
	}

	_, err = govalidator.ValidateStruct(&comment)
	if err != nil {
		helper.BadRequestResponse(c, err.Error())
		return
	}

	err = controller.db.Create(&comment).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helper.NotFoundResponse(c, err.Error())
			return
		}
		helper.InternalServerJsonResponse(c, err.Error())
		return
	}

	helper.WriteJsonResponse(c, http.StatusOK, gin.H{
		"message": "Your comment has been successfully created",
	})
}

func (controller *CommentController) FindAllComment(c *gin.Context) {
	var comments []entity.Comment

	err := controller.db.Preload("User").Preload("Photo").Find(&comments).Error

	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helper.NotFoundResponse(c, err.Error())
			return
		}
		helper.InternalServerJsonResponse(c, err.Error())
		return
	}

	var response dto.CommentGetResponse
	for _, comment := range comments {
		var userData dto.UserCommentResponse
		if comment.User != nil {
			userData = dto.UserCommentResponse{
				Id:       comment.User.Id,
				Username: comment.User.Username,
				Email:    comment.User.Email,
			}
		}

		var photoData dto.PhotoCommentResponse
		if comment.Photo != nil {
			photoData = dto.PhotoCommentResponse{
				Id:       comment.Photo.Id,
				Title:    comment.Photo.Title,
				Caption:  comment.Photo.Caption,
				PhotoUrl: comment.Photo.PhotoUrl,
				UserId:   comment.Photo.UserId,
			}
		}
		commentsResponse := dto.CommentData{
			Id:        comment.Id,
			Message:   comment.Message,
			PhotoId:   comment.PhotoId,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
			User:      userData,
			Photos:    photoData,
		}
		response.Comments = append(response.Comments, commentsResponse)

	}
	helper.WriteJsonResponse(c, http.StatusOK, response)
}

func (controller *CommentController) UpdateComment(c *gin.Context) {
	userId, _ := c.Get("id")
	commentId := c.Param("commentId")
	var commentUpdateRequest dto.CommentUpdateRequest
	var comment entity.Comment

	err := c.ShouldBindJSON(&commentUpdateRequest)
	if err != nil {
		helper.BadRequestResponse(c, err.Error())
		return
	}

	updatePhoto := entity.Comment{
		Message: commentUpdateRequest.Message,
	}

	err = controller.db.First(&comment, commentId).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helper.NotFoundResponse(c, "data not found")
			return
		}
		helper.InternalServerJsonResponse(c, err.Error())
		return
	}

	if comment.UserId != uint(userId.(float64)) {
		helper.WriteJsonResponse(c, http.StatusUnauthorized, gin.H{
			"error":   true,
			"message": "you're not allowed to update or edit this comment",
		})
		return
	}

	err = controller.db.Model(&comment).Updates(updatePhoto).Error
	if err != nil {
		helper.BadRequestResponse(c, err.Error())
		return
	}

	helper.WriteJsonResponse(c, http.StatusOK, gin.H{
		"message": "Your comment has been successfully updated",
	})
}

func (controller *CommentController) DeleteComment(c *gin.Context) {
	userId, _ := c.Get("id")
	commentId := c.Param("commentId")
	var comment entity.Comment

	err := controller.db.First(&comment, commentId).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helper.NotFoundResponse(c, "data not found")
			return
		}
		helper.InternalServerJsonResponse(c, err.Error())
		return
	}

	if comment.UserId != uint(userId.(float64)) {
		helper.WriteJsonResponse(c, http.StatusUnauthorized, gin.H{
			"error":   true,
			"message": "you're not allowed to delete this comment",
		})
		return
	}

	err = controller.db.Delete(&comment).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helper.NotFoundResponse(c, err.Error())
			return
		}
		helper.InternalServerJsonResponse(c, err.Error())
		return
	}

	helper.WriteJsonResponse(c, http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
