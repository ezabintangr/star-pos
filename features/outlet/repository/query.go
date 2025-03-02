package repository

import (
	"star-pos/app/databases"
	outletModel "star-pos/features/outlet/model"

	"github.com/google/uuid"
)

func Create(input outletModel.Outlet) (string, error) {
	input.ID = uuid.NewString()
	tx := databases.DB.Create(&input)
	if tx.Error != nil {
		return "", tx.Error
	}

	return input.ID, nil
}

func ReadAllOutlets() ([]outletModel.Outlet, error) {
	var allOutlets []outletModel.Outlet
	tx := databases.DB.Model(&outletModel.Outlet{}).Find(&allOutlets)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return allOutlets, nil
}

func ReadOutlet(id string) (*outletModel.Outlet, error) {
	var outlet outletModel.Outlet
	tx := databases.DB.Where("id = ?", id).First(&outlet)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &outlet, nil
}

func UpdateOutlet(input outletModel.Outlet) error {
	tx := databases.DB.Where("id = ?", input.ID).Updates(&input)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func DeleteOutlet(id string) error {
	tx := databases.DB.Delete(&outletModel.Outlet{}, "id = ?", id)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
