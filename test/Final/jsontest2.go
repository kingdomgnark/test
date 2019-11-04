package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "strconv"
)

// Users struct which contains
// an array of users
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
    // Open our jsonFile
    jsonFile, err := os.Open("json2.json")
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Successfully Opened users.json")
    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()

    // read our opened xmlFile as a byte array.
    byteValue, _ := ioutil.ReadAll(jsonFile)

    // we initialize our Users array
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