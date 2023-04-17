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

type PhotoController struct {
	db *gorm.DB
}

func NewPhotoController(db *gorm.DB) *PhotoController {
	return &PhotoController{
		db: db,
	}
}

func (controller *PhotoController) CreatePhoto(c *gin.Context) {
	userId, _ := c.Get("id")
	photoRequest := dto.PhotoRequest{}

	err := c.ShouldBindJSON(&photoRequest)
	if err != nil {
		helper.BadRequestResponse(c, err.Error())
		return
	}

	photo := entity.Photo{
		Title:    photoRequest.Title,
		Caption:  photoRequest.Caption,
		PhotoUrl: photoRequest.PhotoUrl,
		UserId:   uint(userId.(float64)),
	}

	_, err = govalidator.ValidateStruct(&photo)
	if err != nil {
		helper.BadRequestResponse(c, err.Error())
		return
	}

	err = controller.db.Create(&photo).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helper.NotFoundResponse(c, err.Error())
			return
		}
		helper.InternalServerJsonResponse(c, err.Error())
		return
	}

	helper.WriteJsonResponse(c, http.StatusOK, gin.H{
		"message": "Your photo has been successfully created",
	})
}

func (controller *PhotoController) FindAllPhoto(c *gin.Context) {
	var photos []entity.Photo

	err := controller.db.Preload("User").Find(&photos).Error

	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helper.NotFoundResponse(c, err.Error())
			return
		}
		helper.InternalServerJsonResponse(c, err.Error())
		return
	}

	var response dto.PhotoGetResponse
	for _, photo := range photos {
		var userData dto.UserPhotoResponse
		if photo.User != nil {
			userData = dto.UserPhotoResponse{
				Username: photo.User.Username,
				Email:    photo.User.Email,
			}
		}
		photosResponse := dto.PhotoData{
			Id:        photo.Id,
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoUrl:  photo.PhotoUrl,
			CreatedAt: photo.CreatedAt,
			UpdatedAt: photo.UpdatedAt,
			User:      userData,
		}
		response.Photos = append(response.Photos, photosResponse)

	}
	helper.WriteJsonResponse(c, http.StatusOK, response)
}

func (controller *PhotoController) UpdatePhoto(c *gin.Context) {
	userId, _ := c.Get("id")
	photoMediaId := c.Param("photoId")
	var photoRequest dto.PhotoRequest
	var photo entity.Photo

	err := c.ShouldBindJSON(&photoRequest)
	if err != nil {
		helper.BadRequestResponse(c, err.Error())
		return
	}

	updatePhoto := entity.Photo{
		Title:    photoRequest.Title,
		Caption:  photoRequest.Caption,
		PhotoUrl: photoRequest.PhotoUrl,
	}

	err = controller.db.First(&photo, photoMediaId).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helper.NotFoundResponse(c, "data not found")
			return
		}
		helper.InternalServerJsonResponse(c, err.Error())
		return
	}

	if photo.UserId != uint(userId.(float64)) {
		helper.WriteJsonResponse(c, http.StatusUnauthorized, gin.H{
			"error":   true,
			"message": "you're not allowed to update or edit this photo",
		})
		return
	}

	err = controller.db.Model(&photo).Updates(updatePhoto).Error
	if err != nil {
		helper.BadRequestResponse(c, err.Error())
		return
	}

	

	helper.WriteJsonResponse(c, http.StatusOK, gin.H{
		"message": "Your photo has been successfully updated",
	})
}

func (controller *PhotoController) DeletePhoto(c *gin.Context) {
	userId, _ := c.Get("id")
	photoId := c.Param("photoId")
	var photo entity.Photo

	err := controller.db.First(&photo, photoId).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helper.NotFoundResponse(c, "data not found")
			return
		}
		helper.InternalServerJsonResponse(c, err.Error())
		return
	}

	if photo.UserId != uint(userId.(float64)) {
		helper.WriteJsonResponse(c, http.StatusUnauthorized, gin.H{
			"error":   true,
			"message": "you're not allowed to delete this photo",
		})
		return
	}

	err = controller.db.Delete(&photo).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helper.NotFoundResponse(c, err.Error())
			return
		}
		helper.InternalServerJsonResponse(c, err.Error())
		return
	}

	helper.WriteJsonResponse(c, http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
