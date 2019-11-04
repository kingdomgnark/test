package main

import (
	"fmt"
	"log"
	"net/url"
	"strconv"
	"encoding/json"
	"net/http"
    "io/ioutil"
	
)

type Forecast struct {
    Forecast []Dayweather `json:"list"`
}

// User struct which contains a name
// a type and a list of social links
type Dayweather struct {
	Day_Time string	`json:"dt_txt"`
    Weather Weather `json:"main"`
}

type Weather struct {
	Mintemp		float32	`json:"temp_min"`
	Maxtemp		float32	`json:"temp_max"`
	Humidity	int	`json:"humidity"`
}

func main() {
	u, err := url.Parse("api.openweathermap.org")
	if err != nil {
		log.Fatal(err)
	}
	zip := 95621
	s := strconv.Itoa(zip)
	u.Scheme = "http"
	u.Host = "api.openweathermap.org"
	u.Path = "data/2.5/forecast"
	q := u.Query()
	q.Set("APPID","12ea4c57a813ac94155a0cabe7788b15")
	q.Set("zip", s)
	q.Set("units","imperial")
	u.RawQuery = q.Encode()
	url1 := u.String()
	res, err := http.Get(url1)
	defer res.Body.Close()
	
	byteValue, _ := ioutil.ReadAll(res.Body)
	
	
	var forecast Forecast

    // we unmarshal our byteArray which contains our
    // jsonFile's content into 'users' which we defined above
    json.Unmarshal(byteValue, &forecast)

    // we iterate through every user within our users array and
    // print out the user Type, their name, and their facebook url
    // as just an example
    for i := 0; i < len(forecast.Forecast); i++ {
		strMintemp := fmt.Sprintf("%f", forecast.Forecast[i].Weather.Mintemp)
		strMaxtemp := fmt.Sprintf("%f", forecast.Forecast[i].Weather.Maxtemp)
        fmt.Println("Date and Time: " + forecast.Forecast[i].Day_Time)
		fmt.Println("Today's Low:"+ strMintemp)
		fmt.Println("Today's High:"+ strMaxtemp)
		fmt.Println("Current Humidity:"+ strconv.Itoa(forecast.Forecast[i].Weather.Humidity)+"%")
    }

}