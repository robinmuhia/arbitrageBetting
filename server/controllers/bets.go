package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/robinmuhia/arbitrageBetting/server/arbitrageBackend/initializers"
	"github.com/robinmuhia/arbitrageBetting/server/arbitrageBackend/models"
)

func getSports() ([]models.Sport, error) {
	apiKey := os.Getenv("ODDS_API_KEY")
	if apiKey == "" {
		return nil,errors.New("missing environment variable")
	}
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

func getOdds() ([]models.Odds,error) {
	apiKey := os.Getenv("ODDS_API_KEY")
	if apiKey == "" {
		return nil,errors.New("missing environment variable")
	}

	region := "uk"
	markets := "h2h"
	oddsFormat := "decimal"
	dateFormat := "iso"

	sports, err := getSports()
	if err != nil {
		return nil,errors.New(err.Error())
	}
	var wg sync.WaitGroup
	var mu sync.Mutex
	var allOdds []models.Odds

	httpClient := &http.Client{}

	for _, sport := range sports {
		if sport.Active {
			wg.Add(1)
			go func(sport models.Sport) {
				defer wg.Done()

				params := url.Values{}
				params.Add("api_key", apiKey)
				params.Add("regions", region)
				params.Add("markets", markets)
				params.Add("oddsFormat", oddsFormat)
				params.Add("dateFormat", dateFormat)

				url := fmt.Sprintf("https://api.the-odds-api.com/v4/sports/%s/odds?%s", sport.Key, params.Encode())

				response, err := httpClient.Get(url)
				if err != nil {
					log.Printf("Failed to fetch odds for %s: %s\n", sport.Title, err)
					return
				}
				defer response.Body.Close()

				if response.StatusCode != http.StatusOK {
					log.Printf("Failed to get odds data for %s: %s\n", sport.Title, response.Status)
					return
				}

				var odds []models.Odds
				if err := json.NewDecoder(response.Body).Decode(&odds); err != nil {
					log.Printf("Failed to decode odds data for %s: %s\n", sport.Title, err)
					return
				}

				mu.Lock()
				allOdds = append(allOdds, odds...)
				mu.Unlock()
			}(sport)
		}
	}
	wg.Wait()
	return allOdds,nil
}


func getArbs()([]models.ThreeOddsBet,[]models.TwoOddsBet,error){
	odds,err := getOdds()

	if err != nil{
		return nil,nil,errors.New(err.Error())
	}

	var ThreeOddsArbs []models.ThreeOddsBet
	var TwoOddsArbs []models.TwoOddsBet

	for _, odd := range odds{
		for i := 0; i < len(odd.Bookmakers);i++{
			for j := 0; j < len(odd.Bookmakers); j++{
				if len(odd.Bookmakers[i].Markets[0].Outcomes) == 3 && odd.Bookmakers[i].Key == "h2h"{
					for k := 0; k < len(odd.Bookmakers); k++{
						homeOdd := odd.Bookmakers[i].Markets[0].Outcomes[0].Price
						awayOdd := odd.Bookmakers[j].Markets[0].Outcomes[1].Price
						drawOdd := odd.Bookmakers[k].Markets[0].Outcomes[2].Price
						arb := ((1/homeOdd)+(1/awayOdd)+(1/drawOdd))
						if arb < float64(1) {
							threewayArb := models.ThreeOddsBet{
								Title: fmt.Sprintf("%s - %s",odd.HomeTeam,odd.AwayTeam) ,
								Home: odd.Bookmakers[i].Title,
								Away: odd.Bookmakers[j].Title,
								Draw: odd.Bookmakers[k].Title,
								HomeOdds: homeOdd,
								AwayOdds: awayOdd,
								DrawOdds: drawOdd,
								GameType: odd.SportKey,
								League: odd.SportTitle,
								BookmarkerRegion: "uk",
								GameTime: odd.CommenceTime,								
							}
							ThreeOddsArbs = append(ThreeOddsArbs, threewayArb)
						}
					}
				} else if len(odd.Bookmakers[i].Markets[0].Outcomes) == 2 && odd.Bookmakers[i].Key == "h2h"{
					for k := 0; k < len(odd.Bookmakers); k++{
						homeOdd := odd.Bookmakers[i].Markets[0].Outcomes[0].Price
						awayOdd := odd.Bookmakers[j].Markets[0].Outcomes[1].Price
						arb := ((1/homeOdd)+(1/awayOdd))
						if arb < float64(1) {
							twowayArb := models.TwoOddsBet{
								Title: fmt.Sprintf("%s - %s",odd.HomeTeam,odd.AwayTeam) ,
								Home: odd.Bookmakers[i].Title,
								Away: odd.Bookmakers[j].Title,
								HomeOdds: homeOdd,
								AwayOdds: awayOdd,
								GameType: odd.SportKey,
								League: odd.SportTitle,
								BookmarkerRegion: "uk",
								GameTime: odd.CommenceTime,								
							}
							TwoOddsArbs = append(TwoOddsArbs, twowayArb)							
						}
					}
				}
			}
		}
	}
	return ThreeOddsArbs,TwoOddsArbs,nil
}
	

func LoadArbsInDB(c *gin.Context){
	threeArbs, twoArbs, err := getArbs()
	if err != nil{
		c.JSON(http.StatusMethodNotAllowed,gin.H{})
		return
	}
	if len(threeArbs) > 0{
		initializers.DB.Migrator().DropTable(&models.ThreeOddsBet{})
		initializers.DB.Migrator().CreateTable(&models.ThreeOddsBet{})
		for _,arbs := range threeArbs{
			initializers.DB.Create(&arbs)
		}
	}
	if len(twoArbs) > 0{
		initializers.DB.Migrator().DropTable(&models.TwoOddsBet{})
		initializers.DB.Migrator().CreateTable(&models.TwoOddsBet{})
		for _,arbs := range twoArbs{
			initializers.DB.Create(&arbs)
		}
	}	
	c.JSON(http.StatusOK,gin.H{})
}


// func GetArbs(c *gin.Context){
// 	_, err := getSports()
// 	if err != nil{
// 		return
// 	}
// }