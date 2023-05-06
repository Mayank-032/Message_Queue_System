package bootstrap

import (
	"go-message_queue_system/db"
	"go-message_queue_system/domain/interfaces/repository"
	"go-message_queue_system/domain/interfaces/usecase"
	_productRepo "go-message_queue_system/pkg/catalog/repository"
	_productRoutes "go-message_queue_system/pkg/catalog/routes"
	_productUCase "go-message_queue_system/pkg/catalog/usecase"

	"github.com/gin-gonic/gin"
)

var (
	productRepo repository.IProductRepo
	productUCase usecase.IProductUCase
)

func initRepos() {
	productRepo = _productRepo.NewProductRepository(db.Client)
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