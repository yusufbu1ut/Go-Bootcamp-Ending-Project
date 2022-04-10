package customer

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/api/login"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/config"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/domain/customer"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/pkg/hashing"
	jwtHelper "github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/pkg/jwt"
	"log"
	"net/http"
	"os"
	"time"
)

type ControllerCustomer struct {
	appConfig       *config.Configuration
	customerService *customer.ServiceCustomer
}

// @BasePath /login

func NewCustomerController(appConfig *config.Configuration, service *customer.ServiceCustomer) *ControllerCustomer {
	return &ControllerCustomer{
		appConfig:       appConfig,
		customerService: service,
	}
}

// Login godoc
// @Summary Customers login with email and password
// @Tags Login
// @Accept  json
// @Produce  json
// @Param login-request body login.RequestLogin true "Login process takes customer' email and password. Checks the inputs in database and returns JWT token."
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /login/customer [post]
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
	//here hashing the role and using it in middleware
	hashedRole, err := hashing.HashWord("customer") //role is hiding with hashing because if someone can reach and change
	if err != nil {
		log.Println("Error occurred:", err.Error())
		g.JSON(http.StatusInternalServerError, gin.H{
			"error_message": err.Error(),
		})
		g.Abort()
		return
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
