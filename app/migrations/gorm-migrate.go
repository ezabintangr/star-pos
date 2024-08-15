package migrations

import (
	"star-pos/app/configs"
	"star-pos/app/databases"
	cartData "star-pos/features/cart/repository"
	cartProductData "star-pos/features/cart_product/repository"
	categoriesData "star-pos/features/categories/repository"
	discountData "star-pos/features/discount/repository"
	outletData "star-pos/features/outlet/repository"
	productData "star-pos/features/product/repository"
	transactionData "star-pos/features/transaction/repository"
	TransactionDetailData "star-pos/features/transaction_detail/repository"
	userData "star-pos/features/user/repository"
)

func InitMigration() {
	databases.InitMySql(configs.InitConfig()).AutoMigrate(&userData.User{}, &productData.Product{}, &cartData.Cart{}, &cartProductData.CartProduct{}, &categoriesData.Categories{}, &discountData.Discount{}, &outletData.Outlet{}, &transactionData.Transaction{}, &TransactionDetailData.TransactionDetail{})
}
