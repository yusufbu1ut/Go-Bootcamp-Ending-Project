package admin

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/api/login"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/config"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/domain/admin"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/pkg/hashing"
	jwtHelper "github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/pkg/jwt"
)

type ControllerAdmin struct {
	appConfig    *config.Configuration
	adminService *admin.ServiceAdmin
}

// @BasePath /login

func NewAdminController(appConfig *config.Configuration, service *admin.ServiceAdmin) *ControllerAdmin {
	return &ControllerAdmin{
		appConfig:    appConfig,
		adminService: service,
	}
}

// Login godoc
// @Summary Admins login with email and password
// @Tags Login
// @Accept  json
// @Produce  json
// @Param login-request body login.RequestLogin true "Login process takes admin' email and password. Checks the inputs in database and returns JWT token."
// @Success 200 {object} map[string]string
// @Router /login/admin [post]
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
	//here hashing the role and using in middleware
	hashedRole, err := hashing.HashWord("admin")
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
