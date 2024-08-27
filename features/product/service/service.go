package service

import (
	"errors"
	"fmt"
	productModel "star-pos/features/product/model"
	"star-pos/features/product/repository"

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

	result, err := repository.CreateProduct(input)
	if err != nil {
		return "", err
	}

	return result, nil
}

func GetAllProducts() ([]productModel.Product, error) {
	result, err := repository.GetAllProducts()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetProduct(id string) (*productModel.Product, error) {
	if id == "" {
		return nil, errors.New("you must login first")
	}

	result, err := repository.GetProduct(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func UpdateProduct(input productModel.Product) error {
	if input.ID == "" {
		return errors.New("you must login first")
	}

	result, err := GetProduct(input.ID)
	if err != nil {
		return err
	}

	if input.CategoriesID == "" && input.ProductName == "" && input.Stock == result.Stock && input.Price == result.Price {
		return errors.New("you didn't change anything")
	}

	err = repository.UpdateProduct(input)
	if err != nil {
		return err
	}

	return nil
}

func DeleteProduct(id string) error {
	if id == "" {
		return errors.New("you must login first")
	}

	err := repository.DeleteProduct(id)
	if err != nil {
		return err
	}

	return nil
}
