package controllers

import (
	"DTS-Kominfo-Hactiv8/Chapter3/Challange2/database"
	"DTS-Kominfo-Hactiv8/Chapter3/Challange2/enums"
	"DTS-Kominfo-Hactiv8/Chapter3/Challange2/helpers"
	"DTS-Kominfo-Hactiv8/Chapter3/Challange2/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType

	User := models.Users{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	role := c.PostForm("role")
    if role != "" {
        User.Role = enums.RoleUser(role)
    }

	err := db.Debug().Create(&User).Error

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": User.ID,
		"email" : User.Email,
		"full_name" : User.FullName,
		"role" : User.Role,
	})
}

func UserLogin(c *gin.Context){
	db:= database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType

	User := models.Users{}
	password := ""

	if contentType ==  appJSON{
		c.ShouldBindJSON(&User)
	} else{
		c.ShouldBind(&User)
	}

	password = User.Password
	err := db.Debug().Where("email = ? ", User.Email).Take(&User).Error

	if err !=nil{
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	comparePass:= helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass{
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email, User.Role)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}