package main

import (
	"github.com/gin-gonic/gin"
	"github.com/robinmuhia/arbitrageBetting/server/arbitrageBackend/controllers/auth"
	"github.com/robinmuhia/arbitrageBetting/server/arbitrageBackend/controllers/bets"
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
	r.POST("/api/v1/signup", auth.SignUp)
	r.POST("/api/v1/login",auth.Login)
	r.GET("/api/v1/logout",middleware.RequireAuth,auth.Logout)
	r.GET("/api/v1/bets",middleware.RequireAuth,bets.GetBets)
	r.Run() 
}