package customer

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/api/login"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/config"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/domain/customer"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/pkg/hashing"
	jwtHelper "github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/pkg/jwt"
	"net/http"
	"os"
	"time"
)

type ControllerCustomer struct {
	appConfig       *config.Configuration
	customerService *customer.ServiceCustomer
}

func NewCustomerController(appConfig *config.Configuration, service *customer.ServiceCustomer) *ControllerCustomer {
	return &ControllerCustomer{
		appConfig:       appConfig,
		customerService: service,
	}
}

func (c *ControllerCustomer) Login(g *gin.Context) {
	var req login.RequestLogin
	err := g.ShouldBind(&req)
	if err != nil || req.Email == "" || req.Password == "" {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Check your request body.",
		})
		g.Abort()
		return
	}
	user := c.customerService.GetUser(req.Email, req.Password)
	if user.ID == 0 {
		g.JSON(http.StatusNotFound, gin.H{
			"error_message": "User not found!",
		})
		g.Abort()
		return
	}
	hashedRole, err := hashing.HashWord("customer") //role is hiding with hashing because if someone can reach and change
	if err != nil {
		fmt.Println("Error occurred:", err.Error())
	}
	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   user.ID,
		"email":    user.Email,
		"username": user.Username,
		"role":     hashedRole,
		"iat":      time.Now().Unix(),
		"iss":      "customer-service" + os.Getenv("ENV"),
		"exp": time.Now().Add(2 *
			time.Hour).Unix(),
	})
	token := jwtHelper.GenerateToken(jwtClaims, c.appConfig.JwtSettings.SecretKey)
	g.JSON(http.StatusOK, token)
}

func (c *ControllerCustomer) VerifyToken(g *gin.Context) {
	token := g.GetHeader("Authorization")
	decodedClaims := jwtHelper.VerifyToken(token, c.appConfig.JwtSettings.SecretKey, os.Getenv("ENV"))

	g.JSON(http.StatusOK, decodedClaims)

}
