package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Weather struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

type CityUnknown struct {
	Cod     string `json:"cod"`
	Message string `json:"message"`
}

func GetWeather(token, location string) {

	var w Weather
	var cu CityUnknown
	//call open weather map and get something
	resp, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", location, token))

	//fail as early as possible
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == 404 {
		read_all, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(read_all, &cu)
		if err != nil {
			log.Fatal(err)
		}

		if cu.Message == "city not found" {
			fmt.Printf("%v\n", cu.Message)
			os.Exit(2)
		}
	}

	read_all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(read_all, &w)
	if err != nil {
		log.Fatal(err)
	}

	celcius := w.Main.Temp - 32*5/9

	fmt.Printf("weather: %v\ncelcius: %v", w.Weather[0].Main, celcius)
}

func main() {

	location := flag.String("location", "", "a city")
	flag.Parse()

	if len(*location) == 0 {
		fmt.Printf("please enter a location\n")
		os.Exit(2)
	}
	flag.Parse()

	token := os.Getenv("WEATHERAPP_TOKEN")
	if len(token) == 0 {
		fmt.Printf("please set a weatherapp token\n")
		os.Exit(2)
	}

	GetWeather(token, *location)
}
