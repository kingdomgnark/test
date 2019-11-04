package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
	"strconv"
)


	type Daily struct {
		Currenttemp float32 `json:"temp"`
		Mintemp		float32	`json:"temp_min"`
		Maxtemp		float32	`json:"temp_max"`
		Humidity	int	`json:"humidity"`
	}

func main() {
    jsonFile, err := os.Open("jsontest1.json")
    if err != nil {
        fmt.Println(err)
    }
	
	
	
    fmt.Println("Successfully Opened jsontest1.json")

    defer jsonFile.Close()

    byteValue, _ := ioutil.ReadAll(jsonFile)

	var data struct{Main Daily}
	
	json.Unmarshal(byteValue, &data)
	

	for i := 0; i < len(v.Main); i++ {
		strCurrenttemp := fmt.Sprintf("%f", v.Main.Currenttemp)
		strMintemp := fmt.Sprintf("%f", v.Main.Mintemp)
		strMaxtemp := fmt.Sprintf("%f", v.Main.Maxtemp)	
		fmt.Println("Current Temperature:"+ strCurrenttemp)
		fmt.Println("Today's Low:"+ strMintemp)
		fmt.Println("Today's High:"+ strMaxtemp)
		fmt.Println("Current Humidity:"+ strconv.Itoa(v.Main.Humidity)+"%")
		}
}