package weather

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

type CityUnknown struct {
	Cod     string
	Message string
}

type Conditions struct {
	City    string
	OneWord string
	Celsius float64
}

func CliOutput(args []string) (output string) {

	token := os.Getenv("WEATHERAPP_TOKEN")
	if len(token) == 0 {
		fmt.Fprintf(os.Stderr, "please set env variable, WEATHERAPP_TOKEN \n")
		os.Exit(2)
	}

	location := LocationFromArgs(args)
	url := BuildURL(token, location)
	callURL := CallURL(url)
	conditions, err := Get(callURL)
	if err != nil {
		log.Fatal(err)
	}

	output = fmt.Sprintf("city: %s\nweather: %s\nCelsius: %v\n", conditions.City, conditions.OneWord, conditions.Celsius)

	return output
}

func LocationFromArgs(args []string) string {

	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "please set a location e.g. london\n")
		os.Exit(2)
	}

	location := strings.Join(args[1:], "%20")

	return location
}

func BuildURL(token string, location string) string {

	domain := "api.openweathermap.org"

	BuildURL := fmt.Sprintf("https://%s/data/2.5/weather?q=%s&appid=%s", domain, location, token)

	return BuildURL
}

func CallURL(url string) *http.Response {

	resp, err := http.Get(url)

	if err != nil {
		log.Printf("an error has occured, %v", err)
		os.Exit(2)
	}

	return resp
}

func Get(resp *http.Response) (Conditions, error) {

	var a apiResponse
	var c Conditions

	var cu CityUnknown

	if resp.StatusCode == http.StatusNotFound {
		err := json.NewDecoder(resp.Body).Decode(&cu)
		if err != nil {
			log.Printf("an error has occured, %v", err)
		}

		if cu.Message == "city not found" {
			fmt.Printf("The city cannot be found, the error code is %v \n", resp.StatusCode)
			os.Exit(2)
		}
	}

	read_all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("an error has occured, %v", err)
		os.Exit(2)
	}

	err = json.Unmarshal(read_all, &a)
	if err != nil {
		log.Printf("an error has occured, %v", err)
		os.Exit(2)
	}

	Celsius := a.Main.Temp - 273.15
	mainWeather := a.Weather[0].Main

	c.Celsius = math.Round(Celsius)
	c.OneWord = mainWeather
	c.City = a.Name

	reqBodyBytes := new(bytes.Buffer)

	json.NewEncoder(reqBodyBytes).Encode(c)
	if err != nil {
		log.Printf("an error has occured, %v", err)
		os.Exit(2)
	}

	return c, nil
}
