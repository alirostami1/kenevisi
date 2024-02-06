package models

import (
	"time"

	"github.com/samber/lo"
	"gorm.io/gorm"
)

type Product struct {
	ID            int            `gorm:"primaryKey" json:"id"`
	Name          string         `gorm:"not null" json:"name"`
	Type          string         `gorm:"not null" json:"type"`
	Price         int            `gorm:"not null" json:"price"`
	Quantity      int            `gorm:"not null" json:"quantity"`
	Description   string         `json:"description"`
	ImageFileName string         `json:"imageFileName"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `json:"deletedAt"`
	Colors        []Color        `json:"colors"`
	Sizes         []Size         `json:"sizes"`
	Filters       []Filter       `json:"filters"`
}

type Color struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"not null" json:"name"`
	ProductId int    `gorm:"not null" json:"productId"`
}

type Size struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"not null" json:"name"`
	ProductId int    `gorm:"not null" json:"productId"`
}

type Filter struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"not null" json:"name"`
	ProductId int    `gorm:"not null" json:"productId"`
}

type CreateProductPayload struct {
	Name          string   `json:"name" binding:"required"`
	Type          string   `json:"type" binding:"required"`
	Price         int      `json:"price" binding:"required"`
	Quantity      int      `json:"quantity"`
	Description   string   `json:"description"`
	ImageFileName string   `json:"imageFileName"`
	Colors        []string `json:"colors"`
	Sizes         []string `json:"sizes"`
	Filters       []string `json:"filters"`
}

func (c *CreateProductPayload) ToProduct() Product {
	return Product{
		Name:          c.Name,
		Type:          c.Type,
		Price:         c.Price,
		Quantity:      c.Quantity,
		Description:   c.Description,
		ImageFileName: c.ImageFileName,
		Colors:        lo.Map(c.Colors, func(color string, _ int) Color { return Color{Name: color} }),
		Sizes:         lo.Map(c.Sizes, func(size string, _ int) Size { return Size{Name: size} }),
		Filters:       lo.Map(c.Filters, func(filter string, _ int) Filter { return Filter{Name: filter} }),
	}
}

type GetProductResponse struct {
	ID            int      `json:"id"`
	Name          string   `json:"name"`
	Type          string   `json:"type"`
	Price         int      `json:"price"`
	Quantity      int      `json:"quantity"`
	Description   string   `json:"description"`
	ImageFileName string   `json:"imageFileName"`
	Colors        []string `json:"colors"`
	Sizes         []string `json:"sizes"`
	Filters       []string `json:"filters"`
}

func (p *Product) ToGetProductResponse() GetProductResponse {
	return GetProductResponse{
		ID:            p.ID,
		Name:          p.Name,
		Type:          p.Type,
		Price:         p.Price,
		Quantity:      p.Quantity,
		Description:   p.Description,
		ImageFileName: p.ImageFileName,
		Colors:        lo.Map(p.Colors, func(color Color, _ int) string { return color.Name }),
		Sizes:         lo.Map(p.Sizes, func(size Size, _ int) string { return size.Name }),
		Filters:       lo.Map(p.Filters, func(filter Filter, _ int) string { return filter.Name }),
	}
}

type ProductIdUri struct {
	ID int `uri:"id" binding:"required,gt=0"`
}
