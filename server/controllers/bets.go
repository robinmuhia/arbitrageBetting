package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/robinmuhia/arbitrageBetting/server/arbitrageBackend/models"
)

func GetArbs(c *gin.Context){
	sports, err := getSports()
	if err != nil{
		return
	}
	fmt.Println(sports)
}

func getSports() ([]models.Sport, error) {
	apiKey := os.Getenv("ODDS_API_KEY")
	params := url.Values{}
	params.Add("api_key", apiKey)
	url := "https://api.the-odds-api.com/v4/sports?" + params.Encode()

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get sports data: %s", response.Status)
	}

	var sports []models.Sport
	if err := json.NewDecoder(response.Body).Decode(&sports); err != nil {
		return nil, err
	}
	return sports, nil
}

// func GetOdds() [][]models.Odds {
// 	apiKey := os.Getenv("ODDS_API_KEY")
// 	if apiKey == "" {
// 		log.Println("Missing ODDS_API_KEY environment variable.")
// 		return nil
// 	}

// 	region := "uk"
// 	markets := "h2h"
// 	oddsFormat := "decimal"
// 	dateFormat := "iso"

// 	sports, err := getSports()
// 	if err != nil {
// 		log.Printf("Failed to fetch sports data: %s\n", err)
// 		return nil
// 	}

// 	var wg sync.WaitGroup
// 	var mu sync.Mutex
// 	var allOdds [][]models.Odds

// 	httpClient := &http.Client{}

// 	for _, sport := range sports {
// 		if sport.Active {
// 			wg.Add(1)
// 			go func(sport models.Sport) {
// 				defer wg.Done()

// 				params := url.Values{}
// 				params.Add("api_key", apiKey)
// 				params.Add("regions", region)
// 				params.Add("markets", markets)
// 				params.Add("oddsFormat", oddsFormat)
// 				params.Add("dateFormat", dateFormat)

// 				url := fmt.Sprintf("https://api.the-odds-api.com/v4/sports/%s/odds?%s", sport.Key, params.Encode())

// 				response, err := httpClient.Get(url)
// 				if err != nil {
// 					log.Printf("Failed to fetch odds for %s: %s\n", sport.Title, err)
// 					return
// 				}
// 				defer response.Body.Close()

// 				if response.StatusCode != http.StatusOK {
// 					log.Printf("Failed to get odds data for %s: %s\n", sport.Title, response.Status)
// 					return
// 				}

// 				var odds []models.Odds
// 				if err := json.NewDecoder(response.Body).Decode(&odds); err != nil {
// 					log.Printf("Failed to decode odds data for %s: %s\n", sport.Title, err)
// 					return
// 				}

// 				mu.Lock()
// 				allOdds = append(allOdds, odds)
// 				mu.Unlock()
// 			}(sport)
// 		}
// 	}
// 	wg.Wait()
// 	fmt.Println(allOdds)
// 	return allOdds
// }


// func getArbs(){
// 	odds := GetOdds()

// 	var ThreeOddsArbs []models.ThreeOddsBet
// 	var TwoOddsArbs []models.TwoOddsBet

// 	for _, odd := range odds{
// 		for i := 0; i < len(odd); i++{
// 			if len(odd[i].Bookmakers[0].Markets[0].Outcomes) == 3{
				
// 			}
// 		}
// 	}
// }