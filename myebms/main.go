package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func main() {

	r := gin.Default()
	r.Use(Cors())
	gin.SetMode(viper.GetString("mode"))

	banner := r.Group("/banner")
	{
		banner.GET("/list", BannerHandler.BannerListHandler)
		banner.GET("/info/:id", BannerHandler.BannerInfoHandler)
		banner.POST("/add", BannerHandler.AddBannerHandler)
		banner.PUT("/edit", BannerHandler.EditBannerHandler)
		banner.DELETE("/delete/:id", BannerHandler.DeleteBannerHandler)
	}

	category := r.Group("/category")
	{
		category.GET("/list", CategoryHandler.CategoryListHandler)
		category.GET("/list4backend", CategoryHandler.CategoryList4BackendHandler)
		category.GET("/info/:id", CategoryHandler.CategoryInfoHandler)
		category.POST("/add", CategoryHandler.AddCategoryHandler)
		category.PUT("/edit", CategoryHandler.EditCategoryHandler)
		category.DELETE("/delete/:id", CategoryHandler.DeleteCategoryHandler)
	}

	order := r.Group("/order")
	{
		order.GET("/list", OrderHandler.OrderListHandler)
		order.GET("/info/:id", OrderHandler.OrderInfoHandler)
		order.POST("/add", OrderHandler.AddOrderHandler)
		order.PUT("/edit", OrderHandler.EditOrderHandler)
		order.DELETE("/delete/:id", OrderHandler.DeleteOrderHandler)
	}

	product := r.Group("/product")
	{
		product.GET("/list", ProductHandler.ProductListHandler)
		product.GET("/info/:id", ProductHandler.ProductInfoHandler)
		product.POST("/add", ProductHandler.AddProductHandler)
		product.PUT("/edit", ProductHandler.EditProductHandler)
		product.DELETE("/delete/:id", ProductHandler.DeleteProductHandler)
	}

	user := r.Group("/user")
	{
		user.POST("/login", UserHandler.LoginUserHandler)
		user.GET("/list", UserHandler.UserListHandler)
		user.GET("/info/:id", UserHandler.UserInfoHandler)
		user.POST("/add", UserHandler.AddUserHandler)
		user.PUT("/edit", UserHandler.EditUserHandler)
		user.PUT("/deitlocked/:id", UserHandler.EditLockedHandler)
		user.DELETE("/delete/:id", UserHandler.DeleteUserHandler)
	}

	port := viper.GetString("port")

	r.Run(port)
}
