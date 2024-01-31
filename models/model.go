package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type Product struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	Type        string         `gorm:"not null" json:"type"`
	Price       int            `gorm:"not null" json:"price"`
	Quantity    int            `gorm:"not null" json:"quantity"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt"`
	Colors      []Color        `gorm:"foreignKey:ProductId;references:ID" json:"colors"`
	Sizes       []Size         `gorm:"foreignKey:ProductId;references:ID" json:"sizes"`
	Filters     []Filter       `gorm:"foreignKey:ProductId;references:ID" json:"filters"`
}

type Color struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	ProductId uuid.UUID `gorm:"not null" json:"productId"`
}

type Size struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	ProductId uuid.UUID `gorm:"not null" json:"productId"`
}

type Filter struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	ProductId uuid.UUID `gorm:"not null" json:"productId"`
}

type CreateProductPayload struct {
	Name        string   `json:"name"`
	Type        string   `json:"type"`
	Price       int      `json:"price"`
	Quantity    int      `json:"quantity"`
	Description string   `json:"description"`
	Colors      []string `json:"colors"`
	Sizes       []string `json:"sizes"`
	Filters     []string `json:"filters"`
}

func (c *CreateProductPayload) ToProduct() Product {
	return Product{
		Name:        c.Name,
		Type:        c.Type,
		Price:       c.Price,
		Quantity:    c.Quantity,
		Description: c.Description,
		Colors:      lo.Map(c.Colors, func(color string, _ int) Color { return Color{Name: color} }),
		Sizes:       lo.Map(c.Sizes, func(size string, _ int) Size { return Size{Name: size} }),
		Filters:     lo.Map(c.Filters, func(filter string, _ int) Filter { return Filter{Name: filter} }),
	}
}

type GetProductResponse struct {
	Name        string   `json:"name"`
	Type        string   `json:"type"`
	Price       int      `json:"price"`
	Quantity    int      `json:"quantity"`
	Description string   `json:"description"`
	Colors      []string `json:"colors"`
	Sizes       []string `json:"sizes"`
	Filters     []string `json:"filters"`
}

func (p *Product) ToGetProductResponse() GetProductResponse {
	return GetProductResponse{
		Name:        p.Name,
		Type:        p.Type,
		Price:       p.Price,
		Quantity:    p.Quantity,
		Description: p.Description,
		Colors:      lo.Map(p.Colors, func(color Color, _ int) string { return color.Name }),
		Sizes:       lo.Map(p.Sizes, func(size Size, _ int) string { return size.Name }),
		Filters:     lo.Map(p.Filters, func(filter Filter, _ int) string { return filter.Name }),
	}
}
