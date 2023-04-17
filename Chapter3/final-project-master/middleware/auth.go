package middleware

import (
	"DTS-Kominfo-Hactiv8/Chapter3/final-project-master/auth"
	"DTS-Kominfo-Hactiv8/Chapter3/final-project-master/database"
	"DTS-Kominfo-Hactiv8/Chapter3/final-project-master/entity"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		headerToken := c.Request.Header.Get("Authorization")
		if headerToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   true,
				"message": "UNAUTHORIZED",
			})
			return
		}

		bearer := strings.HasPrefix(headerToken, "Bearer")
		if !bearer {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   true,
				"message": "UNAUTHORIZED",
			})
			return
		}

		bearerToken := strings.Split(headerToken, "Bearer ")[1]

		verify, err := auth.VerifyToken(bearerToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   true,
				"message": err.Error(),
			})
			return
		}
		data := verify.(jwt.MapClaims)

		c.Set("id", data["id"])
		c.Set("email", data["email"])
		c.Next()
	}
}

func SocialAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.ConnectDB()
		socialId, err := strconv.Atoi(c.Param("socialMediaId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid parameter",
			})
			return
		}
	
		userID, _ := c.Get("id")
		social := entity.Social{}

		err = db.Select("user_id").First(&social, uint(socialId)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "data doesn't exist",
			})
			return
		}

		if social.UserId != uint(userID.(float64)){
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to access this data",
			})
			return
		}
		c.Next()
	}
}

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.ConnectDB()
		photoId, err := strconv.Atoi(c.Param("photoId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid parameter",
			})
			return
		}
		userID, _ := c.Get("id")
		photo := entity.Photo{}

		err = db.Select("user_id").First(&photo, uint(photoId)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "data doesn't exist",
			})
			return
		}

		if photo.UserId != uint(userID.(float64)) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to access this data",
			})
			return
		}
		c.Next()
	}
}
func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.ConnectDB()
		commentId, err := strconv.Atoi(c.Param("commentId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid parameter",
			})
			return
		}
		userID, _ := c.Get("id")
		comment := entity.Comment{}

		err = db.Select("user_id").First(&comment, uint(commentId)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "data doesn't exist",
			})
			return
		}

		if comment.UserId != uint(userID.(float64)) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to access this data",
			})
			return
		}
		c.Next()
	}
}
