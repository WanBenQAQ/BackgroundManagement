package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"wanben/myebms/config"
	"wanben/myebms/handler"
	"wanben/myebms/repository"
	"wanben/myebms/model"
	"wanben/myebms/service"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var (
	DB              *gorm.DB
	BannerHandler   handler.BannerHandler
	CategoryHandler handler.CategoryHandler
	OrderHandler    handler.OrderHandler
	ProductHandler  handler.ProductHandler
	UserHandler     handler.UserHandler
)

func initViper() {
	if err := config.Init(""); err != nil {
		panic(err)
	}
}

func initDB() {
	fmt.Println("数据库 init")
	var err error
	conf := &model.DBConf{
		Host:     viper.GetString("database.host"),
		User:     viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		DbName:   viper.GetString("database.name"),
	}

	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8&parseTime=%t&loc=%s",
		conf.User,
		conf.Password,
		conf.Host,
		conf.DbName,
		true,
		"Local")

	DB, err = gorm.Open("mysql", config)
	if err != nil {
		log.Fatalf("connect error: %v\n", err)
	}
	DB.SingularTable(true)
	// 在数据库中创建表结构
	DB.AutoMigrate(
		&model.Banner{},
		&model.Category{},
		&model.CategoryResult{},
		&model.Order{},
		&model.OrderItem{},
		&model.Product{},
		&model.ShoppingCartItem{},
		&model.Stock{},
		&model.User{},
	)
	fmt.Println("数据库 init 结束...")
}

func initHandler() {
	BannerHandler = handler.BannerHandler{
		BannerSrv: &service.BannerService{
			Repo: &repository.BannerRepository{
				DB: DB,
			},
		}}

	CategoryHandler = handler.CategoryHandler{
		CategorySrv: &service.CategoryService{
			Repo: &repository.CategoryRepository{
				DB: DB,
			},
		},
	}

	OrderHandler = handler.OrderHandler{
		OrderSrv: &service.OrderService{
			Repo: &repository.OrderRepository{
				DB: DB,
			},
		}}

	ProductHandler = handler.ProductHandler{
		ProductSrv: &service.ProductService{
			Repo: &repository.ProductRepository{
				DB: DB,
			},
		}}

	UserHandler = handler.UserHandler{
		UserSrv: &service.UserService{
			Repo: &repository.UserRepository{
				DB: DB,
			},
		}}
}

func init() {
	initViper()
	initDB()
	initHandler()
}
