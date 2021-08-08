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
		fmt.Fprintln(w, "hello world")
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

	// token := os.Getenv("WEATHERAPP_TOKEN")
	// if len(token) == 0 {
	// 	fmt.Printf("please set a weatherapp token\n")
	// 	os.Exit(2)
	// }

	resp := WeatherNewTLSServer()
	conditions, err := weather.Get(resp)

	if err != nil {
		t.Fatal(err)
	}
	if len(conditions) == 0 {
		t.Fatal("no conditions")
	}
}
