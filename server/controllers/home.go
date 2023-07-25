package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robinmuhia/arbitrageBetting/server/arbitrageBackend/models"
)

func Home(c *gin.Context) {
	user, _ := c.Get("user")
	
	c.JSON(http.StatusOK,gin.H{
		"id": user.(models.User).ID,
		"name": user.(models.User).Name,
		"image": user.(models.User).Image,
	})
}