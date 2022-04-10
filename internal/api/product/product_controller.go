package product

import (
	"github.com/gin-gonic/gin"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/domain/category"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/domain/product"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/pkg/pagination"
	"net/http"
	"strconv"
)

type ControllerProduct struct {
	productService  *product.ServiceProduct
	categoryService *category.ServiceCategory
}

// @BasePath /product

func NewProductController(service *product.ServiceProduct, service2 *category.ServiceCategory) *ControllerProduct {
	return &ControllerProduct{
		productService:  service,
		categoryService: service2,
	}
}

// GetAll godoc
// @Summary Gets all products with pagination parameters page and size
// @Tags Product
// @Accept  json
// @Produce  json
// @Param page query int false "Page Index"
// @Param size query int false "Page Size"
// @Success 200 {object} pagination.Pages
// @Failure 404 {object} map[string]string
// @Router /product [get]
func (c *ControllerProduct) GetAll(g *gin.Context) {
	pageIndex, pageSize := pagination.GetPaginationParametersFromRequest(g)
	products, count := c.productService.GetAll(pageIndex, pageSize)
	if len(products) == 0 {
		g.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "No Data found with this pagination.",
		})
		return
	}
	for i, _ := range products {
		products[i].Category = c.categoryService.GetCategoryWithId(int(products[i].CategoryID))
	}
	paginatedResult := pagination.CreatePagingResponse(g, count)
	paginatedResult.Items = products

	g.JSON(http.StatusOK, paginatedResult)
}

// Create godoc
// @Summary Creates products with the given request
// @Tags Product
// @Accept  json
// @Produce  json
// @Param product-request body ResponseProduct true "Takes the products and adds them to db"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Router /product [post]
func (c *ControllerProduct) Create(g *gin.Context) {
	var req product.Product
	err := g.ShouldBind(&req)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Check your request body." + err.Error(),
		})
		g.Abort()
		return
	}
	//there is no category id it will be default category named "other"
	if req.CategoryID == 0 {
		req.CategoryID = c.categoryService.GetCategoryWithCode(0).ID
	}
	//checking given request category id if is not exist returns
	category := c.categoryService.GetCategoryWithId(int(req.CategoryID))
	if category.ID == 0 {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Check your request body. Category id is not found.",
		})
		g.Abort()
		return
	}
	//creating product
	productItem := product.NewProduct(req.Name, req.Price, req.Amount, req.Code, req.Description, req.CategoryID)
	err = c.productService.Create(productItem)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": err.Error(),
		})
		g.Abort()
		return
	}

	g.JSON(http.StatusCreated, gin.H{
		"message": "Product successfully created.",
	})
}

// Search godoc
// @Summary Gets all products with search parameters name, amount or categoryId
// @Tags Product
// @Accept  json
// @Produce  json
// @Param name query string false "Name"
// @Param categoryId query int false "CategoryId"
// @Param amount query int false "Amount"
// @Success 200 {object} ResponseProduct
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /product/search [get]
//Search works with query parameters name, category id and amount, amount process looks equal or if there is more
func (c *ControllerProduct) Search(g *gin.Context) {
	var products []product.Product
	var req product.Product
	var err error

	//taking query parameters
	name, isOk := g.GetQuery("name")
	categoryId, isOk1 := g.GetQuery("categoryId")
	amount, isOk2 := g.GetQuery("amount")

	if isOk {
		req.Name = name
	}
	if isOk1 {
		//checking category-id query parameter
		intCategory, err := strconv.Atoi(categoryId)
		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{
				"error_message": err.Error(),
			})
			g.Abort()
			return
		}
		if intCategory <= 0 {
			g.JSON(http.StatusBadRequest, gin.H{
				"error_message": "Invalid category-id",
			})
			g.Abort()
			return
		}
		req.CategoryID = uint(intCategory)
	}
	if isOk2 {
		//checking amount query parameter
		intAmount, err := strconv.Atoi(amount)
		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{
				"error_message": err.Error(),
			})
			g.Abort()
			return
		}
		if intAmount < 0 {
			g.JSON(http.StatusBadRequest, gin.H{
				"error_message": "Invalid amount",
			})
			g.Abort()
			return
		}
		req.Amount = uint(intAmount)
	}
	//if query parameters exist
	if isOk || isOk1 || isOk2 {
		products, err = c.productService.Search(&req)
		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{
				"error_message": err.Error(),
			})
			g.Abort()
			return
		}
	}
	//if there is no product
	if len(products) == 0 {
		g.JSON(http.StatusNotFound, gin.H{
			"error_message": "No data found with search parameters.",
		})
		g.Abort()
		return
	}
	//shaping response products
	var response []ResponseProduct
	for _, p := range products {
		var resProduct ResponseProduct
		resProduct.ID = p.ID
		resProduct.Name = p.Name
		resProduct.Amount = p.Amount
		resProduct.Code = p.Code
		resProduct.Price = p.Price
		resProduct.Description = p.Description
		resProduct.CategoryId = p.CategoryID
		response = append(response, resProduct)
	}

	g.JSON(http.StatusOK, response)
}

// Delete godoc
// @Summary Deletes given product id, it can be with query or request body
// @Tags Product
// @Accept  json
// @Produce  json
// @Param id query int false "Takes the product id and deletes it, product id is needed"
// @Param product-request body ResponseProduct false "Takes the product infos and deletes it, product id is needed"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Router /product [delete]
//Delete works with query parameter product id and also works with product request body it should contain exist product id
func (c *ControllerProduct) Delete(g *gin.Context) {
	//checking delete from query
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

		err = c.productService.DeleteWithID(deg)
		if err != nil {
			g.JSON(http.StatusNotFound, gin.H{
				"error_message": err.Error(),
			})
			g.Abort()
			return
		}
	} else {
		//checking delete from request body
		var req product.Product
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
				"error_message": "Check your request body. Id is need",
			})
			g.Abort()
			return
		}
		err = c.productService.Delete(&req)
		if err != nil {
			g.JSON(http.StatusNotFound, gin.H{
				"error_message": err.Error(),
			})
			g.Abort()
			return
		}
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Product successfully deleted.",
	})

}

// Update godoc
// @Summary Updates given product
// @Tags Product
// @Accept  json
// @Produce  json
// @Param product-request body ResponseProduct true "Takes the product infos and updates it. Product id, name, category id and code needed"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Router /product [put]
//Update works with all product fields and changes on it
func (c *ControllerProduct) Update(g *gin.Context) {
	var req product.Product
	err := g.ShouldBind(&req)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Check your request body." + err.Error(),
		})
		g.Abort()
		return
	}

	err = c.productService.Update(&req)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": err.Error(),
		})
		g.Abort()
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Product successfully updated.",
	})
}
