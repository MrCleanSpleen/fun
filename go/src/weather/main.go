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
		Name     string `json:"name"`
		Temp     int    `json:"main,temp"`
		Country  string `json:"sys,country"`
		Temp_min int    `json:"main,temp_min"`
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
	var record Data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}
	fmt.Println("City: ", record.Name)
	fmt.Println("Temperature: ", record.Temp)
}
func main() {
	Weather("charlotte")
}
