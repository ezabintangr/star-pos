package service

import (
	"errors"
	outletModel "star-pos/features/outlet/model"
	"star-pos/features/outlet/repository"
)

func Add(input outletModel.Outlet) (string, error) {
	if input.UserID == "" {
		return "", errors.New("user id is required")
	}
	if input.OutletName == "" {
		return "", errors.New("outlet name is required")
	}
	if input.Address == "" {
		return "", errors.New("address is required")
	}
	if input.PhoneOutlet == "" {
		return "", errors.New("phone outlet id is required")
	}

	idNew, err := repository.Create(input)
	if err != nil {
		return "", err
	}

	return idNew, nil
}

func GetAll() ([]outletModel.Outlet, error) {
	allOutlets, err := repository.ReadAllOutlets()
	if err != nil {
		return nil, err
	}

	return allOutlets, nil
}

func GetById(id string) (*outletModel.Outlet, error) {
	if id == "" {
		return nil, errors.New("id must not be empty")
	}

	outlet, err := repository.ReadOutlet(id)
	if err != nil {
		return nil, err
	}

	return outlet, nil
}

func Update(input outletModel.Outlet) error {
	err := repository.UpdateOutlet(input)
	if err != nil {
		return err
	}

	return nil
}

func Delete(id string) error {
	if id == "" {
		return errors.New("id cannot be empty")
	}

	err := repository.DeleteOutlet(id)
	if err != nil {
		return err
	}

	return nil
}
