package useCases

import (
	"log"

	"github.com/vitormoschetta/go/internal/application/requests"
	"github.com/vitormoschetta/go/internal/domain/category"
	"github.com/vitormoschetta/go/internal/domain/general"
	"github.com/vitormoschetta/go/internal/domain/models"
	"github.com/vitormoschetta/go/internal/domain/product"
)

type ProductUseCase struct {
	ProductRepository  product.IProductRepository
	CategoryRepository general.IRepository[category.Category]
}

func NewProductUseCase(pR product.IProductRepository, cR general.IRepository[category.Category]) *ProductUseCase {
	return &ProductUseCase{ProductRepository: pR, CategoryRepository: cR}
}

func (u *ProductUseCase) Save(request requests.CreateProductRequest) (response models.Response, statusCode int) {
	response = request.Validate()
	if len(response.Errors) > 0 {
		return response, 400
	}
	category, err := u.CategoryRepository.FindByID(request.CategoryId)
	if err != nil {
		log.Println("Error on find category: ", err)
		response.Errors = append(response.Errors, err.Error())
		return response, 500
	}
	if category.ID == "" {
		response.Errors = append(response.Errors, "Category not found")
		return response, 404
	}
	product := request.ToProductModel(category)
	response.Data = product
	err = u.ProductRepository.Save(product)
	if err != nil {
		log.Println("Error on save product: ", err)
		response.Errors = append(response.Errors, err.Error())
		return response, 500
	}
	return response, 201
}

func (u *ProductUseCase) Update(request requests.UpdateProductRequest) (response models.Response, statusCode int) {
	response = request.Validate()
	if len(response.Errors) > 0 {
		return response, 400
	}
	product, err := u.ProductRepository.FindByID(request.ID)
	if err != nil {
		log.Println("Error on find product: ", err)
		response.Errors = append(response.Errors, err.Error())
		return response, 500
	}
	if product.ID == "" {
		response.Errors = append(response.Errors, "Product not found")
		return response, 404
	}
	category, err := u.CategoryRepository.FindByID(request.CategoryId)
	if err != nil {
		log.Println("Error on find category: ", err)
		response.Errors = append(response.Errors, err.Error())
		return response, 500
	}
	if category.ID == "" {
		response.Errors = append(response.Errors, "Category not found")
		return response, 404
	}
	product.Update(request.Name, request.Price, category)
	response.Data = product
	err = u.ProductRepository.Update(product)
	if err != nil {
		log.Println("Error on update product: ", err)
		response.Errors = append(response.Errors, err.Error())
		return response, 500
	}
	return response, 200
}

func (u *ProductUseCase) Delete(id string) (response models.Response, statusCode int) {
	product, err := u.ProductRepository.FindByID(id)
	if err != nil {
		log.Println("Error on find product: ", err)
		response.Errors = append(response.Errors, err.Error())
		return response, 500
	}
	if product.ID == "" {
		response.Errors = append(response.Errors, "Product not found")
		return response, 404
	}
	response.Data = product
	err = u.ProductRepository.Delete(product.ID)
	if err != nil {
		log.Println("Error on delete product: ", err)
		response.Errors = append(response.Errors, err.Error())
		return response, 500
	}
	return response, 200
}

func (u *ProductUseCase) ApplyPromotion(request requests.ApplyPromotionProductRequest) (response models.Response, statusCode int) {
	response = request.Validate()
	if len(response.Errors) > 0 {
		return response, 400
	}
	product, err := u.ProductRepository.FindByID(request.ProductId)
	if err != nil {
		log.Println("Error on find product: ", err)
		response.Errors = append(response.Errors, err.Error())
		return response, 500
	}
	if product.ID == "" {
		response.Errors = append(response.Errors, "Product not found")
		return response, 404
	}
	product.ApplyPromotion(request.Percentage)
	response.Data = product
	err = u.ProductRepository.Update(product)
	if err != nil {
		log.Println("Error on apply promotion on product: ", err)
		response.Errors = append(response.Errors, err.Error())
		return response, 500
	}
	return response, 200
}

func (u *ProductUseCase) ApplyPromotionOnProductsByCategory(request requests.ApplyPromotionProductByCategoryRequest) (response models.Response, statusCode int) {
	response = request.Validate()
	if len(response.Errors) > 0 {
		return response, 400
	}
	category, err := u.CategoryRepository.FindByID(request.CategoryId)
	if err != nil {
		log.Println("Error on find category: ", err)
		response.Errors = append(response.Errors, err.Error())
		return response, 500
	}
	if category.ID == "" {
		response.Errors = append(response.Errors, "Category not found")
		return response, 404
	}
	err = u.ProductRepository.ApplyPromotionOnProductsByCategory(request.CategoryId, request.Percentage)
	if err != nil {
		log.Println("Error on apply promotion on products: ", err)
		response.Errors = append(response.Errors, err.Error())
		return response, 500
	}
	return response, 200
}
