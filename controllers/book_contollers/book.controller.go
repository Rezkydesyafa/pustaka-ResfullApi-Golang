package book_contollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(c *gin.Context){
	
	c.JSON(http.StatusOK,gin.H{
		"Hello":"Book",
	})
}
func GetBooksByID(c *gin.Context){
	id := c.Param("id")
	c.JSON(http.StatusOK,gin.H{
		"Hello":"Book",
		"idBook":id,
	})
}