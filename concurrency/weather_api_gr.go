package concurrency

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// const apikey_gr = "c1e91754906e46eaab3729738ec61306"
const apikey_gr = "bd5e378503939ddaee76f12ad7a97608"

// https://api.openweathermap.org/data/2.5/weather?q=chennai&appid=c1e91754906e46eaab3729738ec61306


type WeatherGR struct {
	ID int `json:"id"`
	Main string `json:"main"`
	Description string `json:"description"`
	Icon string `json:"icon"`
}

type ApiResultGR struct {
	Name string `json:"name"`
	Id int `json:"id"`
	TimeZone int `json:"timezone"`
	WeatherGR []WeatherGR `json:"weather"`
}

func weather_api_gr_call(cityName string,ch chan ApiResultGR,wg *sync.WaitGroup) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s",cityName, apikey_gr)
	res, err := http.Get(url)

	if err != nil {
		fmt.Println("error making request: %w", err)
	}

	defer res.Body.Close()

	finalResult := ApiResultGR{}

	if err = json.NewDecoder(res.Body).Decode(&finalResult); err != nil {
		fmt.Println("error decoding JSON response", err)
	}

	wg.Done()

	ch <- finalResult

	// finalResult := make(map[string]interface{})
	// for key, value := range finalResult {
	// 	fmt.Println(key,value)
	// }
}


func WeatherApiGR() {

	start := time.Now()

	wg := sync.WaitGroup{}
	ch := make(chan ApiResultGR)

	city_name := []string{"chennai", "madurai", "trichy", "bengaluru", "coimbatore"}

	for _, value := range city_name {
		wg.Add(1)

		go weather_api_gr_call(value,ch,&wg)
	}

	go func ()  {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println("Current weather report for", result.Name)
		fmt.Println("WeatherGR ID:", result.WeatherGR[0].ID)
		fmt.Println("Main weather condition:", result.WeatherGR[0].Main)
		fmt.Println("Description:", result.WeatherGR[0].Description)
		fmt.Println("Icon:", result.WeatherGR[0].Icon)
		fmt.Println("Timezone:", result.TimeZone)
		fmt.Println("")	
	}

	elapsed := time.Since(start)
	fmt.Printf("Time taken for %s\n", elapsed)


	fmt.Println("WeatherApiNormal")
}