package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/utamadeveloper/golang-rest-api/config"
	"github.com/utamadeveloper/golang-rest-api/models"
)

func CreateOrder(c *gin.Context) {
	var (
		order models.Order
	)

	err := c.BindJSON(&order)
	if err != nil {
		ve := config.GetErrorValidator(err)
		if ve != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"statusCode": 400,
				"error":      ve,
				"message":    "Failed create order",
			})
			return
		}
	}

	config.Psql.Debug().Create(&order)
	c.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"data":       order,
		"message":    "Success create order",
	})
}

func FindAllOrder(c *gin.Context) {
	var (
		orders []models.Order
	)

	config.Psql.Debug().Find(&orders)
	c.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"data":       orders,
		"message":    "Success find all order",
	})
}

func FindOneOrder(c *gin.Context) {
	var (
		order models.Order
	)

	id := c.Param("id")
	config.Psql.Debug().Find(&order, id)

	if order.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": 400,
			"error":      "Order Not Found",
			"message":    "Failed find order",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"data":       order,
		"message":    "Success find order",
	})
}

func UpdateOrder(c *gin.Context) {
	var (
		order models.Order
	)

	id := c.Param("id")
	config.Psql.Debug().Find(&order, id)

	if order.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": 400,
			"error":      "Order Not Found",
			"message":    "Failed find order",
		})
		return
	}

	err := c.BindJSON(&order)
	if err != nil {
		ve := config.GetErrorValidator(err)
		if ve != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"statusCode": 400,
				"error":      ve,
				"message":    "Failed create order",
			})
			return
		}
	}

	config.Psql.Debug().Save(&order)

	c.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"data":       order,
		"message":    "Success update order",
	})
}

func DeleteOrder(c *gin.Context) {
	var (
		order models.Order
	)

	id := c.Param("id")
	config.Psql.Debug().Find(&order, id)

	if order.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": 400,
			"error":      "Order Not Found",
			"message":    "Failed find order",
		})
		return
	}

	config.Psql.Debug().Delete(&order)

	c.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"data":       order,
		"message":    "Success delete order",
	})
}

func FindAllOrderUser(c *gin.Context) {
	var (
		orders []models.Order
	)

	userId := c.Param("user_id")
	config.Psql.Debug().Where("user_id = ?", userId).Find(&orders)

	c.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"data":       orders,
		"message":    "Success find all order user",
	})
}
