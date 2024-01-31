package controllers

import (
	"net/http"

	"github.com/carboncody/go-bootstrapper/initializers"
	"github.com/carboncody/go-bootstrapper/models"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductController struct {
	DB *gorm.DB
}

func NewProductController(DB *gorm.DB) ProductController {
	return ProductController{DB}
}

func (uc *ProductController) CreateProduct(ctx *gin.Context) {
	var payload models.CreateProductPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	product := payload.ToProduct()

	if err := initializers.DB.Create(&product).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "message": "product created successfully"})
}

func (uc *ProductController) GetProducts(ctx *gin.Context) {
	products := []models.Product{}
	if err := initializers.DB.Model(models.Product{}).Preload(clause.Associations).Find(&products).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	response := lo.Map(products, func(p models.Product, _ int) models.GetProductResponse { return p.ToGetProductResponse() })

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "list of products", "data": response})
}
