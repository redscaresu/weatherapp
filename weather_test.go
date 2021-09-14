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

type Conditions struct {
	City    string  `json:"city"`
	OneWord string  `json:"oneword"`
	Celcius float64 `json:"celcius"`
}

func TestGetWeather(t *testing.T) {

	var c Conditions
	var re io.Reader

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

	re = res.Body

	want := Conditions{
		OneWord: "Clouds",
		Celcius: 23.0,
		City:    "Birmingham",
	}

	got, err := weather.Get(re)

	if err != nil {
		t.Fatal(err)
	}

	json.Unmarshal(got, &c)
	// OneWordResponse := &r.OneWord
	// CityResponse := &r.City
	// CelciusResonse := &r.Celcius

	if want != c {
		t.Fatal("want not equal to got")
	}

	// if want.OneWord != *OneWordResponse {
	// 	t.Fatal("want not equal to got")
	// }

	// if want.City != *CityResponse {
	// 	t.Fatal("want not equal to got")
	// }

	// if want.Celcius != *CelciusResonse {
	// 	t.Fatal("want not equal to got")
	// }
}
