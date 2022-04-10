package basket

import (
	"github.com/gin-gonic/gin"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/config"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/domain/basket"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/domain/customer"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/domain/order"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/domain/product"
	jwtHelper "github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/pkg/jwt"
	"net/http"
	"os"
	"strconv"
)

type ControllerBasket struct {
	appConfig       *config.Configuration
	basketService   *basket.ServiceBasket
	productService  *product.ServiceProduct
	customerService *customer.ServiceCustomer
}

// @BasePath /basket

func NewBasketController(appConfig *config.Configuration, service *basket.ServiceBasket, service2 *product.ServiceProduct, service3 *customer.ServiceCustomer) *ControllerBasket {
	return &ControllerBasket{
		appConfig:       appConfig,
		basketService:   service,
		productService:  service2,
		customerService: service3,
	}
}

// GetAll godoc
// @Summary Gets all basket items for the authed customer
// @Tags Basket
// @Accept  json
// @Produce  json
// @Success 201 {object} basket.ResponseBasket
// @Failure 404 {object} map[string]string
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Router /basket [get]
func (c *ControllerBasket) GetAll(g *gin.Context) {
	decodedClaims := jwtHelper.VerifyToken(g.GetHeader("Authorization"), c.appConfig.SecretKey, os.Getenv("ENV"))
	basket := c.basketService.GetAll(decodedClaims.UserId)
	if len(basket) == 0 {
		g.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "No Data found.",
		})
		return
	}
	var response ResponseBasket
	//shaping response
	customerUser := c.customerService.GetUserWithId(decodedClaims.UserId)
	response.ID = customerUser.ID
	response.Name = customerUser.Name
	response.Email = customerUser.Email
	response.Username = customerUser.Username
	response.PhoneNo = customerUser.PhoneNo
	response.Address = customerUser.Address
	for _, b := range basket {
		var responseProduct ResponseProduct
		product := c.productService.GetById(int(b.ProductID))
		responseProduct.ID = product.ID
		responseProduct.Name = product.Name
		responseProduct.Amount = b.Amount
		response.Products = append(response.Products, responseProduct)
	}
	g.JSON(http.StatusOK, response)
}

// Create godoc
// @Summary Adds the given items to basket
// @Tags Basket
// @Accept  json
// @Produce  json
// @Param basket-request body basket.ResponseProduct true "Takes the products and adds them"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Router /basket [post]
//Create adds products to the basket, request should contain product id and amount
func (c *ControllerBasket) Create(g *gin.Context) {
	decodedClaims := jwtHelper.VerifyToken(g.GetHeader("Authorization"), c.appConfig.SecretKey, os.Getenv("ENV"))
	var req []ResponseProduct
	err := g.ShouldBind(&req)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Check your request body." + err.Error(),
		})
		g.Abort()
		return
	}
	//looking request items to add db
	var baskets []order.Order
	for j, r := range req {
		//checking request amount is it zero value and request ID
		if r.Amount == 0 || r.ID == 0 {
			g.JSON(http.StatusBadRequest, gin.H{
				"error_message": "Check your request body. ID and Amount needed. Amount can not be zero.",
			})
			g.Abort()
			return
		}
		//checking is there same productItem in same request
		for i, _ := range req {
			if j != i && r.ID == req[i].ID {
				g.JSON(http.StatusBadRequest, gin.H{
					"error_message": "Check your request body. There cant be same product item in one request",
				})
				g.Abort()
				return
			}
		}
		//checking given request product id for is there such as product with request product id
		productItem := c.productService.GetById(int(r.ID))
		if productItem.ID == 0 {
			g.JSON(http.StatusNotFound, gin.H{
				"error_message": "Product not exist . No product found with id: " + strconv.Itoa(int(r.ID)),
			})
			g.Abort()
			return
		}
		//checking is amount higher than stock amount
		if int(productItem.Amount)-int(r.Amount) < 0 {
			g.JSON(http.StatusBadRequest, gin.H{
				"error_message": "Given amount can not be higher than stock with id:" + strconv.Itoa(int(productItem.ID)) + " Stock is :" + strconv.Itoa(int(productItem.Amount)),
			})
			g.Abort()
			return
		}
		//creating (order)basket item
		basketItem := order.NewOrder(uint(decodedClaims.UserId), productItem.ID, r.Amount)
		//shaping product
		productItem.Amount = productItem.Amount - basketItem.Amount
		basketItem.Product = productItem
		baskets = append(baskets, *basketItem)
	}

	//updating products in product db
	for _, b := range baskets {
		err := c.productService.Update(&b.Product)
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{
				"error_message": err.Error(),
			})
			g.Abort()
			return
		}
	}

	//adding basket items in order db
	err = c.basketService.AddBasket(baskets)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error_message": err.Error(),
		})
		g.Abort()
		return
	}

	g.JSON(http.StatusCreated, gin.H{
		"message": "Basket items successfully added.",
	})
}

// Update godoc
// @Summary Updates given basket item, only amount can be updated
// @Tags Basket
// @Accept  json
// @Produce  json
// @Param basket-request body basket.ResponseProduct true "Takes the basket infos and updates it, basket id and to update amount is needed"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Router /basket [patch]
//Update request should contain product id and amount
func (c *ControllerBasket) Update(g *gin.Context) {
	decodedClaims := jwtHelper.VerifyToken(g.GetHeader("Authorization"), c.appConfig.SecretKey, os.Getenv("ENV"))
	var basketItem order.Order
	basketItem.CustomerID = uint(decodedClaims.UserId)

	var req ResponseProduct
	err := g.ShouldBind(&req)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Check your request body." + err.Error(),
		})
		g.Abort()
		return
	}
	//request should contain id and amount
	if req.ID == 0 || req.Amount == 0 {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Product id and amount is needed.",
		})
		g.Abort()
		return
	}
	basketItem.ProductID = req.ID
	basketItem.Amount = req.Amount

	//finding item with request
	basketItemDB := c.basketService.GetBasket(basketItem)
	if basketItemDB.ID == 0 {
		g.JSON(http.StatusNotFound, gin.H{
			"error_message": "Basket item not found.",
		})
		g.Abort()
		return
	}
	//checking product item for amount process
	productItem := c.productService.GetById(int(basketItemDB.ProductID))
	if productItem.ID == 0 {
		g.JSON(http.StatusNotFound, gin.H{
			"error_message": "Product item not found. Product with id: " + strconv.Itoa(int(basketItemDB.ProductID)),
		})
		g.Abort()
		return
	}
	//update amount counting
	count := int(basketItem.Amount) - int(basketItemDB.Amount)
	if int(productItem.Amount)-count < 0 {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Given amount higher than product stock. Stock is :" + strconv.Itoa(int(productItem.Amount)),
		})
		g.Abort()
		return
	}
	//updating total product amount
	productItem.Amount = uint(int(productItem.Amount) - count)
	err = c.productService.Update(&productItem)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error_message": err.Error(),
		})
		g.Abort()
		return
	}
	//updating total basket amount
	basketItemDB.Amount = basketItem.Amount
	err = c.basketService.Update(basketItemDB)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error_message": err.Error(),
		})
		g.Abort()
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Basket item successfully updated.",
	})

}

// Delete godoc
// @Summary Deletes given basket products id, it can be with query or request body
// @Tags Basket
// @Accept  json
// @Produce  json
// @Param id query int false "Takes the basket product infos and deletes it, product id is needed"
// @Param product-request body basket.ResponseProduct false "Takes the basket product infos and deletes it, product id is needed"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Router /basket [delete]
//Delete needs only product ID from basket items it can be sended from query and body
func (c *ControllerBasket) Delete(g *gin.Context) {
	decodedClaims := jwtHelper.VerifyToken(g.GetHeader("Authorization"), c.appConfig.SecretKey, os.Getenv("ENV"))
	var basketItem order.Order
	basketItem.CustomerID = uint(decodedClaims.UserId)
	//checking query for product id
	id, isOK := g.GetQuery("id")
	if isOK {
		deg, err := strconv.Atoi(id)
		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{
				"error_message": "Query count of id should be an integer;" + err.Error(),
			})
			g.Abort()
			return
		}
		basketItem.ProductID = uint(deg)
	} else {
		//checking request for product id
		var req ResponseProduct
		err := g.ShouldBind(&req)
		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{
				"error_message": "Check your request body." + err.Error(),
			})
			g.Abort()
			return
		}
		if req.ID == 0 {
			g.JSON(http.StatusBadRequest, gin.H{
				"error_message": "Request id is needed",
			})
			g.Abort()
			return
		}
		basketItem.ProductID = req.ID
	}
	if basketItem.ProductID == 0 {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Product id is needed to delete process.",
		})
		g.Abort()
		return
	}
	//getting basket item from db
	basketItem = c.basketService.GetBasket(basketItem)
	//checking is there same basket with given basket
	if basketItem.ID == 0 {
		g.JSON(http.StatusNotFound, gin.H{
			"error_message": "Basket item not found.",
		})
		g.Abort()
		return
	}
	//checking is there still the product exist
	productItem := c.productService.GetById(int(basketItem.ProductID))
	if productItem.ID > 0 {
		//updating product amount adding back from basket amount
		productItem.Amount = productItem.Amount + basketItem.Amount
		err := c.productService.Update(&productItem)
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{
				"error_message": "Check your request body." + err.Error(),
			})
			g.Abort()
			return
		}
	}

	err := c.basketService.DeleteBasket(basketItem)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error_message": "Check your request body." + err.Error(),
		})
		g.Abort()
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Basket item successfully deleted.",
	})
}

// Complete godoc
// @Summary Completes baskets to orders with basket products
// @Tags Basket
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Router /basket/complete [post]
//Complete makes order with basket items
func (c *ControllerBasket) Complete(g *gin.Context) {
	decodedClaims := jwtHelper.VerifyToken(g.GetHeader("Authorization"), c.appConfig.SecretKey, os.Getenv("ENV"))
	basketItems := c.basketService.GetAll(decodedClaims.UserId)
	//checking last time is there still same product exist
	for _, item := range basketItems {
		productItem := c.productService.GetById(int(item.ProductID))
		if productItem.ID == 0 {
			g.JSON(http.StatusNotFound, gin.H{
				"error_message": "Basket product not exist. Product with id:" + strconv.Itoa(int(item.ProductID)),
			})
			g.Abort()
			return
		}
	}
	//Completing order
	err := c.basketService.CompleteOrder(decodedClaims.UserId)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": err.Error(),
		})
		g.Abort()
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Basket items successfully ordered.",
	})
}
