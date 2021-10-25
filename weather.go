package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
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

func RunCLI(args []string) {

	conditions := GetConditions(args)
	fmt.Printf("City: %s\nWeather: %s\nCelsius: %v\n", conditions.City, conditions.OneWord, conditions.TemperatureCelsius)
}

func GetConditions(args []string) Conditions {

	token := os.Getenv("WEATHERAPP_TOKEN")
	if token == "" {
		fmt.Fprintf(os.Stderr, "please set env variable, WEATHERAPP_TOKEN \n")
		os.Exit(2)
	}

	request, err := Request(args, token)
	if err != nil {
		fmt.Fprintf(os.Stderr, "problem setting url', %v\n", err)
		os.Exit(2)
	}

	response, err := Response(request)
	if err != nil {
		fmt.Fprintf(os.Stderr, "an error has occured, %v\n", err)
		os.Exit(2)
	}

	conditions, err := ParseResponse(response)
	if err != nil {
		fmt.Fprintf(os.Stderr, "problem parsing API response\n', %v", err)
		os.Exit(2)
	}

	return conditions
}

func Request(args []string, token string) (string, error) {

	domain := "api.openweathermap.org"

	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "please set a location e.g. london\n")
	}

	location := strings.Join(args[1:], "%20")

	url := fmt.Sprintf("https://%s/data/2.5/weather?q=%s&appid=%s", domain, location, token)

	return url, nil
}

func Response(url string) (io.Reader, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("http error code %v", resp.StatusCode)
	}

	return resp.Body, nil
}

func ParseResponse(r io.Reader) (Conditions, error) {

	var a apiResponse
	var c Conditions

	err := json.NewDecoder(r).Decode(&a)
	if err != nil {
		return Conditions{}, err
	}

	Celsius := a.Main.Temp - 273.15
	mainWeather := a.Weather[0].Main

	c.TemperatureCelsius = math.Round(Celsius)
	c.OneWord = mainWeather
	c.City = a.Name

	return c, nil
}
