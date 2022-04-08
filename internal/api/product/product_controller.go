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

func NewProductController(service *product.ServiceProduct, service2 *category.ServiceCategory) *ControllerProduct {
	return &ControllerProduct{
		productService:  service,
		categoryService: service2,
	}
}

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

func (c *ControllerProduct) CreateProduct(g *gin.Context) {
	var req product.Product
	err := g.ShouldBind(&req)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Check your request body." + err.Error(),
		})
		g.Abort()
		return
	}
	if req.CategoryID == 0 {
		req.CategoryID = c.categoryService.GetCategoryWithCode(0).ID
	}
	category := c.categoryService.GetCategoryWithId(int(req.CategoryID))
	if category.ID == 0 {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Check your request body. Category id is not found.",
		})
		g.Abort()
		return
	}
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

func (c *ControllerProduct) SearchProduct(g *gin.Context) {
	var products []product.Product
	var req product.Product
	var err error

	name, isOk := g.GetQuery("name")
	categoryId, isOk1 := g.GetQuery("categoryId")
	amount, isOk2 := g.GetQuery("amount")

	if isOk {
		req.Name = name
	}
	if isOk1 {
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

	if len(products) == 0 {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "No data found with search parameters.",
		})
		g.Abort()
		return
	}

	g.JSON(http.StatusOK, products)
}

func (c *ControllerProduct) DeleteProduct(g *gin.Context) {
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
			g.JSON(http.StatusBadRequest, gin.H{
				"error_message": err.Error(),
			})
			g.Abort()
			return
		}
	} else {
		var req product.Product
		err := g.ShouldBind(&req)
		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{
				"error_message": "Check your request body." + err.Error(),
			})
			g.Abort()
			return
		}
		err = c.productService.Delete(&req)
		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{
				"error_message": err.Error(),
			})
			g.Abort()
			return
		}
	}

	g.JSON(http.StatusAccepted, gin.H{
		"message": "Product successfully deleted.",
	})

}

func (c *ControllerProduct) UpdateProduct(g *gin.Context) {
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
