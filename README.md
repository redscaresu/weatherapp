# weatherapp

`weatherapp` is a Go library and command-line client for the [Weather](api.openweathermap.org) website wearther information service. It allows you to search for the weather based on the City name.

## Installing the command-line client

To install the client binary, run:

```
go get -u github.com/redscaresu/weatherapp
```

## Using the command-line client

To use the client run:

```
weatherapp london

City: London
Weather: Clouds
Celsius: 13
```

## Setting your API key

To use the client, you will need an API Key for the account. Go to the [Weatherapp](https://openweathermap.org/appid) and follow the instructions.

Before you can use the weatherapp you must set the api key as an env variable like so `export WEATHERAPP_TOKEN=$TOKEN_ID`
