package order

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/config"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/domain/customer"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/domain/order"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/domain/product"
	jwtHelper "github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/pkg/jwt"
)

type ControllerOrder struct {
	appConfig       *config.Configuration
	orderService    *order.ServiceOrder
	productService  *product.ServiceProduct
	customerService *customer.ServiceCustomer
}

// @BasePath /order

func NewOrderController(appConfig *config.Configuration, service *order.ServiceOrder, service2 *product.ServiceProduct, service3 *customer.ServiceCustomer) *ControllerOrder {
	return &ControllerOrder{
		appConfig:       appConfig,
		orderService:    service,
		productService:  service2,
		customerService: service3,
	}
}

// GetAll godoc
// @Summary Gets all order items for the authed customer
// @Tags Order
// @Accept  json
// @Produce  json
// @Success 200 {object} ResponseOrder
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Router /order [get]
//GetAll gets customers all uncanceled orders
func (c *ControllerOrder) GetAll(g *gin.Context) {
	decodedClaims := jwtHelper.VerifyToken(g.GetHeader("Authorization"), c.appConfig.SecretKey, os.Getenv("ENV"))
	var response ResponseOrder
	//getting all orders for the customer
	orders := c.orderService.GetAll(decodedClaims.UserId)
	if len(orders) == 0 {
		g.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "No Data found.",
		})
		return
	}
	//shaping response
	customer := c.customerService.GetUserWithId(decodedClaims.UserId)
	response.ID = customer.ID
	response.Name = customer.Name
	response.Email = customer.Email
	response.Username = customer.Username
	response.PhoneNo = customer.PhoneNo
	response.Address = customer.Address
	for _, o := range orders {
		var resProduct ResponseProduct
		productItem := c.productService.GetById(int(o.ProductID))
		if productItem.ID == 0 {
			continue
		}
		resProduct.ID = productItem.ID
		resProduct.OrderTime = o.UpdatedAt
		resProduct.Name = productItem.Name
		resProduct.Amount = o.Amount
		resProduct.Code = o.OrderCode
		response.Products = append(response.Products, resProduct)
	}
	g.JSON(http.StatusOK, response)
}

// Cancel godoc
// @Summary Cancels the orders with given order code
// @Tags Order
// @Accept  json
// @Produce  json
// @Param code query string true "Takes the order codes and cancel them"
// @Success 200 {object} map[string]string
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Router /order/cancel [patch]
//Cancel cancels the orders with
func (c *ControllerOrder) Cancel(g *gin.Context) {
	decodedClaims := jwtHelper.VerifyToken(g.GetHeader("Authorization"), c.appConfig.SecretKey, os.Getenv("ENV"))
	code := g.Query("code")
	if code == "" {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Query parameter order code is needed.",
		})
		g.Abort()
		return
	}
	orders := c.orderService.GetWithCode(decodedClaims.UserId, code)
	if orders == nil {
		g.JSON(http.StatusNotFound, gin.H{
			"error_message": "No order found to cancel with code: " + code,
		})
		g.Abort()
		return
	}
	for _, o := range orders {
		productItem := c.productService.GetById(int(o.ProductID))
		//checking is there still the product exist
		if productItem.ID > 0 {
			productItem.Amount = productItem.Amount + o.Amount
			err := c.productService.Update(&productItem)
			if err != nil {
				g.JSON(http.StatusInternalServerError, gin.H{
					"error_message": err.Error(),
				})
				g.Abort()
				return
			}
		}
	}
	err := c.orderService.CancelOrders(orders)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error_message": err.Error(),
		})
		g.Abort()
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Order is successfully canceled.",
	})

}
