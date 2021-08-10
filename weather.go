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
)

type Weather struct {
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
}

type CityUnknown struct {
	Cod     string `json:"cod"`
	Message string `json:"message"`
}

type CliResponse struct {
	OneWord string  `json:"oneword"`
	Celcius float64 `json:"celcius"`
}

type Response struct {
	OneWord string  `json:"oneword"`
	Celcius float64 `json:"celcius"`
}

func CliOutput(token, location string) (output string) {

	var r Response

	resp := CallUrl(token, location)
	weatherString, err := Get(resp)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(weatherString, &r)

	output = string(fmt.Sprintf("weather: %s\ncelcius: %v\n", r.OneWord, r.Celcius))
	return output
}

func CallUrl(token, location string) http.Response {

	var cu CityUnknown
	domain := "api.openweathermap.org"

	resp, err := http.Get(fmt.Sprintf("https://%s/data/2.5/weather?q=%s&appid=%s", domain, location, token))

	//fail as early as possible
	if err != nil {
		log.Fatal(err)
	}

	//I dont like this, 404 could be different reasons.  I assume its always because a city cannot be found.
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

	return *resp
}

func Get(resp http.Response) ([]byte, error) {

	var w Weather
	var c CliResponse

	read_all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(read_all, &w)
	if err != nil {
		log.Fatal(err)
	}

	celcius := w.Main.Temp - 273.15
	mainWeather := w.Weather[0].Main

	c.Celcius = math.Round(celcius)
	c.OneWord = mainWeather

	reqBodyBytes := new(bytes.Buffer)

	json.NewEncoder(reqBodyBytes).Encode(c)
	if err != nil {
		log.Fatal(err)
	}

	return reqBodyBytes.Bytes(), err
}
