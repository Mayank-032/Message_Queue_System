package bootstrap

import (
	"go-message_queue_system/db"
	"go-message_queue_system/domain/interfaces/repository"
	"go-message_queue_system/domain/interfaces/usecase"
	_productRepo "go-message_queue_system/pkg/catalog/repository"
	_productRoutes "go-message_queue_system/pkg/catalog/routes"
	_productUCase "go-message_queue_system/pkg/catalog/usecase"
	_userRepo "go-message_queue_system/pkg/user/repository"
	_userUCase "go-message_queue_system/pkg/user/usecase"

	"github.com/gin-gonic/gin"
)

var (
	userRepo     repository.IUserRepo
	userUCase    usecase.IUserUCase
	productRepo  repository.IProductRepo
	productUCase usecase.IProductUCase
)

func initRepos() {
	userRepo = _userRepo.NewUserRepository(db.Client)
	productRepo = _productRepo.NewProductRepository(db.Client)
}

func initUCase() {
	userUCase = _userUCase.NewUserUCase(userRepo)
	productUCase = _productUCase.NewProductUCase(userUCase, productRepo)
}

func initAPIs(apiGroup *gin.RouterGroup) {
	_productRoutes.InitRoutes(apiGroup, productUCase)
}

func Init(apiGroup *gin.RouterGroup) {
	initRepos()
	initUCase()
	initAPIs(apiGroup)
}
