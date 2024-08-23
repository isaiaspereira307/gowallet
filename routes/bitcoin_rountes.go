package routes

import (
	"github.com/gin-gonic/gin"
	bitcoin_handlers "github.com/isaiaspereira307/gowallet/handlers/bitcoin"
)

func InitializeBitcoinRoutes(router *gin.RouterGroup) {
	router.GET("/bitcoin/:id", bitcoin_handlers.GetBitcoin)
	router.GET("/bitcoin-price", bitcoin_handlers.GetBitcoinPriceUSD)
	router.POST("/bitcoin", bitcoin_handlers.CreateBitcoin)
	router.PUT("/bitcoin/:id", bitcoin_handlers.UpdateBitcoin)
	router.DELETE("/bitcoin/:id", bitcoin_handlers.DeleteBitcoin)
}
