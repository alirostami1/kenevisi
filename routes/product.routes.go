package routes

import (
	"github.com/carboncody/go-bootstrapper/controllers"
	"github.com/gin-gonic/gin"
)

type ProductRouteController struct {
	productController controllers.ProductController
}

func NewRouteProductController(productController controllers.ProductController) ProductRouteController {
	return ProductRouteController{productController}
}

func (uc *ProductRouteController) ProductRoute(rg *gin.RouterGroup) {

	router := rg.Group("product")
	router.POST("", uc.productController.CreateProduct)
	router.GET("", uc.productController.GetProducts)
}
