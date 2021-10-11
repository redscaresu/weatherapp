package main

import (
	"fmt"
	"os"
	"weather"
)

func main() {

	resp := weather.RunCLI(os.Args)
	fmt.Print(resp)
}
