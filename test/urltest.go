package main

import (
	"fmt"
	"net/url"
)

func main() {
	// Let's start with a base url
	url := ("api.openweathermap.org/data/2.5/forecast?zip=")

	fmt.Printf(url)
}