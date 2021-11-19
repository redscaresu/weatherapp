package weather_test

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
	"weather"

	"github.com/google/go-cmp/cmp"
)

func TestNewClient(t *testing.T) {
	t.Parallel()
	got := weather.NewClient("dummyToken")
	want := &weather.Client{
		Token:   "dummyToken",
		BaseURL: "https://api.openweathermap.org",
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestConstructUrl(t *testing.T) {
	t.Parallel()
	location, err := weather.ParseArgs([]string{"rio", "de", "janeiro"})
	if err != nil {
		t.Fatal(err)
	}
	want := "https://api.openweathermap.org/data/2.5/weather?q=rio+de+janeiro&appid=dummyToken"
	client := weather.NewClient("dummyToken")
	got := client.Request(location)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("\n%q\n%q", want, got)
	}
}

func TestResponse(t *testing.T) {
	t.Parallel()
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open("testdata/weather.json")
		if err != nil {
			t.Fatal(err)
		}
		io.Copy(w, file)
	}))
	defer ts.Close()
	client := weather.NewClient("dummyToken")
	client.BaseURL = ts.URL
	client.HTTPClient = ts.Client()
	got, err := client.GetWeather("dummy location")
	if err != nil {
		t.Fatal(err)
	}
	want := weather.Conditions{
		OneWord:            "Clouds",
		TemperatureCelsius: 11.590000000000032,
		City:               "Birmingham",
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestParseResponseWeather(t *testing.T) {
	t.Parallel()
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open("testdata/weather.json")
		if err != nil {
			t.Fatal(err)
		}
		io.Copy(w, file)
	}))
	defer ts.Close()

	client := ts.Client()
	res, err := client.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	want := weather.Conditions{
		OneWord:            "Clouds",
		TemperatureCelsius: 11.590000000000032,
		City:               "Birmingham",
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	got, err := weather.ParseResponse(bodyBytes)
	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Fatal(cmp.Diff(want, got))
	}
}

func TestCreateString(t *testing.T) {
	t.Parallel()
	input := weather.Conditions{
		OneWord:            "Clouds",
		TemperatureCelsius: 26.01,
		City:               "Birmingham",
	}

	want := "Clouds 26.0ºC"
	got := input.String()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
