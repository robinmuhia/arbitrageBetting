package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/robinmuhia/arbitrageBetting/server/arbitrageBackend/initializers"
	"github.com/robinmuhia/arbitrageBetting/server/arbitrageBackend/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context){
	var body struct{
		Email string `binding:"required"`
		Password string `binding:"required"`
		Name string `binding:"required"`
	}

	if err := c.Bind(&body); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Failed to get all values or incorrect data-types were sent",
		})
		return
	}

	hashedPassword,err := bcrypt.GenerateFromPassword([]byte(body.Password),10)

	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Failed to hash password",
		})
		return
	}
	user := models.User{
		Name: body.Name,
		Password: string(hashedPassword),
		Email: body.Email,
		BookmarkerRegion: "uk",
		SubscriptionPaid: true,
	}

	result := initializers.DB.Create(&user)
	if result.Error != nil{
		c.JSON(http.StatusConflict,gin.H{
			"error":"User already exists",
		})
		return
	}
	c.JSON(http.StatusCreated,gin.H{
		"id": user.ID,
		"name": user.Name,
	})
}


func Login(c *gin.Context){
	var body struct{
		Email string
		Password string
	}
	
	if err := c.Bind(&body); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Failed to get all values from body",
		})
		return
	}
	var user models.User
	initializers.DB.First(&user,"email = ?",body.Email)
	if user.ID == 0 {
		c.JSON(http.StatusForbidden,gin.H{
			"error":"Invalid email or passowrd",
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(body.Password)); err != nil{
		c.JSON(http.StatusForbidden,gin.H{
			"error":"Invalid email or password",
		})
		return
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"sub":user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	jwtSecret := os.Getenv("JWT_SECRET")
	jwtTokenString, err := jwtToken.SignedString([]byte(jwtSecret))

	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":"Failed to create jwt token",
		})
		return
	}
	c.SetSameSite(http.SameSiteDefaultMode)
	c.SetCookie("Authorization",jwtTokenString,3600*24*30,"","",true,true)

	c.JSON(http.StatusOK,gin.H{
		"id": user.ID,
		"name": user.Name,
	})
}

func Logout(c *gin.Context) {
    c.SetCookie("Authorization", "", -1, "", "", true, true)
	c.JSON(http.StatusOK,gin.H{
		"message":"User successfully logged out",
	})
}