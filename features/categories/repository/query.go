package repository

import (
	"star-pos/app/databases"
	categoriesModel "star-pos/features/categories/model"

	"github.com/google/uuid"
)

func Create(input categoriesModel.Categories) (string, error) {
	input.ID = uuid.New().String()
	tx := databases.DB.Create(&input)
	if tx.Error != nil {
		return "", tx.Error
	}

	return input.ID, nil
}

func ReadAllCategories() ([]categoriesModel.Categories, error) {
	var allCategories []categoriesModel.Categories
	tx := databases.DB.Model(&categoriesModel.Categories{}).Find(&allCategories)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return allCategories, nil
}

func ReadCategory(id string) (*categoriesModel.Categories, error) {
	var currentCategory categoriesModel.Categories
	tx := databases.DB.Where("id = ?", id).First(&currentCategory)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &currentCategory, nil
}

func UpdateCategory(input categoriesModel.Categories) error {
	tx := databases.DB.Where("id = ?", input.ID).Updates(&input)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func DeleteCategory(id string) error {
	tx := databases.DB.Delete(&categoriesModel.Categories{}, "id = ?", id)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
