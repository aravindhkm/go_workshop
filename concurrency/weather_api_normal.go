package concurrency

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// https://api.openweathermap.org/data/2.5/weather?q=chennai&appid=c1e91754906e46eaab3729738ec61306

type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type ApiResult struct {
	Name     string    `json:"name"`
	Id       int       `json:"id"`
	TimeZone int       `json:"timezone"`
	Weather  []Weather `json:"weather"`
}

func weather_api_normal_call(cityName string) {

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", cityName, apikey_gr)
	res, err := http.Get(url)

	if err != nil {
		fmt.Println("error making request: %w", err)
	}

	defer res.Body.Close()

	finalResult := ApiResult{}

	if err = json.NewDecoder(res.Body).Decode(&finalResult); err != nil {
		fmt.Println("error decoding JSON response", err)
	}

	fmt.Println("Current weather report for", finalResult.Name)
	fmt.Println("Weather ID:", finalResult.Weather[0].ID)
	fmt.Println("Main weather condition:", finalResult.Weather[0].Main)
	fmt.Println("Description:", finalResult.Weather[0].Description)
	fmt.Println("Icon:", finalResult.Weather[0].Icon)
	fmt.Println("Timezone:", finalResult.TimeZone)
	fmt.Println("")

	// finalResult := make(map[string]interface{})
	// for key, value := range finalResult {
	// 	fmt.Println(key,value)
	// }
}

func WeatherApiNormal() {

	start := time.Now()

	city_name := []string{"chennai", "madurai", "trichy", "bengaluru", "coimbatore"}

	for _, value := range city_name {
		weather_api_normal_call(value)
	}

	elapsed := time.Since(start)

	fmt.Printf("Time taken for %s\n", elapsed)

	fmt.Println("WeatherApiNormal")
}
