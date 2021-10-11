package weather_test

import (
	"testing"
	"weather"
)

func TestConstructUrl(t *testing.T) {

	token := "foo"
	response := weather.Request([]string{"PATH", "rio", "de", "janeiro"}, token)
	got := response.URL.String()
	want := "https://api.openweathermap.org/data/2.5/weather?q=rio%20de%20janeiro&appid=foo"

	if want != got {
		t.Fatalf("want: %q got: %q", want, got)
	}

}

// func TestParseResponseWeather(t *testing.T) {

// 	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		file, err := os.Open("testdata/weather.json")
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		io.Copy(w, file)
// 	}))
// 	defer ts.Close()

// 	client := ts.Client()
// 	res, err := client.Get(ts.URL)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	want := weather.Conditions{
// 		OneWord:            "Clouds",
// 		TemperatureCelsius: 23.0,
// 		City:               "Birmingham",
// 	}

// 	got, err := weather.ParseResponse(res)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if want != got {
// 		t.Fatal("want not equal to got")
// 	}
// }
