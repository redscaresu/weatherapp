package weather_test

import (
	"bufio"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"weather"
)

func GetTestCases() []string {

	file, err := os.Open("testcases.txt")
	if err != nil {
		fmt.Printf("%v", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func WeatherNewTLSServer(testcases string) (r http.Response) {

	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, testcases)
	}))
	defer ts.Close()

	client := ts.Client()
	res, err := client.Get(ts.URL)
	if err != nil {
		fmt.Printf("%v", err)
	}

	return *res
}

func TestGetWeather(t *testing.T) {

	// var conditions []byte

	testcases := GetTestCases()
	for _, testcase := range testcases {
		resp := WeatherNewTLSServer(testcase)
		conditions, err := weather.Get(resp)

		if err != nil {
			t.Fatal(err)
		}
		if len(conditions) == 0 {
			t.Fatal("no conditions")
		}
	}
}
