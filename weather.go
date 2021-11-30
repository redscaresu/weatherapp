package weather

import (
	"encoding/json"
	"errors"
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
	TempScale          struct {
		Farenheit float64
		Celsius   float64
		Kelvin    float64
		Rankine   float64
	}
}

func (c Conditions) String() string {
	return fmt.Sprintf("%s %.1fºC", c.OneWord, c.TempScale.Celsius)
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

	mainWeather := a.Weather[0].Main

	c.TempScale.Farenheit = a.Main.Temp
	c.TempScale.Celsius = a.Main.Temp - 273.15
	c.TempScale.Kelvin = (a.Main.Temp-32)*5/9 + 273.15
	c.TempScale.Rankine = a.Main.Temp + 459.67

	c.OneWord = mainWeather
	c.City = a.Name

	return c, nil
}
