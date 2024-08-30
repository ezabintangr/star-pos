package repository

import (
	"star-pos/app/databases"
	discountModel "star-pos/features/discount/model"

	"github.com/google/uuid"
)

func Create(input discountModel.Discount) (string, error) {
	input.ID = uuid.New().String()
	tx := databases.DB.Create(&input)
	if tx.Error != nil {
		return "", tx.Error
	}

	return input.ID, nil
}

func ReadAllDiscounts() ([]discountModel.Discount, error) {
	var allDiscounts []discountModel.Discount
	tx := databases.DB.Model(&discountModel.Discount{}).Find(&allDiscounts)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return allDiscounts, nil
}

func ReadDiscount(id string) (*discountModel.Discount, error) {
	var currentDiscount discountModel.Discount
	tx := databases.DB.Where("id = ?", id).First(&currentDiscount)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &currentDiscount, nil
}

func UpdateDiscount(input discountModel.Discount) error {
	tx := databases.DB.Where("id = ?", input.ID).Updates(&input)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func DeleteDiscount(id string) error {
	tx := databases.DB.Delete(&discountModel.Discount{}, "id = ?", id)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
