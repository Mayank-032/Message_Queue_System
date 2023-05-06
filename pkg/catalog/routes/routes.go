package routes

import (
	"go-message_queue_system/domain/interfaces/usecase"

	"github.com/gin-gonic/gin"
)

var ProductUCase usecase.IProductUCase
func InitRoutes(apiGroup *gin.RouterGroup, productUCase usecase.IProductUCase) {
	ProductUCase = productUCase

	apiGroup.POST("/product", createProduct)
}

func createProduct(c *gin.Context) {
}