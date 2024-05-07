package main

import (
	"udemy/restapi/db"
	"udemy/restapi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	servers := gin.Default()

	routes.RegisterRoutes(servers)

	servers.Run(":8000")
}
