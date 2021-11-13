package weather_test

import (
	"io/ioutil"
	"testing"
	"weather"

	"github.com/google/go-cmp/cmp"
)

func TestConstructUrl(t *testing.T) {

	token := "foo"
	location, err := weather.ParseArgs([]string{"PATH", "rio", "de", "janeiro"})
	if err != nil {
		t.Fatal(err)
	}
	got, err := weather.Request(location, token)
	if err != nil {
		t.Fatal(err)
	}

	want := "https://api.openweathermap.org/data/2.5/weather?q=rio%20de%20janeiro&appid=foo"

	if want != got {
		t.Fatalf("want: %q got: %q", want, got)
	}

}

func TestParseResponseWeather(t *testing.T) {

	file, err := ioutil.ReadFile("testdata/weather.json")
	if err != nil {
		t.Fatal(err)
	}

	want := weather.Conditions{
		OneWord:            "Clouds",
		TemperatureCelsius: 11.590000000000032,
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
