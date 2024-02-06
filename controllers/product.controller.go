package controllers

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"net/http"
	"path/filepath"

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

func (uc *ProductController) UploadProductImage(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	config, err := initializers.LoadConfig()

	// Inside the UploadProductImage function
	randomHash, err := generateRandomHashString(16)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "failed to generate random hash"})
		return
	}

	filename := randomHash + filepath.Ext(filepath.Base(file.Filename))
	filePath := filepath.Clean(config.UploadFolder + "/" + filename)
	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "message": "product image uploaded successfully", "filename": filename})
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

func generateRandomHashString(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func (uc *ProductController) DeleteProduct(ctx *gin.Context) {
	// productIdInt := ctx.GetInt("id")
	productIdUri := models.ProductIdUri{}
	err := ctx.BindUri(&productIdUri)
	if err != nil {
		return
	}

	log.Println(productIdUri.ID)

	if err := initializers.DB.Delete(&models.Product{}, productIdUri.ID).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "product deleted successfully"})
}
