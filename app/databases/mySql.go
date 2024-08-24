package databases

import (
	"database/sql"
	"fmt"
	"star-pos/app/configs"

	cartData "star-pos/features/cart/model"
	cartProductData "star-pos/features/cart_product/model"
	categoriesData "star-pos/features/categories/model"
	discountData "star-pos/features/discount/model"
	outletData "star-pos/features/outlet/model"
	productData "star-pos/features/product/model"
	transactionData "star-pos/features/transaction/model"
	TransactionDetailData "star-pos/features/transaction_detail/model"
	userModel "star-pos/features/user/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var DBSql *sql.DB

func InitMySql(cfg *configs.AppConfig) {
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

	DBSql, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	err = DBSql.Ping()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&userModel.User{}, &productData.Product{}, &cartData.Cart{}, &cartProductData.CartProduct{}, &categoriesData.Categories{}, &discountData.Discount{}, &outletData.Outlet{}, &transactionData.Transaction{}, &TransactionDetailData.TransactionDetail{})
	DB = db
}
