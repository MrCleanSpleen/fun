package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	//	s "strings"
)

func Weather(loc string) {
	units := "imperial"
	//	m := "main"
	//	s := "sys"
	loct := url.QueryEscape(loc)
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=", loct, "&appid=6a4cbda1e239084151dea640a95d2a0c&units=", units)
	type Data struct {
		Name string `json:"name"`
		//Temp float64 `json:"main,temp"`
		Main []struct {
			Temp     float64 `json:"temp"`
			Pressure int     `json:"pressure"`
		}
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}
	var weather Data
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		log.Println(err)
		//		err := json.Unmarshal(b, &record)
	}
	//[foo2 := {}
	//[	getJson("http://example.com", &foo2)
	//[	println(foo2.Bar)
	fmt.Println("City: ", weather.Name)
	//	fmt.Println("Temperature: ", weather.Temp)
	//	fmt.Println("Country: ", weather.Country)
}
func main() {
	Weather("charlotte")
}
