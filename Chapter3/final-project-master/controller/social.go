package controller

import (
	"DTS-Kominfo-Hactiv8/Chapter3/final-project-master/helper"
	"DTS-Kominfo-Hactiv8/Chapter3/final-project-master/dto"
	"DTS-Kominfo-Hactiv8/Chapter3/final-project-master/entity"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SocialController struct {
	db *gorm.DB
}

func NewSocialController(db *gorm.DB) *SocialController {
	return &SocialController{
		db: db,
	}
}

func (controller *SocialController) CreateSocial(c *gin.Context) {
	userId, _ := c.Get("id")
	socialRequest := dto.SocialRequest{}

	err := c.ShouldBindJSON(&socialRequest)
	if err != nil {
		helper.BadRequestResponse(c, err.Error())
		return
	}

	social := entity.Social{
		Name:           socialRequest.Name,
		SocialMediaUrl: socialRequest.SocialMediaUrl,
		UserId:         uint(userId.(float64)),
	}

	_, err = govalidator.ValidateStruct(&social)
	if err != nil {
		helper.BadRequestResponse(c, err.Error())
		return
	}

	err = controller.db.Create(&social).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helper.NotFoundResponse(c, err.Error())
			return
		}
		helper.InternalServerJsonResponse(c, err.Error())
		return
	}

	helper.WriteJsonResponse(c, http.StatusOK, gin.H{
		"message": "Your social media has been successfully created",
	})
}

func (controller *SocialController) FindAllSocial(c *gin.Context) {
	var socials []entity.Social

	err := controller.db.Preload("User").Find(&socials).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helper.NotFoundResponse(c, err.Error())
			return
		}
		helper.InternalServerJsonResponse(c, err.Error())
		return
	}

	var response dto.SocialGetResponse
	for _, social := range socials {
		var userData dto.UserSocialResponse
		if social.User != nil {
			userData = dto.UserSocialResponse{
				Id:       social.User.Id,
				Username: social.User.Username,
			}
		}
		socialMediasResponse := dto.SocialData{
			Id:             social.Id,
			Name:           social.Name,
			SocialMediaUrl: social.SocialMediaUrl,
			CreatedAt:      social.CreatedAt,
			UpdatedAt:      social.UpdatedAt,
			User:           userData,
		}
		response.Socials = append(response.Socials, socialMediasResponse)

	}
	helper.WriteJsonResponse(c, http.StatusOK, response)
}

func (controller *SocialController) UpdateSocial(c *gin.Context) {
	userId, _ := c.Get("id")
	socialMediaId := c.Param("socialMediaId")
	var socialRequest dto.SocialRequest
	var social entity.Social

	err := c.ShouldBindJSON(&socialRequest)
	if err != nil {
		helper.BadRequestResponse(c, err.Error())
		return
	}

	updatedSocial := entity.Social{
		Name:           socialRequest.Name,
		SocialMediaUrl: socialRequest.SocialMediaUrl,
		UserId:         uint(userId.(float64)),
	}

	err = controller.db.First(&social, socialMediaId).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helper.NotFoundResponse(c, "data not found")
			return
		}
		helper.InternalServerJsonResponse(c, err.Error())
		return
	}

	if social.UserId != uint(userId.(float64)) {
		helper.WriteJsonResponse(c, http.StatusUnauthorized, "you're not allowed to update or edit this social media")
		return
	}

	err = controller.db.Model(&social).Updates(updatedSocial).Error
	if err != nil {
		helper.BadRequestResponse(c, err.Error())
		return
	}
	helper.WriteJsonResponse(c, http.StatusOK, gin.H{
		"message": "Your social media has been successfully updated",
	})
}

func (controller *SocialController) DeleteSocial(c *gin.Context) {
	userId, _ := c.Get("id")
	socialMediaId := c.Param("socialMediaId")
	var social entity.Social

	err := controller.db.First(&social, socialMediaId).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helper.NotFoundResponse(c, "data not found")
			return
		}
		helper.InternalServerJsonResponse(c, err.Error())
		return
	}

	if social.UserId != uint(userId.(float64)) {
		helper.WriteJsonResponse(c, http.StatusUnauthorized, gin.H{
			"error":   true,
			"message": "you're not allowed to delete this social media",
		})
		return
	}

	err = controller.db.Delete(&social).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helper.NotFoundResponse(c, err.Error())
			return
		}
		helper.InternalServerJsonResponse(c, err.Error())
		return
	}

	helper.WriteJsonResponse(c, http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}
