package bootsrap

import (
	"pustaka-api-Golang/config"
	"pustaka-api-Golang/database"
	"pustaka-api-Golang/routes"

	"github.com/gin-gonic/gin"
)

func BootsrapApp() {
	database.ConnectDatabase()
	app := gin.Default()

	routes.InitRoutes(app)

	

	app.Run(config.PORT)
}