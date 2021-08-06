package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"weather"
)

type Response struct {
	OneWord string  `json:"oneword"`
	Celcius float64 `json:"celcius"`
}

func main() {

	var r Response

	location := flag.String("location", "", "a city")
	flag.Parse()

	if len(*location) == 0 {
		fmt.Printf("please enter a location\n")
		os.Exit(2)
	}
	flag.Parse()

	token := os.Getenv("WEATHERAPP_TOKEN")
	if len(token) == 0 {
		fmt.Printf("please set a weatherapp token\n")
		os.Exit(2)
	}

	weatherString, err := weather.Get(token, *location)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(weatherString, &r)

	fmt.Printf("weather: %s\ncelcius: %v\n", r.OneWord, r.Celcius)
}
