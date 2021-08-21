package weather_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"weather"
)

func GetTestCases() string {

	b, err := ioutil.ReadFile("testcases.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	testcases := string(b)
	return testcases
}

func WeatherNewTLSServer(testcases string) (r http.Response) {

	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, testcases)
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

	// var conditions []byte

	testcases := GetTestCases()
	resp := WeatherNewTLSServer(testcases)
	conditions, err := weather.Get(resp)

	if err != nil {
		t.Fatal(err)
	}
	if len(conditions) == 0 {
		t.Fatal("no conditions")
	}
}
