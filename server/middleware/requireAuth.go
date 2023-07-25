package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/robinmuhia/arbitrageBetting/server/arbitrageBackend/initializers"
	"github.com/robinmuhia/arbitrageBetting/server/arbitrageBackend/models"
)

func RequireAuth(c *gin.Context){
	jwtTokenString, err := c.Cookie("Authorization")
	if err != nil{
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	jwtToken, err := jwt.Parse(jwtTokenString,func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}	
		jwtSecret := os.Getenv("JWT_SECRET")
		return []byte(jwtSecret), nil
	})

	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":"Failed to parse jwt token",
		})
		return
	}

	if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64){
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		var user models.User
		initializers.DB.First(&user,claims["sub"])
		if user.ID == 0{
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("user",user)
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}