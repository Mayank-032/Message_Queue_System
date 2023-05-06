package routes

import (
	"go-message_queue_system/domain/interfaces/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var ProductUCase usecase.IProductUCase
func InitRoutes(apiGroup *gin.RouterGroup, productUCase usecase.IProductUCase) {
	ProductUCase = productUCase

	apiGroup.PUT("/product", upsertProduct)
}

func upsertProduct(c *gin.Context) {
	resData := gin.H{"status": false}
	request := CreateProductReq{}

	err := c.ShouldBindJSON(&request)
	if err == nil {
		err = request.validate()
	}
	if err != nil {
		log.Printf("Error: %v, invalid_request", err.Error())
		resData["message"] = "invalid request"
		c.JSON(http.StatusBadRequest, resData)
		return
	}

	product := request.toProductDto()
	err = ProductUCase.UpsertProduct(c, product)
	if err != nil {
		resData["message"] = "unable to create product"
		c.JSON(http.StatusOK, resData)
	}
	resData["status"] = true
	resData["message"] = "successfully upserted product"
	c.JSON(http.StatusOK, resData)
}