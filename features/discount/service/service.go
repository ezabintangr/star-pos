package service

import (
	"errors"
	discountModel "star-pos/features/discount/model"
	"star-pos/features/discount/repository"
)

func Add(input discountModel.Discount) (string, error) {
	if input.UserID == "" {
		return "", errors.New("user id is required")
	}
	if input.OutletID == "" {
		return "", errors.New("outlet id is required")
	}
	if input.ProductID == "" {
		return "", errors.New("product id is required")
	}

	idNew, err := repository.Create(input)
	if err != nil {
		return "", err
	}

	return idNew, nil
}

func GetAllDiscounts() ([]discountModel.Discount, error) {
	allDiscounts, err := repository.ReadAllDiscounts()
	if err != nil {
		return nil, err
	}

	return allDiscounts, nil
}

func GetDiscount(id string) (*discountModel.Discount, error) {
	if id == "" {
		return nil, errors.New("you must login first")
	}

	result, err := repository.ReadDiscount(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func UpdateCurrentDiscount(input discountModel.Discount) error {
	if input.ID == "" {
		return errors.New("you must login first")
	}

	if input.UserID == "" && input.OutletID == "" && input.ProductID == "" {
		return errors.New("user_id, outlet_id, product_id field is required")
	}

	err := repository.UpdateDiscount(input)
	if err != nil {
		return err
	}

	return nil
}

func DeleteDiscount(id string) error {
	if id == "" {
		return errors.New("you must login first")
	}

	result, err := GetDiscount(id)
	if err != nil {
		return err
	}

	err = repository.DeleteDiscount(result.ID)
	if err != nil {
		return err
	}

	return nil
}
