package migrations

import (
	"encoding/json"
	"log"
	"os"
	"star-pos/app/databases"
	cartModel "star-pos/features/cart/model"
	cartProductModel "star-pos/features/cart_product/model"
	categoriesModel "star-pos/features/categories/model"
	discountModel "star-pos/features/discount/model"
	outletModel "star-pos/features/outlet/model"
	productModel "star-pos/features/product/model"
	transactionModel "star-pos/features/transaction/model"
	transactionDetailModel "star-pos/features/transaction_detail/model"
	userModel "star-pos/features/user/model"
	"time"

	"gorm.io/gorm"
)

func InitMigration() {
	roll := databases.DB.Begin()

	file, err := os.ReadFile("dummy/userdummy.json")
	if err != nil {
		log.Fatal("error reading file: ", err)
	}
	var users []userModel.User
	err = json.Unmarshal(file, &users)
	if err != nil {
		log.Fatal("error unmarshalling JSON: ", err)
	}

	for _, data := range users {
		var existingUser userModel.User
		result := databases.DB.Where("id = ?", data.ID).First(&existingUser)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				tx := databases.DB.Create(&data)
				if tx.Error != nil {
					roll.Rollback()
					log.Fatal("error insert data: ", tx.Error)
				}
			}
		} else {
			log.Println("table already exist")
			break
		}
	}

	file, err = os.ReadFile("dummy/outletdummy.json")
	if err != nil {
		log.Fatal("error reading file: ", err)
	}
	var outlet []outletModel.Outlet
	err = json.Unmarshal(file, &outlet)
	if err != nil {
		log.Fatal("error unmarshalling JSON: ", err)
	}

	for i, data := range outlet {
		var existingOutlet outletModel.Outlet
		result := databases.DB.Where("id = ?", data.ID).First(&existingOutlet)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				data.UserID = users[i].ID
				tx := databases.DB.Create(&data)
				if tx.Error != nil {
					log.Fatal("error insert data: ", tx.Error)
				}
			}
		} else {
			log.Println("table already exist")
			break
		}
	}

	file, err = os.ReadFile("dummy/categoriesdummy.json")
	if err != nil {
		log.Fatal("error reading file: ", err)
	}
	var categories []categoriesModel.Categories
	err = json.Unmarshal(file, &categories)
	if err != nil {
		log.Fatal("error unmarshalling JSON: ", err)
	}

	for i, data := range categories {
		var existingCategories categoriesModel.Categories
		result := databases.DB.Where("id = ?", data.ID).First(&existingCategories)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				data.UserID = users[i].ID
				tx := databases.DB.Create(&data)
				if tx.Error != nil {
					roll.Rollback()
					log.Fatal("error insert data: ", tx.Error)
				}

			}
		} else {
			log.Println("table already exist")
			break
		}
	}

	file, err = os.ReadFile("dummy/productdummy.json")
	if err != nil {
		log.Fatal("error reading file: ", err)
	}
	var products []productModel.Product
	err = json.Unmarshal(file, &products)
	if err != nil {
		log.Fatal("error unmarshalling JSON: ", err)
	}

	for i, data := range products {
		var existingProduct productModel.Product
		result := databases.DB.Where("id = ?", data.ID).First(&existingProduct)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				data.UserID = users[i].ID
				data.CategoriesID = categories[i].ID
				tx := databases.DB.Create(&data)
				if tx.Error != nil {
					roll.Rollback()
					log.Fatal("error insert data: ", tx.Error)
				}
			}
		} else {
			log.Println("table already exist")
			break
		}
	}

	file, err = os.ReadFile("dummy/cartdummy.json")
	if err != nil {
		log.Fatal("error reading file: ", err)
	}

	var carts []cartModel.Cart
	err = json.Unmarshal(file, &carts)
	if err != nil {
		log.Fatal("error unmarshalling JSON: ", err)
	}

	for i, data := range carts {
		var existingCart cartModel.Cart
		_, err = time.Parse("2006-01-02", data.Date.Format("2006-01-02"))
		if err != nil {
			log.Fatal("error parse date: ", err)
		}

		result := databases.DB.Where("id = ?", data.ID).First(&existingCart)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				data.UserID = users[i].ID
				data.OutletID = outlet[i].ID
				tx := databases.DB.Create(&data)
				if tx.Error != nil {
					roll.Rollback()
					log.Fatal("error insert data: ", tx.Error)
				}
			}
		} else {
			log.Println("table already exist")
			break
		}
	}

	file, err = os.ReadFile("dummy/cartproductdummy.json")
	if err != nil {
		log.Fatal("error reading file: ", err)
	}
	var cartProduct []cartProductModel.CartProduct
	err = json.Unmarshal(file, &cartProduct)
	if err != nil {
		log.Fatal("error unmarshalling JSON: ", err)
	}

	for i, data := range cartProduct {
		var existingCartProduct cartProductModel.CartProduct
		result := databases.DB.Where("id = ?", data.ID).First(&existingCartProduct)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				data.CartID = carts[i].ID
				data.ProductID = products[i].ID
				tx := databases.DB.Create(&data)
				if tx.Error != nil {
					roll.Rollback()
					log.Fatal("error insert data: ", tx.Error)
				}
			}
		} else {
			log.Println("table already exist")
			break
		}
	}

	file, err = os.ReadFile("dummy/discountdummy.json")
	if err != nil {
		log.Fatal("error reading file: ", err)
	}
	var discounts []discountModel.Discount
	err = json.Unmarshal(file, &discounts)
	if err != nil {
		log.Fatal("error unmarshalling JSON: ", err)
	}

	for i, data := range discounts {
		var existingDiscount discountModel.Discount
		result := databases.DB.Where("id = ?", data.ID).First(&existingDiscount)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				data.UserID = users[i].ID
				data.OutletID = outlet[i].ID
				data.ProductID = products[i].ID
				tx := databases.DB.Create(&data)
				if tx.Error != nil {
					roll.Rollback()
					log.Fatal("error insert data: ", tx.Error)
				}
			}
		} else {
			log.Println("table already exist")
			break
		}
	}

	file, err = os.ReadFile("dummy/transactiondummy.json")
	if err != nil {
		log.Fatal("error reading file: ", err)
	}
	var transactions []transactionModel.Transaction
	err = json.Unmarshal(file, &transactions)
	if err != nil {
		log.Fatal("error unmarshalling JSON: ", err)
	}

	for i, data := range transactions {
		var existingTransaction transactionModel.Transaction
		result := databases.DB.Where("id = ?", data.ID).First(&existingTransaction)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				data.UserID = users[i].ID
				data.OutletID = outlet[i].ID
				tx := databases.DB.Create(&data)
				if tx.Error != nil {
					roll.Rollback()
					log.Fatal("error insert data: ", tx.Error)
				}
			}
		} else {
			log.Println("table already exist")
			break
		}
	}

	file, err = os.ReadFile("dummy/transactiondetailsdummy.json")
	if err != nil {
		log.Fatal("error reading file: ", err)
	}
	var transactionsDetails []transactionDetailModel.TransactionDetail
	err = json.Unmarshal(file, &transactionsDetails)
	if err != nil {
		log.Fatal("error unmarshalling JSON: ", err)
	}

	for i, data := range transactionsDetails {
		var existingTransactionDetail transactionDetailModel.TransactionDetail
		result := databases.DB.Where("id = ?", data.ID).First(&existingTransactionDetail)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				data.TransactionID = transactions[i].ID
				tx := databases.DB.Create(&data)
				if tx.Error != nil {
					roll.Rollback()
					log.Fatal("error insert data: ", tx.Error)
				}
			}
		} else {
			log.Println("table already exist")
			break
		}
	}

	roll.Commit()
}
