package api

import (
	"github.com/gin-gonic/gin"
	basket2 "github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/api/basket"
	categoryApi "github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/api/category"
	adminApi "github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/api/login/admin"
	customerApi "github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/api/login/customer"
	orderApi "github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/api/order"
	productApi "github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/api/product"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/api/sign_up"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/config"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/domain/admin"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/domain/basket"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/domain/category"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/domain/customer"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/domain/order"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/domain/product"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/pkg/database_handler"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/pkg/middleware"
	"log"
)

var AppConfig = &config.Configuration{}

func RegisterHandlers(r *gin.Engine, cfgFile string) {
	//Reading configuration file
	AppConfig, err := config.GetAllConfigValues(cfgFile)
	if err != nil {
		log.Fatalf("Failed to read configs file. %v", err.Error())
	}

	//connection to MySQL DB
	db := database_handler.MySQLDBConnect(AppConfig.DatabaseSettings.DatabaseURI)

	//Creating repositories and adding migrations
	repoAdmin := admin.NewRepositoryAdmin(db)
	repoAdmin.Migration()
	repoAdmin.InsertSampleData() //Adding admins

	repoCategory := category.NewRepositoryCategory(db)
	repoCategory.Migration()
	repoCategory.InsertSampleData() //Adding just one data named Other for not matched product category-id

	repoCustomer := customer.NewRepositoryCustomer(db)
	repoCustomer.Migration()

	repoBasket := basket.NewRepositoryBasket(db)
	repoBasket.Migration()

	repoOrder := order.NewRepositoryOrder(db)
	repoOrder.Migration()

	repoProduct := product.NewRepositoryProduct(db)
	repoProduct.Migration()

	//Creating services
	servAdmin := admin.NewServiceAdmin(repoAdmin)
	servCustomer := customer.NewServiceCustomer(repoCustomer)
	servCategory := category.NewServiceCategory(repoCategory)
	servProduct := product.NewServiceProduct(repoProduct)
	servOrder := order.NewServiceOrder(repoOrder)
	servBasket := basket.NewServiceBasket(repoBasket)

	//Creating Controllers
	adminController := adminApi.NewAdminController(AppConfig, servAdmin)
	customerController := customerApi.NewCustomerController(AppConfig, servCustomer)
	categoryController := categoryApi.NewCategoryController(servCategory)
	signupController := sign_up.NewSignupController(AppConfig, servCustomer)
	productController := productApi.NewProductController(servProduct, servCategory)
	orderController := orderApi.NewOrderController(AppConfig, servOrder, servProduct, servCustomer)
	basketController := basket2.NewBasketController(AppConfig, servBasket, servProduct, servCustomer)

	//Router Groups
	loginGroup := r.Group("/login")
	loginGroup.POST("/admin", adminController.Login)
	//loginGroup.GET("/admin", middleware.AdminMiddleware(AppConfig.JwtSettings.SecretKey), adminController.VerifyToken)

	loginGroup.POST("/customer", customerController.Login)
	//loginGroup.GET("/customer", middleware.CustomerMiddleware(AppConfig.JwtSettings.SecretKey), customerController.VerifyToken)

	signupGroup := r.Group("/signup")
	signupGroup.POST("", signupController.Signup)

	categoryGroup := r.Group("/category")
	categoryGroup.GET("", categoryController.GetAll)
	categoryGroup.POST("", middleware.AdminMiddleware(AppConfig.JwtSettings.SecretKey), categoryController.CreateWithCollectedData)

	productGroup := r.Group("/product")
	productGroup.GET("", productController.GetAll)
	productGroup.GET("/search", productController.Search)
	productGroup.POST("", middleware.AdminMiddleware(AppConfig.JwtSettings.SecretKey), productController.Create)
	productGroup.PUT("", middleware.AdminMiddleware(AppConfig.JwtSettings.SecretKey), productController.Update)
	productGroup.DELETE("", middleware.AdminMiddleware(AppConfig.JwtSettings.SecretKey), productController.Delete)

	basketGroup := r.Group("/basket")
	basketGroup.GET("", middleware.CustomerMiddleware(AppConfig.JwtSettings.SecretKey), basketController.GetAll)
	basketGroup.POST("", middleware.CustomerMiddleware(AppConfig.JwtSettings.SecretKey), basketController.Create)
	basketGroup.DELETE("", middleware.CustomerMiddleware(AppConfig.JwtSettings.SecretKey), basketController.Delete)
	basketGroup.PATCH("", middleware.CustomerMiddleware(AppConfig.JwtSettings.SecretKey), basketController.Update)
	basketGroup.POST("/complete", middleware.CustomerMiddleware(AppConfig.JwtSettings.SecretKey), basketController.Complete)

	orderGroup := r.Group("/order", middleware.CustomerMiddleware(AppConfig.JwtSettings.SecretKey))
	orderGroup.GET("", middleware.CustomerMiddleware(AppConfig.JwtSettings.SecretKey), orderController.GetAll)
	orderGroup.PATCH("/cancel", middleware.CustomerMiddleware(AppConfig.JwtSettings.SecretKey), orderController.Cancel)

}
