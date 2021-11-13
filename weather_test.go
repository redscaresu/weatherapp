package weather_test

import (
	"os"
	"testing"
	"weather"
)

func TestConstructUrl(t *testing.T) {

	token := "foo"
	got, err := weather.Request([]string{"PATH", "rio", "de", "janeiro"}, token)
	if err != nil {
		t.Fatal(err)
	}

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
	defer file.Close()

	want := weather.Conditions{
		OneWord:            "Clouds",
		TemperatureCelsius: 23.0,
		City:               "Birmingham",
	}

	got, err := weather.ParseResponse(file)

	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Fatal(cmp.Diff(want, got))
	}
}
