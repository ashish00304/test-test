package main

import (
	"github.com/ashish00304/test-test/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.RegisterRoutes(router)

	router.Run(":8080")

}
