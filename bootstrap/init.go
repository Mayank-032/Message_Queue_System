package bootstrap

import (
	"database/sql"
	"go-message_queue_system/domain/interfaces/repository"
	"go-message_queue_system/domain/interfaces/usecase"
	_productRepo "go-message_queue_system/pkg/catalog/repository"
	_productUCase "go-message_queue_system/pkg/catalog/usecase"
	_productRoutes "go-message_queue_system/pkg/catalog/routes"

	"github.com/gin-gonic/gin"
)

var (
	productRepo repository.IProductRepo
	productUCase usecase.IProductUCase
)

func initRepos() {
	productRepo = _productRepo.NewProductRepository(&sql.DB{})
}

func initUCase() {
	productUCase = _productUCase.NewProductUCase(productRepo)
}

func initAPIs(apiGroup *gin.RouterGroup) {
	_productRoutes.InitRoutes(apiGroup, productUCase)
}

func Init(apiGroup *gin.RouterGroup) {
	initRepos()
	initUCase()
	initAPIs(apiGroup)
}