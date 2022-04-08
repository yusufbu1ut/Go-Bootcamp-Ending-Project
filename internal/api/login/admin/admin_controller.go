package admin

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/api/login"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/config"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/domain/admin"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/pkg/hashing"
	jwtHelper "github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/pkg/jwt"
	"net/http"
	"os"
	"time"
)

type ControllerAdmin struct {
	appConfig    *config.Configuration
	adminService *admin.ServiceAdmin
}

func NewAdminController(appConfig *config.Configuration, service *admin.ServiceAdmin) *ControllerAdmin {
	return &ControllerAdmin{
		appConfig:    appConfig,
		adminService: service,
	}
}

func (c *ControllerAdmin) Login(g *gin.Context) {
	var req login.RequestLogin
	err := g.ShouldBind(&req)
	if err != nil || req.Email == "" || req.Password == "" {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Check your request body.",
		})
		g.Abort()
		return
	}
	user := c.adminService.GetUser(req.Email, req.Password)
	if user.ID == 0 {
		g.JSON(http.StatusNotFound, gin.H{
			"error_message": "User not found!",
		})
		g.Abort()
		return
	}
	hashedRole, err := hashing.HashWord("admin")
	if err != nil {
		fmt.Println("Error occurred:", err.Error())
	}
	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   user.ID,
		"email":    user.Email,
		"username": user.Username,
		"role":     hashedRole,
		"iat":      time.Now().Unix(),
		"iss":      "admin-service" + os.Getenv("ENV"),
		"exp": time.Now().Add(2 *
			time.Hour).Unix(),
	})
	token := jwtHelper.GenerateToken(jwtClaims, c.appConfig.JwtSettings.SecretKey)
	g.JSON(http.StatusOK, token)
}

func (c *ControllerAdmin) VerifyToken(g *gin.Context) {
	token := g.GetHeader("Authorization")
	decodedClaims := jwtHelper.VerifyToken(token, c.appConfig.JwtSettings.SecretKey, os.Getenv("ENV"))

	g.JSON(http.StatusOK, decodedClaims)

}
