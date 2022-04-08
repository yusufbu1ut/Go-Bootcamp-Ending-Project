package order

import (
	"github.com/gin-gonic/gin"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/config"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/domain/order"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/domain/product"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/pkg/pagination"
	"net/http"
)

type ControllerOrder struct {
	appConfig      *config.Configuration
	orderService   *order.ServiceOrder
	productService *product.ServiceProduct
}

func NewOrderController(appConfig *config.Configuration, service *order.ServiceOrder, service2 *product.ServiceProduct) *ControllerOrder {
	return &ControllerOrder{
		appConfig:      appConfig,
		orderService:   service,
		productService: service2,
	}
}

func (c *ControllerOrder) GetAll(g *gin.Context) {
	pageIndex, pageSize := pagination.GetPaginationParametersFromRequest(g)
	orders, count := c.orderService.GetAll(pageIndex, pageSize)
	if len(orders) == 0 {
		g.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "No Data found with this pagination.",
		})
		return
	}
	paginatedResult := pagination.CreatePagingResponse(g, count)
	paginatedResult.Items = orders

	g.JSON(http.StatusOK, paginatedResult)
}
