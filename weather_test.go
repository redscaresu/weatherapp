package weather_test

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
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

	input := weather.Conditions{
		OneWord:            "Clouds",
		TemperatureCelsius: 26.01,
		City:               "Birmingham",
	}

	want := ""
}
