package databases

import (
	"fmt"
	"star-pos/app/configs"

	cartData "star-pos/features/cart/repository"
	cartProductData "star-pos/features/cart_product/repository"
	categoriesData "star-pos/features/categories/repository"
	discountData "star-pos/features/discount/repository"
	outletData "star-pos/features/outlet/repository"
	productData "star-pos/features/product/repository"
	transactionData "star-pos/features/transaction/repository"
	TransactionDetailData "star-pos/features/transaction_detail/repository"
	userData "star-pos/features/user/repository"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitMySql(cfg *configs.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOSTNAME, cfg.DB_PORT, cfg.DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&userData.User{}, &productData.Product{}, &cartData.Cart{}, &cartProductData.CartProduct{}, &categoriesData.Categories{}, &discountData.Discount{}, &outletData.Outlet{}, &transactionData.Transaction{}, &TransactionDetailData.TransactionDetail{})
	return db
}
