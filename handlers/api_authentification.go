package handlers

import (
	"net/http"
	"time"

	"github.com/dgb35/telemogus_backend/db"
	"github.com/dgb35/telemogus_backend/models"
	"github.com/dgb35/telemogus_backend/utils"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var credentials struct {
		ProxyId  string `json:"proxyUsername"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{ProxyIds: []string{credentials.ProxyId}, PasswordHash: credentials.Password}
	db.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func Login(c *gin.Context) {
	var credentials struct {
		ProxyId  string `json:"proxyId"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{ProxyIds: []string{credentials.ProxyId}}
	db.DB.Model(models.User{}).Find(&user)

	if user.Id == 0 || bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(credentials.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Incorrect username or password"})
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"proxyIds": user.ProxyIds,
		"masterId": user.Id,
		"exp":      expirationTime.Unix(),
	})

	tokenString, err := token.SignedString(utils.JWTKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func GetMasterUser(c *gin.Context) {
	var credentials struct {
		ProxyId string `json:"proxyUsername"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{ProxyIds: []string{credentials.ProxyId}}
	db.DB.Model(models.User{}).Find(&user)

	c.JSON(http.StatusOK, gin.H{"user": user})
}
