package main

import (
	"fmt"
	"io"
	"net/http"
)

func getTemperature(city, format string) (string, error) {
	route := fmt.Sprintf("http://wttr.in/%s?format=%s", city, format)
	fmt.Println(route)

	resp, err := http.Get(route)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func main() {
	// Example usage
	temperature, err := getTemperature("Zurich", "%t")
	if err != nil {
		fmt.Println("Error fetching temperature:", err)
		return
	}
	fmt.Println("Temperature in Zurich:", temperature)

	// If you know the specific format for morning or night, use it here
	morningWeather, err := getTemperature("Zurich", "%c")
	if err != nil {
		fmt.Println("Error fetching morning weather:", err)
		return
	}
	fmt.Println("Morning weather in Zurich:", morningWeather)
}
