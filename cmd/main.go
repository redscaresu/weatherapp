package main

import (
	"fmt"
	"os"
	"weather"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "please set a location e.g. london\n")
		os.Exit(2)
	}

	location := os.Args[1]

	token := os.Getenv("WEATHERAPP_TOKEN")
	if len(token) == 0 {
		fmt.Fprintf(os.Stderr, "please set env variable, WEATHERAPP_TOKEN \n")
		os.Exit(2)
	}

	resp := weather.CliOutput(token, location)
	fmt.Printf("%v", resp)
}
