package main

import (
	"flag"
	"fmt"
	"os"
	"weather"
)

func main() {

	location := flag.String("location", "", "a city")
	flag.Parse()

	fmt.Printf("%s", *location)

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

	foo, err := weather.Get(token, *location)
	fmt.Printf(string(foo), err)
}
