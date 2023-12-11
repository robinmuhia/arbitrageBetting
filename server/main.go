package main

import (
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/robinmuhia/arbitrageBetting/server/arbitrageBackend/controllers"
	"github.com/robinmuhia/arbitrageBetting/server/arbitrageBackend/initializers"
	"github.com/robinmuhia/arbitrageBetting/server/arbitrageBackend/middleware"
)

// Initialize required dependencies such as the database connection and environment variables
func init(){
	initializers.LoadEnvironmentVariables()
	initializers.ConnectToDatabase()
	initializers.SyncDatabase()
}

// Configures the gin instance such as routers
func configureGin(wg *sync.WaitGroup){
	r := gin.Default()
	r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"PUT", "PATCH", "GET", "DELETE"},
        AllowHeaders:     []string{"Origin"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        AllowOriginFunc: func(origin string) bool {
            return origin == "http://localhost:3000"
        },
        MaxAge: 12 * time.Hour,
    }))
	r.POST("/api/v1/signup", controllers.SignUp)
	r.POST("/api/v1/login",controllers.Login)
	r.GET("/api/v1/logout",middleware.RequireAuth,controllers.Logout)
	r.GET("/api/v1/twoarbsbets",middleware.RequireAuth,controllers.GetTwoArbs)
	r.GET("/api/v1/threearbsbets",middleware.RequireAuth,controllers.GetThreeArbs)
	r.Run() 
	wg.Done()
}

// Main function that runs the entire backend service
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go configureGin(&wg)
	wg.Wait()
	controllers.LoadArbsInDB()
}