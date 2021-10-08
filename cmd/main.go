package main

import (
	"fmt"
	"os"
	"weather"
)

func main() {

	resp := weather.CliOutput(os.Args)
	fmt.Print(resp)
}
