package weather

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

type Client struct {
	Token      string
	BaseURL    string
	HTTPClient *http.Client
}

func NewClient(token string) *Client {
	return &Client{
		BaseURL: "https://api.openweathermap.org",
		Token:   token,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c Client) GetWeather(location string) (Conditions, error) {
	URI := c.Request(location)
	data, err := c.Response(URI)
	if err != nil {
		return Conditions{}, err
	}
	conditions, err := ParseResponse(data)
	if err != nil {
		return Conditions{}, err
	}
	return conditions, nil
}

func (c Client) Request(location string) string {
	return fmt.Sprintf("%s/data/2.5/weather?q=%s&appid=%s", c.BaseURL, url.QueryEscape(location), c.Token)
}

func (c *Client) Response(url string) ([]byte, error) {
	r, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}

	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response status: %s", r.Status)
	}

	if r.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("location not found: %q", os.Args[0])
	}

	resp, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
