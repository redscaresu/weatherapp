package weather_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"weather"
)

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
		city:    "Birmingham",
		weather: "rain",
		celcius: 10,
	}

	got, err := weather.Get(*res)

	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Fatal("want not equal to got")
	}

}
