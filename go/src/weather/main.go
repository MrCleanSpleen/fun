package main

import (
	"encoding/json"
	"fmt"
	//"log"
	"net/http"
	//"net/url"
	"time"
	//	s "strings"
)

func get_json(url string, target interface{}) error {
	// Gets JSON data from a Web API and puts it into a struct
	var myClient = &http.Client{Timeout: 10 * time.Second}
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}
func Weather(loc string) {
	// Gets weather data for a given city and returns it.
	units := "imperial"
	//loct := url.QueryEscape(loc)
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=london&appid=b6907d289e10d714a6e88b30761fae22&units=", units)
	type Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	}
	type Data struct {
		Name    string             `json:"name"`
		Id      int                `json:"id"`
		Main    map[string]float64 `json:"main"`
		Sys     map[string]float64 `json:"sys"`
		Coord   map[string]float64 `json:"coord"`
		Weather map[string]float64 `json:"weather"`
	}
	get_json(url, &Data{})
	data2 := Data{}
	temp := (data2.Main["temp"] * 1.8) - 459.67
	fmt.Println("In", data2.Name, ", the temperature is", temp, "F", ", with a humidity of", data2.Main["humidity"], "%.")
	print(data2.Name)
}
func main() {
	Weather("london")
}
