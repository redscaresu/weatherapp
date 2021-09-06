package main

import (
	"fmt"
	"os"
	"strings"
	"weather"
)

func LocationFromArgs(args []string) string {

	var command strings.Builder

	for i, v := range os.Args {
		if i > 0 {
			command.WriteString(v)
			command.WriteString(" ")
		}
	}

	return command.String()
}

func main() {

	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "please set a location e.g. london\n")
		os.Exit(2)
	}

	location := LocationFromArgs(os.Args)

	token := os.Getenv("WEATHERAPP_TOKEN")
	if len(token) == 0 {
		fmt.Fprintf(os.Stderr, "please set env variable, WEATHERAPP_TOKEN \n")
		os.Exit(2)
	}

	resp := weather.CliOutput(token, location)
	fmt.Printf("%v", resp)
}
