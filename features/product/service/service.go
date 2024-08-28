package service

import (
	"errors"
	"fmt"
	productModel "star-pos/features/product/model"
	productData "star-pos/features/product/repository"

	validator "github.com/go-playground/validator/v10"
)

func AddProduct(input productModel.Product) (string, error) {
	validate := validator.New()

	err := validate.Struct(input)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return "", fmt.Errorf("%s is %s", err.Field(), err.Tag())
		}
	}

	result, err := productData.CreateProduct(input)
	if err != nil {
		return "", err
	}

	return result, nil
}

func GetAllProducts() ([]productModel.Product, error) {
	result, err := productData.GetAllProducts()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetProduct(id string) (*productModel.Product, error) {
	if id == "" {
		return nil, errors.New("you must login first")
	}

	result, err := productData.GetProduct(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func UpdateProduct(input productModel.Product) error {
	if input.ID == "" {
		return errors.New("you must login first")
	}

	if input.CategoriesID == "" && input.ProductName == "" && input.Stock == 0 && input.Price == 0 {
		return errors.New("you didn't change anything")
	}

	err := productData.UpdateProduct(input)
	if err != nil {
		return err
	}

	return nil
}

func DeleteProduct(id string) error {
	if id == "" {
		return errors.New("you must login first")
	}

	err := productData.DeleteProduct(id)
	if err != nil {
		return err
	}

	return nil
}
