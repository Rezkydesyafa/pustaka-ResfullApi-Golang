package routes

import (
	"pustaka-api-Golang/controllers/book_contollers"
	"pustaka-api-Golang/controllers/user_controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(app *gin.Engine) {

	route := app
	route.GET("/users",user_controllers.GetAllUsers)
	route.GET("/",func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"status": 200,
   			"name": "Pustaka-Resfull-Api",
    		"version": "1.0.0",
    		"docs_link": "https://comingsoon.com/docs",
    		"health_check": 100,
    		"is_open": false,
		})
	})

	route.GET("/books",book_contollers.GetAllBooks)
	route.GET("/books/:id",book_contollers.GetBooksByID)
}