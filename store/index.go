package store

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db       *gorm.DB
	WalletDb *gorm.DB
)

func Init() {
	host := viper.GetString("DB_HOST")
	user := viper.GetString("DB_USER")
	password := viper.GetString("DB_PASSWORD")
	name := viper.GetString("DB_NAME")
	port := viper.GetString("DB_PORT")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, name)
	fmt.Println(dsn)
	mysqlDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	Db = mysqlDb

	walletInit()
}

func walletInit() {
	host := viper.GetString("WALLET_DB_HOST")
	user := viper.GetString("WALLET_DB_USER")
	password := viper.GetString("WALLET_DB_PASSWORD")
	name := viper.GetString("WALLET_DB_NAME")
	port := viper.GetString("WALLET_DB_PORT")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, name)
	fmt.Println(dsn)
	mysqlDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	WalletDb = mysqlDb
}
