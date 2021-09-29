package weather_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"weather"
)

func TestConstructUrl(t *testing.T) {

	token := "foo"
	args := []string{"PATH", "rio", "de", "janeiro"}
	want := "https://api.openweathermap.org/data/2.5/weather?q=rio%20de%20janeiro&appid=foo"

	got := weather.BuildURL(token, args)

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

	url := "https://api.openweathermap.org/data/2.5/weather?q=rio%20de%20janeiro&appid=foo"
	got, err := weather.Get(url)

	if err != nil {
		t.Fatal(err)
	}

	want := weather.Conditions{
		OneWord: "Clouds",
		Celcius: 23.0,
		City:    "Birmingham",
	}

	fmt.Printf("%v", want)
	fmt.Printf("%v", got)

	// if want != got {
	// 	t.Fatal("want not equal to got")
	// }
}
