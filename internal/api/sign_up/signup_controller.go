package sign_up

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/config"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/domain/customer"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/pkg/hashing"
	jwtHelper "github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/pkg/jwt"
	"net/http"
	"os"
	"time"
)

type ControllerSignup struct {
	appConfig       *config.Configuration
	customerService *customer.ServiceCustomer
}

func NewSignupController(appConfig *config.Configuration, service *customer.ServiceCustomer) *ControllerSignup {
	return &ControllerSignup{
		appConfig:       appConfig,
		customerService: service,
	}
}

func (c *ControllerSignup) Signup(g *gin.Context) {
	var req customer.Customer
	err := g.ShouldBind(&req)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Check your request body.",
		})
		g.Abort()
		return
	}
	customer := customer.NewCustomer(req.Name, req.Username, req.Email, req.Password, req.PhoneNo, req.Address)
	err = c.customerService.Create(customer)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": err.Error(),
		})
		g.Abort()
		return
	}
	user := c.customerService.GetUser(req.Email, req.Password)
	hashedRole, err := hashing.HashWord("customer") //role is hiding with hashing because if someone can reach and change
	if err != nil {
		fmt.Println("Error occurred:", err.Error())
	}
	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   user.ID,
		"email":    user.Email,
		"username": user.Username,
		"role":     hashedRole, //hashing uygulanabilir
		"iat":      time.Now().Unix(),
		"iss":      "customer-service" + os.Getenv("ENV"),
		"exp": time.Now().Add(2 *
			time.Hour).Unix(),
	})
	token := jwtHelper.GenerateToken(jwtClaims, c.appConfig.JwtSettings.SecretKey)
	g.JSON(http.StatusCreated, token)
}
