package main

import (
	"github.com/cf-neto/nettrace/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.Setup(router)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}