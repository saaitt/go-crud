package service

import (
	"github.com/saaitt/go-crud/model"
	"github.com/saaitt/go-crud/request"
	"github.com/saaitt/go-crud/response"
)

type ProductRepo interface {
	Create(product *model.Product) error
	List(pageNo int) ([]model.Product, error)
}
type ProductService struct {
	Repo ProductRepo
}

func (p ProductService) Create(request request.CreateProductRequest) (*response.ProductResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	product := model.Product{
		Title:       request.Title,
		Description: request.Description,
		IsActive:    true,
	}
	if err := p.Repo.Create(&product); err != nil {
		return nil, err
	}
	return &response.ProductResponse{
		ID:          int(product.ID),
		Description: product.Description,
		Title:       product.Title,
		IsActive:    product.IsActive,
	}, nil
}

func (p ProductService) List(pageNo int) ([]response.ProductResponse, error) {
	products, err := p.Repo.List(pageNo)
	if err != nil {
		return nil, err
	}
	responses := []response.ProductResponse{}
	for _, product := range products {
		responses = append(responses, response.ProductResponse{
			ID:          int(product.ID),
			Title:       product.Title,
			Description: product.Description,
			IsActive:    product.IsActive,
		})
	}
	return responses, nil
}
