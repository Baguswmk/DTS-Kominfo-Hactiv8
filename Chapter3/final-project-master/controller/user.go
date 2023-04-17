package controller

import (
	"DTS-Kominfo-Hactiv8/Chapter3/final-project-master/auth"
	"DTS-Kominfo-Hactiv8/Chapter3/final-project-master/helper"
	"DTS-Kominfo-Hactiv8/Chapter3/final-project-master/dto"
	"DTS-Kominfo-Hactiv8/Chapter3/final-project-master/entity"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	db *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		db: db,
	}
}

func (controller *UserController) CreateUser(c *gin.Context) {
	user := entity.User{}

	err := c.ShouldBindJSON(&user)
	if err != nil {
		helper.BadRequestResponse(c, err.Error())
		return
	}

	_, err = govalidator.ValidateStruct(&user)
	if err != nil {
		helper.BadRequestResponse(c, err.Error())
		return
	}

	err = controller.db.Create(&user).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helper.NotFoundResponse(c, err.Error())
			return
		}
		helper.InternalServerJsonResponse(c, err.Error())
		return
	}

	helper.WriteJsonResponse(c, http.StatusOK, gin.H{
		"message": "Your account has been successfully created",
	})
}

func (controller *UserController) UserLogin(c *gin.Context) {
	user := entity.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		helper.BadRequestResponse(c, err.Error())
		return
	}

	password := user.Password
	err = controller.db.Debug().Where("email = ?", user.Email).Take(&user).Error

	if err != nil {
		helper.WriteJsonResponse(c, http.StatusUnauthorized, gin.H{
			"error":   true,
			"message": "username / password is not match",
		})
		return
	}

	comparePass := auth.ComparePassword(user.Password, password)
	if !comparePass {
		helper.WriteJsonResponse(c, http.StatusUnauthorized, gin.H{
			"error":   true,
			"message": "username / password is not match",
		})
		return
	}
	token := auth.GenerateToken(user.Id, user.Email)
	c.JSON(http.StatusOK, dto.UserLoginResponse{
		Token: token,
	})
}

func (controller *UserController) UpdateUser(c *gin.Context) {
	userId, _ := c.Get("id")
	userReq := dto.UserUpdateRequest{}
	user := entity.User{}

	err := c.ShouldBindJSON(&userReq)
	if err != nil {
		helper.BadRequestResponse(c, err.Error())
		return
	}

	updatedUser := entity.User{
		Email:    userReq.Email,
		Username: userReq.Username,
	}

	_, err = govalidator.ValidateStruct(&userReq)
	if err != nil {
		helper.BadRequestResponse(c, err.Error())
		return
	}

	err = controller.db.First(&user, userId).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helper.NotFoundResponse(c, "User data not found")
			return
		}
		helper.InternalServerJsonResponse(c, err.Error())
		return
	}

	err = controller.db.Model(&user).Updates(updatedUser).Error
	if err != nil {
		helper.BadRequestResponse(c, err.Error())
		return
	}

	helper.WriteJsonResponse(c, http.StatusOK, gin.H{
		"message": "Your account has been successfully updated",
	})
}

func (controller *UserController) DeleteUser(c *gin.Context) {
	userId, _ := c.Get("id")
	var user entity.User

	err := controller.db.First(&user, userId).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helper.NotFoundResponse(c, "User not found")
			return
		}
		helper.BadRequestResponse(c, err.Error())
		return
	}

	err = controller.db.Delete(&user).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helper.NotFoundResponse(c, err.Error())
			return
		}
		helper.InternalServerJsonResponse(c, err.Error())
		return
	}

	helper.WriteJsonResponse(c, http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}
