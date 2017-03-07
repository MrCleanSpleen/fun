package main

import (
	"encoding/json"
	"fmt"
	//"log"
	"net/http"
	"net/url"
	"time"
	//	s "strings"
)

func get_json(url string, target interface{}) error {
	//if it works, it works
	var myClient = &http.Client{Timeout: 10 * time.Second}
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}
func Weather(loc string) {
	units := "imperial"
	loct := url.QueryEscape(loc)
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=", loct, "&appid=6a4cbda1e239084151dea640a95d2a0c&units=", units)
	type Data struct {
		Name string `json:"name"`
		Temp int    `json:"main,temp"`
		//Main []struct {
		//	Temp     float64 `json:"temp"`
		//	Pressure int     `json:"pressure"`
		//}
	}
	data2 := Data{}
	get_json(url, &data2)
	fmt.Println(data2.Temp)

}
func main() {
	Weather("charlotte")
}
