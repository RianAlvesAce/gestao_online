package server

import (
	"github.com/RianAlvesAce/gestao_online/internal/server/routes"
	"github.com/gin-gonic/gin"
)

func InitApi() {
	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")
}