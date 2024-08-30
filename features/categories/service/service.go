package service

import (
	"errors"
	categoriesModel "star-pos/features/categories/model"
	"star-pos/features/categories/repository"
)

func AddCategories(input categoriesModel.Categories) (string, error) {
	if input.UserID == "" {
		return "", errors.New("user id is required")
	} else if input.CategoryName == "" {
		return "", errors.New("category name is required")
	}

	idNew, err := repository.Create(input)
	if err != nil {
		return "", err
	}

	return idNew, nil
}

func GetAllCategories() ([]categoriesModel.Categories, error) {
	allCategories, err := repository.ReadAllCategories()
	if err != nil {
		return nil, err
	}

	return allCategories, nil
}

func GetCategory(id string) (*categoriesModel.Categories, error) {
	if id == "" {
		return nil, errors.New("you must login first")
	}

	result, err := repository.ReadCategory(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func UpdateCurrentCategory(input categoriesModel.Categories) error {
	if input.ID == "" {
		return errors.New("you must login first")
	}

	if input.UserID == "" && input.CategoryName == "" {
		return errors.New("this field is required")
	}

	err := repository.UpdateCategory(input)
	if err != nil {
		return err
	}

	return nil
}

func DeleteCategory(id string) error {
	if id == "" {
		return errors.New("you must login first")
	}

	result, err := GetCategory(id)
	if err != nil {
		return err
	}

	err = repository.DeleteCategory(result.ID)
	if err != nil {
		return err
	}

	return nil
}
