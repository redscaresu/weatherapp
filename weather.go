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

type apiErrorResponse struct {
	Cod     string
	Message string
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
		log.Printf("an error has occured, %v", err)
		os.Exit(2)
	}

	response, err := Response(request)
	if err != nil {
		log.Printf("an error has occured, %v", err)
		os.Exit(2)
	}

	conditions, err := ParseResponse(response)
	if err != nil {
		log.Printf("problem parsing API response', %v", err)
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
		log.Printf("problem setting url', %v", err)
		return request, err

	}
	return request, err
}

func Response(request *http.Request) (*http.Response, error) {

	client := &http.Client{}

	resp, err := client.Do(request)

	return resp, err
}

func ParseResponse(resp *http.Response) (Conditions, error) {

	var a apiResponse
	var c Conditions

	var cu apiErrorResponse

	if resp.StatusCode == http.StatusNotFound {
		err := json.NewDecoder(resp.Body).Decode(&cu)
		if err != nil {
			log.Printf("an error has occured, %v", err)
		}

		if cu.Message == "city not found" {
			fmt.Printf("The city cannot be found, the error code is %v \n", resp.StatusCode)
		}
	}

	read_all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("an error has occured, %v", err)
	}

	err = json.Unmarshal(read_all, &a)
	if err != nil {
		log.Printf("an error has occured, %v", err)
	}

	Celsius := a.Main.Temp - 273.15
	mainWeather := a.Weather[0].Main

	c.TemperatureCelsius = math.Round(Celsius)
	c.OneWord = mainWeather
	c.City = a.Name

	reqBodyBytes := new(bytes.Buffer)

	json.NewEncoder(reqBodyBytes).Encode(c)
	if err != nil {
		log.Printf("an error has occured, %v", err)
	}

	return c, nil
}
