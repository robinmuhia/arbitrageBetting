package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/robinmuhia/arbitrageBetting/server/arbitrageBackend/controllers"
	"github.com/robinmuhia/arbitrageBetting/server/arbitrageBackend/initializers"
	"github.com/robinmuhia/arbitrageBetting/server/arbitrageBackend/middleware"
	"github.com/robinmuhia/arbitrageBetting/server/arbitrageBackend/models"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)


func setUpRouter() *gin.Engine{
    gin.SetMode(gin.TestMode)
    router := gin.Default()
    return router
}

func TestSignUpUserExists(t *testing.T) {
    // Create a user with the same email in the database
    initializers.DB.Migrator().DropTable(&models.User{})
    initializers.DB.Migrator().CreateTable(&models.User{})
    user := models.User{
        Name:             "TestUser",
        Email:            "test@example.com",
        Password:         "Password123",
        BookmarkerRegion: "uk",
        SubscriptionPaid: true,
    }
    initializers.DB.Create(&user)

    r := setUpRouter()
    r.POST("/api/v1/signup", controllers.SignUp)

    newUser := models.User{
        Name:     "NewUser",
        Email:    "test@example.com", // Use the same email as an existing user
        Password: "NewPassword123",
    }
    requestBody, _ := json.Marshal(newUser)
    req, _ := http.NewRequest("POST", "/api/v1/signup", bytes.NewBuffer(requestBody))
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusConflict, w.Code)
}

func TestSignUpIncorrectDataTypes(t *testing.T) {
    r := setUpRouter()
    r.POST("/api/v1/signup", controllers.SignUp)

    // Send invalid JSON data (missing required fields)
    requestBody := `{"InvalidKey": "invalidValue"}`

    req, _ := http.NewRequest("POST", "/api/v1/signup", bytes.NewBufferString(requestBody))
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSuccessfulSignUp(t *testing.T) {
    initializers.DB.Migrator().DropTable(&models.User{})
    initializers.DB.Migrator().CreateTable(&models.User{})
    mockResponse := `{"id":1,"name":"TestUser"}`
    r := setUpRouter()
    r.POST("/api/v1/signup", controllers.SignUp)
    user := models.User{
        Name: "TestUser",
        Email: "test@example.com",
        Password: "Password123",
    }
    requestBody,_ := json.Marshal(user)
    req, _ := http.NewRequest("POST", "/api/v1/signup", bytes.NewBuffer(requestBody))
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    responseData, _ := io.ReadAll(w.Body)
    assert.Equal(t, mockResponse, string(responseData))
    assert.Equal(t, http.StatusOK, w.Code)
}

func TestLogin_Success(t *testing.T) {
    // Prepare a user with known credentials in the database
    initializers.DB.Migrator().DropTable(&models.User{})
    initializers.DB.Migrator().CreateTable(&models.User{})
    user := models.User{
        Name:     "TestUser",
        Email:    "test@example.com",
        Password: "Password123",
    }
    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
    user.Password = string(hashedPassword)
    initializers.DB.Create(&user)

    // Create a request with valid credentials
    requestBody := map[string]string{
        "Email":    "test@example.com",
        "Password": "Password123",
    }
    reqBody, _ := json.Marshal(requestBody)
    req, _ := http.NewRequest("POST", "/api/v1/login", bytes.NewBuffer(reqBody))
    r := setUpRouter()
    r.POST("/api/v1/login", controllers.Login)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    // Check the response
    assert.Equal(t, http.StatusOK, w.Code)
    responseData, _ := io.ReadAll(w.Body)
    assert.Contains(t, string(responseData), "id")
    assert.Contains(t, string(responseData), "name")
}

func TestLogin_IncorrectCredentials(t *testing.T) {
    // Prepare an empty database
    initializers.DB.Migrator().DropTable(&models.User{})
    initializers.DB.Migrator().CreateTable(&models.User{})

    // Prepare a user with known credentials in the database
    initializers.DB.Migrator().DropTable(&models.User{})
    initializers.DB.Migrator().CreateTable(&models.User{})
    user := models.User{
        Name:     "TestUser",
        Email:    "test@example.com",
        Password: "Password123",
    }
    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
    user.Password = string(hashedPassword)
    initializers.DB.Create(&user)

    // Create a request with incorrect credentials
    requestBody := map[string]string{
        "Email":    "test@example.com",
        "Password": "WrongPassword",
    }
    reqBody, _ := json.Marshal(requestBody)
    req, _ := http.NewRequest("POST", "/api/v1/login", bytes.NewBuffer(reqBody))
    r := setUpRouter()
    r.POST("/api/v1/login", controllers.Login)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    // Check the response
    assert.Equal(t, http.StatusForbidden, w.Code)
    responseData, _ := io.ReadAll(w.Body)
    assert.Contains(t, string(responseData), "error")
    assert.Contains(t, string(responseData), "Invalid email or password")
}

func TestLogin_NoUser(t *testing.T) {
    // Prepare an empty database
    initializers.DB.Migrator().DropTable(&models.User{})
    initializers.DB.Migrator().CreateTable(&models.User{})

    // Create a request with incorrect credentials
    requestBody := map[string]string{
        "Email":    "test@example.com",
        "Password": "WrongPassword",
    }
    reqBody, _ := json.Marshal(requestBody)
    req, _ := http.NewRequest("POST", "/api/v1/login", bytes.NewBuffer(reqBody))
    r := setUpRouter()
    r.POST("/api/v1/login", controllers.Login)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    // Check the response
    assert.Equal(t, http.StatusForbidden, w.Code)
    responseData, _ := io.ReadAll(w.Body)
    assert.Contains(t, string(responseData), "error")
    assert.Contains(t, string(responseData), "Invalid email or password")
}

func TestLogout(t *testing.T){
    // Prepare a user with known credentials in the database
    initializers.DB.Migrator().DropTable(&models.User{})
    initializers.DB.Migrator().CreateTable(&models.User{})
    user := models.User{
        Name:     "TestUser",
        Email:    "test@example.com",
        Password: "Password123",
    }
    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
    user.Password = string(hashedPassword)
    initializers.DB.Create(&user)


    jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"sub":1,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	jwtSecret := os.Getenv("JWT_SECRET")
    jwtTokenString, err := jwtToken.SignedString([]byte(jwtSecret))
    if err != nil {
        t.Fatal(err)
    }

    // Create a request with the JWT token in the Cookie
    req, _ := http.NewRequest("GET", "/api/v1/logout", nil)
    req.AddCookie(&http.Cookie{Name: "Authorization", Value: jwtTokenString})
    r := setUpRouter()

    // Use the RequireAuth middleware and the Logout controller
    r.GET("/api/v1/logout", middleware.RequireAuth, controllers.Logout)

    // Send the request and record the response
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    // Check the response
    assert.Equal(t, http.StatusOK, w.Code)
    
    // Ensure that the "Authorization" cookie is cleared
    cookies := w.Result().Cookies()
    assert.Len(t, cookies, 1) // Only one cookie should remain
    assert.Equal(t, "Authorization", cookies[0].Name)
    assert.Empty(t, cookies[0].Value) // The cookie value should be empt

}

func TestGetTwoArbs(t *testing.T) {
    r := setUpRouter()

    // Create some example data for the test
twoArbs := []models.TwoOddsBet{
    {
        Title:           "Bet 1",
        Home:            "Team A",
        Away:            "Team B",
        HomeOdds:        2.0,
        AwayOdds:        2.5,
        GameType:        "Football",
        League:          "Premier League",
        Profit:          0.1,
        BookmarkerRegion: "UK",
        GameTime:        "2023-10-01 15:00:00",
    },
    {
        Title:           "Bet 2",
        Home:            "Team X",
        Away:            "Team Y",
        HomeOdds:        1.8,
        AwayOdds:        2.2,
        GameType:        "Basketball",
        League:          "NBA",
        Profit:          0.2,
        BookmarkerRegion: "US",
        GameTime:        "2023-10-02 18:30:00",
    },
    {
        Title:           "Bet 3",
        Home:            "Team 1",
        Away:            "Team 2",
        HomeOdds:        1.5,
        AwayOdds:        2.0,
        GameType:        "Tennis",
        League:          "Wimbledon",
        Profit:          0.05,
        BookmarkerRegion: "EU",
        GameTime:        "2023-10-03 10:00:00",
    },
    {
        Title:           "Bet 4",
        Home:            "Team Red",
        Away:            "Team Blue",
        HomeOdds:        2.2,
        AwayOdds:        2.8,
        GameType:        "Soccer",
        League:          "La Liga",
        Profit:          0.15,
        BookmarkerRegion: "ES",
        GameTime:        "2023-10-04 20:45:00",
    },
    {
        Title:           "Bet 5",
        Home:            "Team Alpha",
        Away:            "Team Beta",
        HomeOdds:        1.6,
        AwayOdds:        1.9,
        GameType:        "Hockey",
        League:          "NHL",
        Profit:          0.12,
        BookmarkerRegion: "CA",
        GameTime:        "2023-10-05 19:15:00",
    },
    }
    initializers.DB.Migrator().DropTable(&models.TwoOddsBet{})
    initializers.DB.Migrator().CreateTable(&models.TwoOddsBet{})
    for _,arbs := range twoArbs{
        initializers.DB.Create(&arbs)
    }

    // Simulate a request to the "/api/v1/two-arbs" endpoint
    jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"sub":1,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	jwtSecret := os.Getenv("JWT_SECRET")
    jwtTokenString, err := jwtToken.SignedString([]byte(jwtSecret))
    if err != nil {
        t.Fatal(err)
    }

    // Use the RequireAuth middleware and the Logout controller
    r.GET("/api/v1/twoarbsbets",middleware.RequireAuth,controllers.GetTwoArbs)

    // Create a request with the JWT token in the Cookie   
    req, _ := http.NewRequest("GET", "/api/v1/twoarbsbets", nil)
    req.AddCookie(&http.Cookie{Name: "Authorization", Value: jwtTokenString})
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    // Check the response status code
    assert.Equal(t, http.StatusOK, w.Code)

    // Parse the JSON response
    var response struct {
        TwoArbs []models.TwoOddsBet `json:"twoArbs"`
    }

    err = json.NewDecoder(w.Body).Decode(&response)
    if err != nil {
        t.Fatal(err)
    }

    // Verify that the response contains the expected data
    assert.Len(t, response.TwoArbs, len(twoArbs))
}

func TestGetUnauthorizedTwoArbs(t *testing.T) {
    r := setUpRouter()

    // Create some example data for the test
twoArbs := []models.TwoOddsBet{
    {
        Title:           "Bet 1",
        Home:            "Team A",
        Away:            "Team B",
        HomeOdds:        2.0,
        AwayOdds:        2.5,
        GameType:        "Football",
        League:          "Premier League",
        Profit:          0.1,
        BookmarkerRegion: "UK",
        GameTime:        "2023-10-01 15:00:00",
    },
    {
        Title:           "Bet 2",
        Home:            "Team X",
        Away:            "Team Y",
        HomeOdds:        1.8,
        AwayOdds:        2.2,
        GameType:        "Basketball",
        League:          "NBA",
        Profit:          0.2,
        BookmarkerRegion: "US",
        GameTime:        "2023-10-02 18:30:00",
    },
    {
        Title:           "Bet 3",
        Home:            "Team 1",
        Away:            "Team 2",
        HomeOdds:        1.5,
        AwayOdds:        2.0,
        GameType:        "Tennis",
        League:          "Wimbledon",
        Profit:          0.05,
        BookmarkerRegion: "EU",
        GameTime:        "2023-10-03 10:00:00",
    },
    {
        Title:           "Bet 4",
        Home:            "Team Red",
        Away:            "Team Blue",
        HomeOdds:        2.2,
        AwayOdds:        2.8,
        GameType:        "Soccer",
        League:          "La Liga",
        Profit:          0.15,
        BookmarkerRegion: "ES",
        GameTime:        "2023-10-04 20:45:00",
    },
    {
        Title:           "Bet 5",
        Home:            "Team Alpha",
        Away:            "Team Beta",
        HomeOdds:        1.6,
        AwayOdds:        1.9,
        GameType:        "Hockey",
        League:          "NHL",
        Profit:          0.12,
        BookmarkerRegion: "CA",
        GameTime:        "2023-10-05 19:15:00",
    },
    }
    initializers.DB.Migrator().DropTable(&models.TwoOddsBet{})
    initializers.DB.Migrator().CreateTable(&models.TwoOddsBet{})
    for _,arbs := range twoArbs{
        initializers.DB.Create(&arbs)
    }

    // Use the RequireAuth middleware and the Logout controller
    r.GET("/api/v1/twoarbsbets",middleware.RequireAuth,controllers.GetTwoArbs)

    // Create a request with the JWT token in the Cookie   
    req, _ := http.NewRequest("GET", "/api/v1/twoarbsbets", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    // Check the response status code
    assert.Equal(t, http.StatusUnauthorized, w.Code)
    
    // Parse the JSON response
    var response struct {
        TwoArbs []models.ThreeOddsBet `json:"threeArbs"`
    }

    err := json.NewDecoder(w.Body).Decode(&response)
    if err != nil {
        t.Fatal(err)
    }

    // Verify that the response contains the expected data
    assert.Len(t, response.TwoArbs, 0)
}


func TestGetThreeArbs(t *testing.T) {
    r := setUpRouter()

    // Create some example data for the test
threeArbs := []models.ThreeOddsBet{
    {
        Title:           "Bet 1",
        Home:            "Team A",
        Away:            "Team B",
        Draw:            "Team C", 
        HomeOdds:        2.0,
        DrawOdds:        1.5,
        AwayOdds:        2.5,
        GameType:        "Football",
        League:          "Premier League",
        Profit:          0.1,
        BookmarkerRegion: "UK",
        GameTime:        "2023-10-01 15:00:00",
    },
    {
        Title:           "Bet 2",
        Home:            "Team X",
        Away:            "Team Y",
        Draw:            "Team Z",
        HomeOdds:        1.8,
        DrawOdds:        3.1, 
        AwayOdds:        2.2,
        GameType:        "Basketball",
        League:          "NBA",
        Profit:          0.2,
        BookmarkerRegion: "US",
        GameTime:        "2023-10-02 18:30:00",
    },
    }
    initializers.DB.Migrator().DropTable(&models.ThreeOddsBet{})
    initializers.DB.Migrator().CreateTable(&models.ThreeOddsBet{})
    for _,arbs := range threeArbs{
        initializers.DB.Create(&arbs)
    }

    // Simulate a request to the "/api/v1/two-arbs" endpoint
    jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"sub":1,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	jwtSecret := os.Getenv("JWT_SECRET")
    jwtTokenString, err := jwtToken.SignedString([]byte(jwtSecret))
    if err != nil {
        t.Fatal(err)
    }

    // Use the RequireAuth middleware and the Logout controller
    r.GET("/api/v1/threearbsbets",middleware.RequireAuth,controllers.GetThreeArbs)

    // Create a request with the JWT token in the Cookie   
    req, _ := http.NewRequest("GET", "/api/v1/threearbsbets", nil)
    req.AddCookie(&http.Cookie{Name: "Authorization", Value: jwtTokenString})
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    // Check the response status code
    assert.Equal(t, http.StatusOK, w.Code)

    // Parse the JSON response
    var response struct {
        TwoArbs []models.ThreeOddsBet `json:"threeArbs"`
    }

    err = json.NewDecoder(w.Body).Decode(&response)
    if err != nil {
        t.Fatal(err)
    }

    // Verify that the response contains the expected data
    assert.Len(t, response.TwoArbs, len(threeArbs))
    
}


func TestUnaouthorizedGetThreeArbs(t *testing.T) {
    r := setUpRouter()

    // Create some example data for the test
threeArbs := []models.ThreeOddsBet{
    {
        Title:           "Bet 1",
        Home:            "Team A",
        Away:            "Team B",
        Draw:            "Team C", 
        HomeOdds:        2.0,
        DrawOdds:        1.5,
        AwayOdds:        2.5,
        GameType:        "Football",
        League:          "Premier League",
        Profit:          0.1,
        BookmarkerRegion: "UK",
        GameTime:        "2023-10-01 15:00:00",
    },
    {
        Title:           "Bet 2",
        Home:            "Team X",
        Away:            "Team Y",
        Draw:            "Team Z",
        HomeOdds:        1.8,
        DrawOdds:        3.1, 
        AwayOdds:        2.2,
        GameType:        "Basketball",
        League:          "NBA",
        Profit:          0.2,
        BookmarkerRegion: "US",
        GameTime:        "2023-10-02 18:30:00",
    },
    }
    initializers.DB.Migrator().DropTable(&models.ThreeOddsBet{})
    initializers.DB.Migrator().CreateTable(&models.ThreeOddsBet{})
    for _,arbs := range threeArbs{
        initializers.DB.Create(&arbs)
    }

    r.GET("/api/v1/threearbsbets",middleware.RequireAuth,controllers.GetThreeArbs)

    // Create a request with the JWT token in the Cookie   
    req, _ := http.NewRequest("GET", "/api/v1/threearbsbets", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    // Check the response status code
    assert.Equal(t, http.StatusUnauthorized, w.Code)

    // Parse the JSON response
    var response struct {
        TwoArbs []models.ThreeOddsBet `json:"threeArbs"`
    }

    err := json.NewDecoder(w.Body).Decode(&response)
    if err != nil {
        t.Fatal(err)
    }

    // Verify that the response contains the expected data
    assert.Len(t, response.TwoArbs, 0)
}
