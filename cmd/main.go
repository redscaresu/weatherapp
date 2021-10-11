package main

import (
	"fmt"
	"os"
	"weather"
)

func main() {

	resp := weather.RunCli(os.Args)
	fmt.Print(resp)
}
