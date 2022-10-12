package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/utamadeveloper/golang-rest-api/config"
	"github.com/utamadeveloper/golang-rest-api/models"
)

func CreateUser(c *gin.Context) {
	var (
		user models.User
	)

	err := c.BindJSON(&user)
	if err != nil {
		ve := config.GetErrorValidator(err)
		if ve != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"statusCode": 400,
				"error":      ve,
				"message":    "Failed create user",
			})
			return
		}
	}

	config.Psql.Debug().Create(&user)
	c.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"data":       user,
		"message":    "Success create user",
	})
}

func FindAllUser(c *gin.Context) {
	var (
		users []models.User
	)

	config.Psql.Debug().Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"data":       users,
		"message":    "Success find all user",
	})
}

func FindOneUser(c *gin.Context) {
	var (
		user models.User
	)

	id := c.Param("id")
	config.Psql.Debug().Find(&user, id)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": 400,
			"error":      "User Not Found",
			"message":    "Failed find user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"data":       user,
		"message":    "Success find user",
	})
}

func UpdateUser(c *gin.Context) {
	var (
		user models.User
	)

	id := c.Param("id")
	config.Psql.Debug().Find(&user, id)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": 400,
			"error":      "User Not Found",
			"message":    "Failed find user",
		})
		return
	}

	err := c.BindJSON(&user)
	if err != nil {
		ve := config.GetErrorValidator(err)
		if ve != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"statusCode": 400,
				"error":      ve,
				"message":    "Failed create user",
			})
			return
		}
	}

	config.Psql.Debug().Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"data":       user,
		"message":    "Success update user",
	})
}

func DeleteUser(c *gin.Context) {
	var (
		user models.User
	)

	id := c.Param("id")
	config.Psql.Debug().Find(&user, id)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": 400,
			"error":      "User Not Found",
			"message":    "Failed find user",
		})
		return
	}

	config.Psql.Debug().Delete(&user)

	c.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"data":       user,
		"message":    "Success delete user",
	})
}
