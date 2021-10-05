package main

import (
	"fmt"
	"os"
	"weather"
)

func main() {

	token := os.Getenv("WEATHERAPP_TOKEN")
	if len(token) == 0 {
		fmt.Fprintf(os.Stderr, "please set env variable, WEATHERAPP_TOKEN \n")
		os.Exit(2)
	}

	resp := weather.CliOutput(token, os.Args)
	fmt.Print(resp)
}
