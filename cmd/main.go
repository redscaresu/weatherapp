package main

import (
	"fmt"
	"os"
	"weather"
)

// type Response struct {
// 	OneWord string  `json:"oneword"`
// 	Celcius float64 `json:"celcius"`
// 	City    string
// }

func main() {

	location := os.Args[1]

	if len(os.Args) == 0 {
		fmt.Fprintf(os.Stderr, "please set a location e.g. london\n")
		os.Exit(2)
	}

	token := os.Getenv("WEATHERAPP_TOKEN")
	if len(token) == 0 {
		fmt.Fprintf(os.Stderr, "please set env variable, WEATHERAPP_TOKEN \n")
		os.Exit(2)
	}

	resp := weather.CliOutput(token, location)
	fmt.Printf("%v", resp)
}
