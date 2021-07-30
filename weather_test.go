package weather_test

import (
	"fmt"
	"os"
	"testing"
	"weather"
)

func TestGet(t *testing.T) {
	var conditions []byte

	token := os.Getenv("WEATHERAPP_TOKEN")
	if len(token) == 0 {
		fmt.Printf("please set a weatherapp token\n")
		os.Exit(2)
	}

	conditions, err := weather.Get("london", token)
	if err != nil {
		t.Fatal(err)
	}
	if len(conditions) == 0 {
		t.Fatal("no conditions")
	}
}
