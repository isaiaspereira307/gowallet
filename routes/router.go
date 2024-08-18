package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gowallet/configs"
	"github.com/isaiaspereira307/gowallet/internal/db"
)

var (
	logger *configs.Logger
)

func Initialize(queries *db.Queries) {

	r := gin.Default()
	InitializeRoutes(r, queries)
	port := fmt.Sprintf(":%s", configs.GetServerPort())

	if err := r.Run(port); err != nil {
		logger.Errorf("Failed to run server: %v", err)
	}
}
