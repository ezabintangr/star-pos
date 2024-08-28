package repository

import (
	"star-pos/app/databases"
	productModel "star-pos/features/product/model"

	"github.com/google/uuid"
)

func CreateProduct(input productModel.Product) (string, error) {
	input.ID = uuid.New().String()
	tx := databases.DB.Create(&input)
	if tx.Error != nil {
		return "", tx.Error
	}

	return input.ID, nil
}

func GetAllProducts() ([]productModel.Product, error) {
	var allProducts []productModel.Product
	tx := databases.DB.Model(&productModel.Product{}).Find(&allProducts)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return allProducts, nil
}

func GetProduct(id string) (*productModel.Product, error) {
	var currentProduct productModel.Product
	tx := databases.DB.Where("id = ?", id).First(&currentProduct)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &currentProduct, nil
}

func UpdateProduct(input productModel.Product) error {
	tx := databases.DB.Where("id = ?", input.ID).Updates(&input)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func DeleteProduct(id string) error {
	tx := databases.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	err := tx.Exec("ALTER TABLE cart_products DROP FOREIGN KEY fk_cart_products_product")
	if err.Error != nil {
		tx.Rollback()
		return err.Error
	}

	err = tx.Exec(`ALTER TABLE cart_products
	ADD CONSTRAINT fk_cart_products_product
	FOREIGN KEY (product_id) REFERENCES products(id)
	ON DELETE CASCADE;`)
	if err.Error != nil {
		tx.Rollback()
		return err.Error
	}

	err = tx.Exec("ALTER TABLE discounts DROP FOREIGN KEY fk_discounts_product")
	if err.Error != nil {
		tx.Rollback()
		return err.Error
	}

	err = tx.Exec(`ALTER TABLE discounts
	ADD CONSTRAINT fk_discounts_product
	FOREIGN KEY (product_id) REFERENCES products(id)
	ON DELETE CASCADE;`)
	if err.Error != nil {
		tx.Rollback()
		return err.Error
	}

	err = databases.DB.Delete(&productModel.Product{}, "id = ?", id)
	if err.Error != nil {
		tx.Rollback()
		return err.Error
	}

	tx.Commit()

	return nil
}
