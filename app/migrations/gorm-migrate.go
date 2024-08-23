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

	var count int
	err = databases.DBSql.QueryRow("SELECT COUNT(id) FROM users").Scan(&count)
	if err != nil {
		log.Fatal("error executing query: ", err)
	}

	if count > 0 {
		log.Println("table already exist")
	} else {
		for _, data := range users {
			tx := databases.DB.Create(&data)
			if tx.Error != nil {
				roll.Rollback()
				log.Fatal("error insert data: ", tx.Error)
			}
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

	err = databases.DBSql.QueryRow("SELECT COUNT(id) FROM outlets").Scan(&count)
	if err != nil {
		log.Fatal("error executing query: ", err)
	}

	if count > 0 {
		log.Println("table already exist")
	} else {
		for i, data := range outlet {
			data.UserID = users[i].ID
			tx := databases.DB.Create(&data)
			if tx.Error != nil {
				log.Fatal("error insert data: ", tx.Error)
			}
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

	err = databases.DBSql.QueryRow("SELECT COUNT(id) FROM categories").Scan(&count)
	if err != nil {
		log.Fatal("error executing query: ", err)
	}

	if count > 0 {
		log.Println("table already exist")
	} else {
		for i, data := range categories {
			data.UserID = users[i].ID
			tx := databases.DB.Create(&data)
			if tx.Error != nil {
				roll.Rollback()
				log.Fatal("error insert data: ", tx.Error)
			}
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

	err = databases.DBSql.QueryRow("SELECT COUNT(id) FROM products").Scan(&count)
	if err != nil {
		log.Fatal("error executing query: ", err)
	}

	if count > 0 {
		log.Println("table already exist")
	} else {
		for i, data := range products {
			data.UserID = users[i].ID
			data.CategoriesID = categories[i].ID
			tx := databases.DB.Create(&data)
			if tx.Error != nil {
				roll.Rollback()
				log.Fatal("error insert data: ", tx.Error)
			}
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

	err = databases.DBSql.QueryRow("SELECT COUNT(id) FROM carts").Scan(&count)
	if err != nil {
		log.Fatal("error executing query: ", err)
	}

	if count > 0 {
		log.Println("table already exist")
	} else {
		for i, data := range carts {
			_, err = time.Parse("2006-01-02", data.Date.Format("2006-01-02"))
			if err != nil {
				log.Fatal("error parse date: ", err)
			}
			data.UserID = users[i].ID
			data.OutletID = outlet[i].ID
			tx := databases.DB.Create(&data)
			if tx.Error != nil {
				roll.Rollback()
				log.Fatal("error insert data: ", tx.Error)
			}
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

	err = databases.DBSql.QueryRow("SELECT COUNT(id) FROM cart_products").Scan(&count)
	if err != nil {
		log.Fatal("error executing query: ", err)
	}

	if count > 0 {
		log.Println("table already exist")
	} else {
		for i, data := range cartProduct {
			data.CartID = carts[i].ID
			data.ProductID = products[i].ID
			tx := databases.DB.Create(&data)
			if tx.Error != nil {
				roll.Rollback()
				log.Fatal("error insert data: ", tx.Error)
			}
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

	err = databases.DBSql.QueryRow("SELECT COUNT(id) FROM discounts").Scan(&count)
	if err != nil {
		log.Fatal("error executing query: ", err)
	}

	if count > 0 {
		log.Println("table already exist")
	} else {
		for i, data := range discounts {
			data.UserID = users[i].ID
			data.OutletID = outlet[i].ID
			data.ProductID = products[i].ID
			tx := databases.DB.Create(&data)
			if tx.Error != nil {
				roll.Rollback()
				log.Fatal("error insert data: ", tx.Error)
			}
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

	err = databases.DBSql.QueryRow("SELECT COUNT(id) FROM transactions").Scan(&count)
	if err != nil {
		log.Fatal("error executing query: ", err)
	}

	if count > 0 {
		log.Println("table already exist")
	} else {
		for i, data := range transactions {
			data.UserID = users[i].ID
			data.OutletID = outlet[i].ID
			tx := databases.DB.Create(&data)
			if tx.Error != nil {
				roll.Rollback()
				log.Fatal("error insert data: ", tx.Error)
			}
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

	err = databases.DBSql.QueryRow("SELECT COUNT(id) FROM transaction_details").Scan(&count)
	if err != nil {
		log.Fatal("error executing query: ", err)
	}

	if count > 0 {
		log.Println("table already exist")
	} else {
		for i, data := range transactionsDetails {
			data.TransactionID = transactions[i].ID
			tx := databases.DB.Create(&data)
			if tx.Error != nil {
				roll.Rollback()
				log.Fatal("error insert data: ", tx.Error)
			}
		}
	}

	roll.Commit()
}
