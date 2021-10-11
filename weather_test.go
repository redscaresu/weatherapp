package weather_test

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"weather"
)

func TestConstructUrl(t *testing.T) {

	token := "foo"
	response, err := weather.Request([]string{"PATH", "rio", "de", "janeiro"}, token)
	if err != nil {
		t.Fatal(err)
	}

	got := response.URL.String()
	want := "https://api.openweathermap.org/data/2.5/weather?q=rio%20de%20janeiro&appid=foo"

	if want != got {
		t.Fatalf("want: %q got: %q", want, got)
	}

}

func TestParseResponseWeather(t *testing.T) {

	file, err := os.Open("testdata/weather.json")
	if err != nil {
		t.Fatal(err)
	}
	r := ioutil.NopCloser(bufio.NewReader(file))

	res, err := &http.Response{
		Status:     fmt.Sprint(http.StatusOK),
		StatusCode: http.StatusOK,
		Body:       r,
	}, nil

	want := weather.Conditions{
		OneWord:            "Clouds",
		TemperatureCelsius: 23.0,
		City:               "Birmingham",
	}

	got, err := weather.ParseResponse(res)

	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Fatal("want not equal to got")
	}
}
