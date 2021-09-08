package weather

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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
	Name string `json:"name"`
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

func CliOutput(token string, location string) (output string) {

	var r Conditions

	resp := CallUrl(token, location)
	weatherString, err := Get(resp)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(weatherString, &r)

	output = fmt.Sprintf("city: %s\nweather: %s\ncelcius: %v\n", r.City, r.OneWord, r.Celcius)
	return output
}

func CallUrl(token string, location string) io.Reader {

	var r io.Reader

	var cu CityUnknown
	domain := "api.openweathermap.org"

	resp, err := http.Get(fmt.Sprintf("https://%s/data/2.5/weather?q=%s&appid=%s", domain, location, token))

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
			fmt.Printf("The city %scannot be found, the error code is %v \n", location, resp.StatusCode)
			os.Exit(2)
		}
	}

	r = resp.Body
	return r
}

func Get(resp io.Reader) ([]byte, error) {

	var w Weather
	var c Conditions

	read_all, err := ioutil.ReadAll(resp)
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
	c.City = w.Name

	reqBodyBytes := new(bytes.Buffer)

	json.NewEncoder(reqBodyBytes).Encode(c)
	if err != nil {
		log.Fatal(err)
	}

	return reqBodyBytes.Bytes(), err
}
