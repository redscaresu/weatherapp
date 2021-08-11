package weather_test

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"weather"
)

func GetTestCases() []string {

	file, err := os.Open("testcases.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	lines := make([]string, 0)

	// Read through 'tokens' until an EOF is encountered.
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func WeatherNewTLSServer(testcase string) (r http.Response) {

	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, testcase)
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

	TestCases := GetTestCases()
	for _, v := range TestCases {
		resp := WeatherNewTLSServer(v)
		conditions, err := weather.Get(resp)

		if err != nil {
			t.Fatal(err)
		}
		if len(conditions) == 0 {
			t.Fatal("no conditions")
		}
	}
}
