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

type Weather struct {
	Name    string
	Weather []struct {
		ID          int
		Main        string
		Description string
		Icon        string
	}
	Main struct {
		Temp      float64
		FeelsLike float64
		TempMin   float64
		TempMax   float64
		Humidity  int
	}
}

type CityUnknown struct {
	Cod     string `json:"cod"`
	Message string `json:"message"`
}

type Conditions struct {
	City    string
	OneWord string  `json:"oneword"`
	Celcius float64 `json:"celcius"`
}

func CliOutput(token string, args []string) (output string) {

	location := LocationFromArgs(args)
	resp := BuildURL(token, location)
	c, err := Get(resp)
	if err != nil {
		log.Fatal(err)
	}

	output = fmt.Sprintf("city: %s\nweather: %s\ncelcius: %v\n", c.City, c.OneWord, c.Celcius)

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

func Call(url string) *http.Response {

	resp, err := http.Get(url)

	if err != nil {
		log.Printf("an error has occured, %v", err)
		os.Exit(2)
	}

	return resp
}

func Get(url string) (Conditions, error) {

	var w Weather
	var c Conditions

	var cu CityUnknown

	resp, err := http.Get(url)

	if err != nil {
		log.Printf("an error has occured, %v", err)
		os.Exit(2)
	}

	if resp.StatusCode == http.StatusNotFound {
		err = json.NewDecoder(resp.Body).Decode(&cu)
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

	err = json.Unmarshal(read_all, &w)
	if err != nil {
		log.Printf("an error has occured, %v", err)
		os.Exit(2)
	}

	celcius := w.Main.Temp - 273.15
	mainWeather := w.Weather[0].Main

	c.Celcius = math.Round(celcius)
	c.OneWord = mainWeather
	c.City = w.Name

	reqBodyBytes := new(bytes.Buffer)

	json.NewEncoder(reqBodyBytes).Encode(c)
	if err != nil {
		log.Printf("an error has occured, %v", err)
		os.Exit(2)
	}

	return c, nil
}
