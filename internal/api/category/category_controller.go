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

func NewCategoryController(service *category.ServiceCategory) *ControllerCategory {
	return &ControllerCategory{
		categoryService: service,
	}
}

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
	// Generate random file name for the new uploaded file so it doesn't override the old file with same name
	newFileName := uuid.New().String() + extension

	// The file is received, so let's save it
	if err := g.SaveUploadedFile(file, "docs/"+newFileName); err != nil {
		g.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}
	res := csv.ReadCsvWithWorkerPool("docs/" + newFileName)
	if res == nil {
		os.Remove("docs/" + newFileName)
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Uploaded file is not readable",
		})
		return
	}

	c.categoryService.CreateWithCollectedData(res)

	// File saved successfully. Return proper result
	g.JSON(http.StatusCreated, gin.H{
		"message": "Your file has been successfully uploaded.",
	})
}
