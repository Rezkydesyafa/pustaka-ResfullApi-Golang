package user_controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context){
	
	c.JSON(http.StatusOK,gin.H{
		"Testing":"123",
	})
}