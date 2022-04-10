package category

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/domain/category"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/pkg/csv"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/pkg/pagination"
	"net/http"
	"os"
	"path/filepath"
)

type ControllerCategory struct {
	categoryService *category.ServiceCategory
}

// @BasePath /category

func NewCategoryController(service *category.ServiceCategory) *ControllerCategory {
	return &ControllerCategory{
		categoryService: service,
	}
}

// GetAll godoc
// @Summary Gets all categories with pagination parameters page and size
// @Tags Category
// @Accept  json
// @Produce  json
// @Param page query int false "Page Index"
// @Param size query int false "Page Size"
// @Success 200 {object} pagination.Pages
// @Failure 404 {object} map[string]string
// @Router /category [get]
//GetAll func needs page and size values from query if there is not, shows with default
func (c *ControllerCategory) GetAll(g *gin.Context) {
	pageIndex, pageSize := pagination.GetPaginationParametersFromRequest(g)
	categories, count := c.categoryService.GetAll(pageIndex, pageSize)
	if len(categories) == 0 {
		g.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "No Data found with this pagination.",
		})
		return
	}
	paginatedResult := pagination.CreatePagingResponse(g, count)
	paginatedResult.Items = categories

	g.JSON(http.StatusOK, paginatedResult)
}

// CreateWithCollectedData godoc
// @Summary Create categories with given csv file uploads it after that reads and creates (adding form-data not implemented for swagger)
// @Tags Category
// @Accept  multipart/form-data
// @Produce  json
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Router /category [post]
//CreateWithCollectedData needs form-data from body key should be named as "file"
func (c *ControllerCategory) CreateWithCollectedData(g *gin.Context) {
	file, err := g.FormFile("file")
	// The file cannot be received.
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}
	// Retrieve file information
	extension := filepath.Ext(file.Filename)
	// Generate random file name for the new uploaded file, so it doesn't override the old file with same name
	newFileName := uuid.New().String() + extension

	// The file is received, saving
	if err := g.SaveUploadedFile(file, "assets/"+newFileName); err != nil {
		g.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}
	res := csv.ReadCsvWithWorkerPool("assets/" + newFileName)
	// if file is not readable deletes the file that uploaded
	if res == nil {
		os.Remove("assets/" + newFileName)
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Uploaded file is not readable",
		})
		return
	}

	c.categoryService.CreateWithCollectedData(res)

	// File saved successfully
	g.JSON(http.StatusCreated, gin.H{
		"message": "Your file has been successfully uploaded.",
	})
}
