package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_learn/models"
	"net/http"
)

func CheckLogin(c *gin.Context) {
	requestInfo := models.SearchPostReq{}
	err := c.BindJSON(&requestInfo)
	if err != nil {
		fmt.Println("Parse body failed. Error:", err)
		c.JSON(http.StatusOK, gin.H{
			"status": 1,
			"msg":    "params 格式错误",
		})
		return
	}
	//user := models.GetUser(requestInfo.UserName, requestInfo.UserPasswd)
	if requestInfo.UserName == "sjc" && requestInfo.UserPasswd == "123" {
		c.JSON(http.StatusOK, gin.H{
			"status": 2,
			"msg": "login success!",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": 3,
			"msg": "login failed...",
		})
	}
}