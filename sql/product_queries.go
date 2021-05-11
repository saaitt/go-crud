package sql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	model "github.com/saaitt/go-crud/model"
)

type ProductRepo struct {
	DB *gorm.DB
}

func (p ProductRepo) Create(product *model.Product) error {
	if err := p.DB.Create(&product).Error; err != nil {
		return err
	}
	return nil
}

func (p ProductRepo) List(pageNo int) ([]model.Product, error) {
	products := []model.Product{}
	if err := p.DB.Limit(10).Offset(pageNo * 5).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
func (p ProductRepo) Disable(productID int) error {
	if err := p.DB.Model(&model.Product{}).Where("id = ?",productID).Update("is_active","false").Error;err !=nil{
		fmt.Println(err)
		return err
	}
	return nil
}
func (p ProductRepo) ListActive(pageNo int) ([]model.Product, error) {
	products := []model.Product{}
	if err := p.DB.Limit(10).Offset(pageNo * 5).Where("is_active= true").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
