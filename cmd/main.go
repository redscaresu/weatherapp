package main

import (
	"flag"
	"fmt"
	"os"
	"weather"
)

type Response struct {
	OneWord string  `json:"oneword"`
	Celcius float64 `json:"celcius"`
}

func main() {

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

	resp := weather.CliOutput(token, *location)
	fmt.Printf("%v", resp)
}
