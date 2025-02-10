package main

import (
	"blogging-platform-api/src/database"
	"blogging-platform-api/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()
	router := gin.Default()

	routes.BlogRoutes(router)

	router.Run(":3001")
}
