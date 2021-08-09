package weather_test

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"weather"
)

func WeatherNewTLSServer() (r http.Response) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "{\"coord\":{\"lon\":-86.8025,\"lat\":33.5207},\"weather\":[{\"id\":804,\"main\":\"Clouds\",\"description\":\"overcast clouds\",\"icon\":\"04n\"}],\"base\":\"stations\",\"main\":{\"temp\":296.06,\"feels_like\":296.91,\"temp_min\":294.16,\"temp_max\":297.63,\"pressure\":1018,\"humidity\":96,\"sea_level\":1018,\"grnd_level\":996},\"visibility\":10000,\"wind\":{\"speed\":1.03,\"deg\":131,\"gust\":0.99},\"clouds\":{\"all\":99},\"dt\":1628495953,\"sys\":{\"type\":2,\"id\":2006051,\"country\":\"US\",\"sunrise\":1628507130,\"sunset\":1628555996},\"timezone\":-18000,\"id\":4049979,\"name\":\"Birmingham\",\"cod\":200}")
	}))
	defer ts.Close()

	client := ts.Client()
	res, err := client.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}

	return *res
}

func TestGetWeather(t *testing.T) {

	var conditions []byte

	resp := WeatherNewTLSServer()
	conditions, err := weather.Get(resp)

	if err != nil {
		t.Fatal(err)
	}
	if len(conditions) == 0 {
		t.Fatal("no conditions")
	}
}
