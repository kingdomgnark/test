package main

import (
	"fmt"
	"log"
	"net/url"
	"strconv"
	"encoding/json"
	"net/http"
    "io/ioutil"
	"html/template"
    "os"
    "reflect"
	
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


var templateFuncs = template.FuncMap{"rangeStruct": RangeStructer}

var htmlTemplate = `{{range .}}<tr>
{{range rangeStruct .}} <td>{{.}}</td>
{{end}}</tr>
{{end}}`

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
    container := []Forecast
	t := template.New("t").Funcs(templateFuncs)
    t, err := t.Parse(htmlTemplate)
    if err != nil {
        panic(err)
    }

    err = t.Execute(os.Stdout, container)
    if err != nil {
        panic(err)
    }

}

func RangeStructer(args ...interface{}) []interface{} {
    if len(args) == 0 {
        return nil
    }

    v := reflect.ValueOf(args[0])
    if v.Kind() != reflect.Struct {
        return nil
    }

    out := make([]interface{}, v.NumField())
    for i := 0; i < v.NumField(); i++ {
        out[i] = v.Field(i).Interface()
    }

    return out
}