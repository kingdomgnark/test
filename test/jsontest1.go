package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
	"bytes"
)

type ignored [0]byte

func (i *ignored) UnmarshalJSON([]byte) error {
		return nil
	}
	
type Daily struct {
	Currenttemp float64 `json:"temp"`
	Mintemp     float64 `json:"temp_min"`
	Maxtemp     float64 `json:"temp_max"`
	Humidity    float64 `json:"humidity"`
	Pressure    ignored
	}

type overall struct {
		Cod			ignored
		Name		ignored
		Sunset		ignored
		Sunrise		ignored
		Country		ignored
		Type		ignored
		Id			ignored
		Timezone	ignored
		Sys			ignored
		Dt			ignored
		Clouds		ignored
		Wind		ignored
		Visibility 	ignored
		Coord   	ignored
		Weather 	ignored
		Base    	ignored
		Main    	Daily
	}

func main() {
    jsonFile, err := os.Open("jsontest1.json")
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Successfully Opened jsontest1.json")

    defer jsonFile.Close()

    byteValue, _ := ioutil.ReadAll(jsonFile)
	

	dec := json.NewDecoder(bytes.NewReader(byteValue))
	dec.DisallowUnknownFields()
	dec.UseNumber()
	var all overall
	if err := dec.Decode(&all); err != nil {
		panic(err)
	}
	
		strCurrenttemp := fmt.Sprintf("%f", all.Main.Currenttemp)
		strMintemp := fmt.Sprintf("%f", all.Main.Mintemp)
		strMaxtemp := fmt.Sprintf("%f", all.Main.Maxtemp)
		strHumidity := fmt.Sprintf("%f", all.Main.Humidity)		
		fmt.Println("Current Temperature:"+ strCurrenttemp)
		fmt.Println("Today's Low:"+ strMintemp)
		fmt.Println("Today's High:"+ strMaxtemp)
		fmt.Println("Current Humidity:"+ strHumidity+"%")
}

