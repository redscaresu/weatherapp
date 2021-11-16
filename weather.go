package weather

import (
	"encoding/json"
	"fmt"
	"io"
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

func RunCLI() {

	conditions := GetConditions(args)
	fmt.Printf("City: %s\nWeather: %s\nCelsius: %v\n", conditions.City, conditions.OneWord, conditions.TemperatureCelsius)
}

func GetConditions(args []string) Conditions {

	token := os.Getenv("WEATHERAPP_TOKEN")
	if token == "" {
		fmt.Fprintf(os.Stderr, "please set env variable, WEATHERAPP_TOKEN \n")
		os.Exit(2)
	}

	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "please set a location e.g. london\n")
		os.Exit(2)
	}

	location, err := ParseArgs(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	request, err := Request(location, token)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Usage: %s LOCATION\n", os.Args[0])
		os.Exit(2)
	}

	response, err := Response(request)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	conditions, err := ParseResponse(response)
	if err != nil {
		fmt.Fprintf(os.Stderr, "problem parsing API response\n', %v", err)
		os.Exit(2)
	}

	celcius := fmt.Sprintf("%.1f", conditions.TemperatureCelsius)

	fmt.Printf("City: %s\nWeather: %s\nCelsius: %v\n", conditions.City, conditions.OneWord, celcius)
}

func ParseArgs(args []string) (string, error) {

	location := strings.Join(args[1:], "%20")
	if len(location) == 0 {
		return "", errors.New("the location is empty")
	}
	return location, nil
}

func Request(location string, token string) (string, error) {

	domain := "api.openweathermap.org"

	url := fmt.Sprintf("https://%s/data/2.5/weather?q=%s&appid=%s", domain, location, token)

	return url, nil
}

func Response(url string) ([]byte, error) {

	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http not OK, http code %v ", r.StatusCode)
	}

	if r.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("location not found: %q", os.Args[0])
	}

	resp, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func ParseResponse(r []byte) (Conditions, error) {

	var a apiResponse
	var c Conditions

	err := json.Unmarshal(r, &a)
	if err != nil {
		return Conditions{}, err
	}

	if len(a.Name) == 0 {
		return Conditions{}, fmt.Errorf("empty apiResponse struct: %v", apiResponse{})
	}

	Celsius := a.Main.Temp - 273.15
	mainWeather := a.Weather[0].Main

	c.TemperatureCelsius = Celsius
	c.OneWord = mainWeather
	c.City = a.Name

	return c, nil
}
