package weather

import (
	"encoding/json"
	"errors"
	"flag"
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
	City      string
	OneWord   string
	TempScale struct {
		Farenheit float64
		Celsius   float64
		Kelvin    float64
		Rankine   float64
	}
}

func Farenheit(c Conditions) string {
	return fmt.Sprintf("%s %.1fºF", c.OneWord, c.TempScale.Farenheit)
}

func Kelvin(c Conditions) string {
	return fmt.Sprintf("%s %.1fºK", c.OneWord, c.TempScale.Kelvin)
}

func Rankine(c Conditions) string {
	return fmt.Sprintf("%s %.1fºR", c.OneWord, c.TempScale.Rankine)
}

func Celius(c Conditions) string {
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

	tempType := flag.String("temp", "f", "please choose either f,k,r or c")

	flag.Parse()

	location, err := ParseArgs(os.Args[2:])
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

	switch *tempType {
	case "f":
		fmt.Println(Farenheit(conditions))
	case "c":
		fmt.Println(Farenheit(conditions))
	case "k":
		fmt.Println(Kelvin(conditions))
	case "r":
		fmt.Println(Rankine(conditions))
	}
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
