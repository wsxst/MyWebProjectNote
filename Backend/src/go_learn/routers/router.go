package routers

import (
	"go_learn/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoadRouters(router *gin.Engine) {

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"data":   "ok",
		})
	})
	router.POST("/login", controllers.CheckLogin)
}
