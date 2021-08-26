package weather_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"weather"
)

type Response struct {
	OneWord string  `json:"oneword"`
	Celcius float64 `json:"celcius"`
	City    string  `json:"city"`
}

func TestGetWeather(t *testing.T) {

	var r Response

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
		City:    "Birmingham",
		OneWord: "Clouds",
		Celcius: 10,
	}

	got, err := weather.Get(*res)

	if err != nil {
		t.Fatal(err)
	}

	json.Unmarshal(got, &r)
	OneWordResponse := &r.OneWord
	CityResponse := &r.City

	if want.OneWord != *OneWordResponse {
		t.Fatal("want not equal to got")
	}

	if want.City != *CityResponse {
		t.Fatal("want not equal to got")
	}
}
