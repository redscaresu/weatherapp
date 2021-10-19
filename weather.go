package weather

import (
	"encoding/json"
	"errors"
	"fmt"
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

	token := os.Getenv("WEATHERAPP_TOKEN")
	if token == "" {
		fmt.Fprintf(os.Stderr, "please set env variable, WEATHERAPP_TOKEN \n")
		os.Exit(2)
	}

	request, err := Request(args, token)
	if err != nil {
		fmt.Fprintf(os.Stderr, "problem setting url', %v", err)
		os.Exit(2)
	}

	response, err := Response(request)
	if err != nil {
		fmt.Fprintf(os.Stderr, "an error has occured, %v", err)
		os.Exit(2)
	}

	conditions, err := ParseResponse(response)
	if err != nil {
		fmt.Fprintf(os.Stderr, "problem parsing API response', %v", err)
		os.Exit(2)
	}

	fmt.Printf("City: %s\nWeather: %s\nCelsius: %v\n", conditions.City, conditions.OneWord, conditions.TemperatureCelsius)

}

func Request(args []string, token string) (*http.Request, error) {

	domain := "api.openweathermap.org"

	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "please set a location e.g. london\n")
	}

	location := strings.Join(args[1:], "%20")

	url := fmt.Sprintf("https://%s/data/2.5/weather?q=%s&appid=%s", domain, location, token)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func Response(request *http.Request) (*http.Response, error) {

	client := &http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func ParseResponse(resp *http.Response) (Conditions, error) {

	var a apiResponse
	var c Conditions

	if resp.StatusCode == http.StatusNotFound {
		return Conditions{}, errors.New("location not found")
	}

	err := json.NewDecoder(resp.Body).Decode(&a)
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
