package main

import (
	"github.com/gin-gonic/gin"
	"github.com/robinmuhia/arbitrageBetting/server/arbitrageBackend/controllers"
	"github.com/robinmuhia/arbitrageBetting/server/arbitrageBackend/initializers"
	"github.com/robinmuhia/arbitrageBetting/server/arbitrageBackend/middleware"
)


func init(){
	initializers.LoadEnvironmentVariables()
	initializers.ConnectToDatabase()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/api/v1/signup", controllers.SignUp)
	r.POST("/api/v1/login",controllers.Login)
	r.GET("/api/v1/logout",middleware.RequireAuth,controllers.Logout)
	r.GET("/api/v1/bets",middleware.RequireAuth,controllers.LoadArbsInDB)
	r.Run() 
}