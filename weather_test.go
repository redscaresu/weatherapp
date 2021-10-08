package weather_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"weather"
)

func TestConstructUrl(t *testing.T) {

	token := "foo"
	location := weather.LocationFromArgs([]string{"PATH", "rio", "de", "janeiro"})
	want := "https://api.openweathermap.org/data/2.5/weather?q=rio%20de%20janeiro&appid=foo"

	got := weather.BuildURL(token, location)

	if want != got {
		t.Fatalf("want: %q got: %q", want, got)
	}
}

func TestGetWeather(t *testing.T) {

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
		TemperatureCelsius: 23.0,
		City:               "Birmingham",
	}

	got, err := weather.Get(res)

	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Fatal("want not equal to got")
	}
}
