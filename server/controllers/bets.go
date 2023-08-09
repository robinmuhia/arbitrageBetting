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
	"time"

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


func getOdds() ([]models.Odds, error) {
	apiKey := os.Getenv("ODDS_API_KEY")
	if apiKey == "" {
		return nil, errors.New("missing environment variable")
	}

	region := "uk"
	markets := "h2h"
	oddsFormat := "decimal"
	dateFormat := "iso"

	sports, err := getSports()
	if err != nil {
		return nil, errors.New(err.Error())
	}
	var wg sync.WaitGroup
	var mu sync.Mutex
	var allOdds []models.Odds

	httpClient := &http.Client{}
	ticker := time.NewTicker(1 * time.Second) // Set the wait duration to 1 second

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

			// Wait for the specified duration before launching the next goroutine
			<-ticker.C
		}
	}
	wg.Wait()
	ticker.Stop() // Stop the ticker before returning

	return allOdds, nil
}


func getArbs() ([]models.ThreeOddsBet, []models.TwoOddsBet, error) {
	odds, err := getOdds()
	if err != nil {
		return nil, nil, err
	}

	var ThreeOddsArbs []models.ThreeOddsBet
	var TwoOddsArbs []models.TwoOddsBet

	// Create channels to receive arbitrage results
	threeOddsCh := make(chan models.ThreeOddsBet)
	twoOddsCh := make(chan models.TwoOddsBet)

	// Create a wait group to ensure all goroutines finish before returning
	var wg sync.WaitGroup
	var once sync.Once

	for _, odd := range odds {
		wg.Add(1)
		go processOdd(odd, threeOddsCh, twoOddsCh, &wg)
	}

	// Close the channels once all goroutines finish processing
	go func() {
		wg.Wait()
		once.Do(func() {
			close(threeOddsCh)
			close(twoOddsCh)
		})
	}()

	// Collect the results from channels
	for {
        select {
        case arb, ok := <-threeOddsCh:
            if !ok {
                threeOddsCh = nil // Set to nil to exit the loop when both channels are closed
            } else {
                ThreeOddsArbs = append(ThreeOddsArbs, arb)
            }
        case arb, ok := <-twoOddsCh:
            if !ok {
                twoOddsCh = nil // Set to nil to exit the loop when both channels are closed
            } else {
                TwoOddsArbs = append(TwoOddsArbs, arb)
            }
        }
        // Exit the loop when both channels are closed
        if threeOddsCh == nil && twoOddsCh == nil {
            break
        }
    }

    return ThreeOddsArbs, TwoOddsArbs, nil

}


func processOdd(odd models.Odds, threeOddsCh chan<- models.ThreeOddsBet, twoOddsCh chan<- models.TwoOddsBet, wg *sync.WaitGroup) {
	defer wg.Done()

	if len(odd.Bookmakers) < 2 {
		return // Skip if there are not enough bookmakers for comparison
	}

	for i := 0; i < len(odd.Bookmakers); i++ {
		if len(odd.Bookmakers[i].Markets) == 0 || odd.Bookmakers[i].Markets[0].Key != "h2h" {
			return // Skip if the market is not 'h2h'
			}
		for j := 0; j < len(odd.Bookmakers); j++ {
			if len(odd.Bookmakers[j].Markets) == 0 || odd.Bookmakers[j].Markets[0].Key != "h2h" {
				return // Skip if the market is not 'h2h'
				}
			if len(odd.Bookmakers[i].Markets[0].Outcomes) == 2 && len(odd.Bookmakers[j].Markets[0].Outcomes) == 2 {
				homeOdd := odd.Bookmakers[i].Markets[0].Outcomes[0].Price
				awayOdd := odd.Bookmakers[j].Markets[0].Outcomes[1].Price
				arb := (1 / homeOdd) + (1 / awayOdd)
				if arb < 1.0 {
					profit := (1 - arb) * 100
					twowayArb := models.TwoOddsBet{
						Title:            fmt.Sprintf("%s - %s", odd.HomeTeam, odd.AwayTeam),
						Home:             odd.Bookmakers[i].Title,
						Away:             odd.Bookmakers[j].Title,
						HomeOdds:         homeOdd,
						AwayOdds:         awayOdd,
						GameType:         odd.SportKey,
						League:           odd.SportTitle,
						Profit:           profit,
						BookmarkerRegion: "uk",
						GameTime:         odd.CommenceTime,
						}
						twoOddsCh <- twowayArb
					}
				} else if len(odd.Bookmakers[i].Markets[0].Outcomes) == 3 && len(odd.Bookmakers[j].Markets[0].Outcomes) == 3 {
					for k := 0; k < len(odd.Bookmakers); k++ {
						if len(odd.Bookmakers[k].Markets) == 0 || odd.Bookmakers[k].Markets[0].Key != "h2h" {
							return // Skip if the market is not 'h2h'
							}
						if len(odd.Bookmakers[k].Markets[0].Outcomes) == 3 {
							homeOdd := odd.Bookmakers[i].Markets[0].Outcomes[0].Price
							awayOdd := odd.Bookmakers[j].Markets[0].Outcomes[1].Price
							drawOdd := odd.Bookmakers[k].Markets[0].Outcomes[2].Price
							arb := (1 / homeOdd) + (1 / awayOdd) + (1 / drawOdd)
							if arb < 1.0 {
								profit := (1 - arb) * 100
								threewayArb := models.ThreeOddsBet {
									Title:            fmt.Sprintf("%s - %s", odd.HomeTeam, odd.AwayTeam),
									Home:             odd.Bookmakers[i].Title,
									Away:             odd.Bookmakers[j].Title,
									Draw:             odd.Bookmakers[k].Title,
									HomeOdds:         homeOdd,
									AwayOdds:         awayOdd,
									DrawOdds:         drawOdd,
									GameType:         odd.SportKey,
									League:           odd.SportTitle,
									BookmarkerRegion: "uk",
									Profit:           profit,
									GameTime:         odd.CommenceTime,
								}
								threeOddsCh <- threewayArb
						}
					}
				}
			}
		}
	}
} 
	

func LoadArbsInDB(){
	ticker := time.NewTicker(time.Hour*24)
	for ; ; <-ticker.C	{
		threeArbs, twoArbs, err := getArbs()
		if err != nil{
			continue
		}
		if len(threeArbs) > 2{
			initializers.DB.Migrator().DropTable(&models.ThreeOddsBet{})
			initializers.DB.Migrator().CreateTable(&models.ThreeOddsBet{})
			for _,arbs := range threeArbs{
				initializers.DB.Create(&arbs)
			}
		}
		if len(twoArbs) > 2{
			initializers.DB.Migrator().DropTable(&models.TwoOddsBet{})
			initializers.DB.Migrator().CreateTable(&models.TwoOddsBet{})
			for _,arbs := range twoArbs{
				initializers.DB.Create(&arbs)
			}
		}
		log.Println("Arbs successfully added")
	}	
}

func GetTwoArbs(c *gin.Context){
	var twoArbs []models.TwoOddsBet
	initializers.DB.Order("profit desc").Find(&twoArbs)
	c.JSON(http.StatusOK,gin.H{
		"twoArbs": twoArbs,
	})
}

func GetThreeArbs(c *gin.Context){
	var threeArbs []models.ThreeOddsBet
	initializers.DB.Order("profit desc").Find(&threeArbs)
	c.JSON(http.StatusOK,gin.H{
		"threeArbs":threeArbs,
	})
}
