package mock

import (
	"context"

	"github.com/vitormoschetta/go/internal/domain/product"
	"github.com/vitormoschetta/go/pkg/pagination"
)

type ProductRepositoryFake struct {
	storage []product.Product
}

func NewProductRepositoryFake() *ProductRepositoryFake {
	return &ProductRepositoryFake{
		storage: []product.Product{},
	}
}

func (r *ProductRepositoryFake) Seed(items []product.Product) {
	r.storage = items
}

func (r *ProductRepositoryFake) FindAll(ctx context.Context, pagination *pagination.Pagination) (products []product.Product, err error) {
	pagination.Total = len(r.storage)
	pagination.BuildLastPage()
	for _, product := range r.storage {
		products = append(products, product)
	}
	return
}

func (r *ProductRepositoryFake) FindByID(ctx context.Context, id string) (product product.Product, err error) {
	for _, product := range r.storage {
		if product.ID == id {
			return product, nil
		}
	}
	return product, nil
}

func (r *ProductRepositoryFake) Save(ctx context.Context, p product.Product) error {
	r.storage = append(r.storage, p)
	return nil
}

func (r *ProductRepositoryFake) Update(ctx context.Context, p product.Product) error {
	for i, product := range r.storage {
		if product.ID == p.ID {
			r.storage[i] = p
		}
	}
	return nil
}

func (r *ProductRepositoryFake) Delete(ctx context.Context, id string) error {
	for i, product := range r.storage {
		if product.ID == id {
			r.storage = append(r.storage[:i], r.storage[i+1:]...)
		}
	}
	return nil
}

func (r *ProductRepositoryFake) FindByCategory(ctx context.Context, categoryID string) (products []product.Product, err error) {
	for _, product := range r.storage {
		if product.Category.ID == categoryID {
			products = append(products, product)
		}
	}
	return products, nil
}

func (r *ProductRepositoryFake) ApplyPromotionOnProductsByCategory(ctx context.Context, categoryId string, percentage float64) error {
	for i, product := range r.storage {
		if product.Category.ID == categoryId {
			r.storage[i].Price = product.Price - (product.Price * percentage)
		}
	}
	return nil
}
