package weather

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type apiResponse struct {
	Name    string
	Weather []struct {
		Main string
	}
	Main struct {
		Temp float64
	}
}

type Conditions struct {
	City               string
	OneWord            string
	TemperatureCelsius float64
}

func (c Conditions) String() string {
	return fmt.Sprintf("%s %.1fºC", c.OneWord, c.TemperatureCelsius)
}

func RunCLI() {
	token := os.Getenv("WEATHERAPP_TOKEN")
	if token == "" {
		fmt.Fprintf(os.Stderr, "please set env variable, WEATHERAPP_TOKEN \n")
		os.Exit(1)
	}

	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s LOCATION\n", os.Args[0])
		os.Exit(1)
	}

	location, err := ParseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	client := NewClient(token)
	conditions, err := client.GetWeather(location)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(conditions)
}

func ParseArgs(args []string) (string, error) {
	location := strings.Join(args, " ")
	if location == "" {
		return "", errors.New("the location is empty")
	}
	return location, nil
}

func ParseResponse(r []byte) (Conditions, error) {
	var a apiResponse
	var c Conditions

	err := json.Unmarshal(r, &a)
	if err != nil {
		return Conditions{}, err
	}

	if a.Name == "" {
		return Conditions{}, fmt.Errorf("empty apiResponse struct: %v", apiResponse{})
	}

	Celsius := a.Main.Temp - 273.15
	mainWeather := a.Weather[0].Main

	c.TemperatureCelsius = Celsius
	c.OneWord = mainWeather
	c.City = a.Name

	return c, nil
}
