package sign_up

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/config"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/domain/customer"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/pkg/hashing"
	jwtHelper "github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/pkg/jwt"
	"log"
	"net/http"
	"os"
	"time"
)

type ControllerSignup struct {
	appConfig       *config.Configuration
	customerService *customer.ServiceCustomer
}

// @BasePath /signup

func NewSignupController(appConfig *config.Configuration, service *customer.ServiceCustomer) *ControllerSignup {
	return &ControllerSignup{
		appConfig:       appConfig,
		customerService: service,
	}
}

// Signup godoc
// @Summary Customers can sign up with username,email and password(Needed fields)
// @Tags SignUp
// @Accept  json
// @Produce  json
// @Param customer body RequestCustomer true "Sing-up takes  username, email and password necessarily. Other fields not necessary but customers can add them too. Checks the customer in database adds it and returns JWT token."
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /signup [post]
//Signup works with request body it should contain username,email and password
func (c *ControllerSignup) Signup(g *gin.Context) {
	var req RequestCustomer
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
		"role":     hashedRole, //hashing uygulanabilir
		"iat":      time.Now().Unix(),
		"iss":      "customer-service" + os.Getenv("ENV"),
		"exp": time.Now().Add(2 *
			time.Hour).Unix(),
	})
	token := jwtHelper.GenerateToken(jwtClaims, c.appConfig.JwtSettings.SecretKey)
	g.JSON(http.StatusCreated, token)
}
